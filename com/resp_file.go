package com

import (
	"errors"
	"fmt"
	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/AaGo/aenum"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func (resp *RespStruct) toHTTPError(err error) *ae.Error {
	if os.IsNotExist(err) {
		return ae.NewE(404)
	}
	if os.IsPermission(err) {
		return ae.NewE(403)
	}
	return ae.NewE(500)
}

// The algorithm uses at most sniffLen bytes to make its decision.
const sniffLen = 512

/*
1. set Last-Modified
2. check preconditions
	* check If-Match
	* check If-Unmodified-Since
	* check If-None-Match
	* check If-Modified-Since
	* check If-Range
*/
func (resp *RespStruct) ServeFile(f string, hps ...map[string]string) error {

	info, err := os.Stat(f)

	if os.IsNotExist(err) {
		resp.Write(resp.toHTTPError)
		return fmt.Errorf("Stat file %s fail %s", f, err.Error())
	}
	// http.ServeFile(w, resp.req.r, f)
	fi, err := os.Open(f)
	if err != nil {
		resp.Write(resp.toHTTPError)
		return fmt.Errorf("Open file %s fail %s", f, err.Error())
	}
	defer fi.Close()

	sizeFunc := func() (int64, error) { return info.Size(), nil }

	return resp.serveContent(f, info.ModTime(), sizeFunc, fi)
}

func (resp *RespStruct) serveContent(name string, modtime time.Time, sizeFunc func() (int64, error), content io.ReadSeeker) error {
	resp.SetHeader("Last-Modified", modtime.UTC().Format(http.TimeFormat))
	done, rangeReq := resp.checkPreconditions(modtime)
	// 403, 412
	if done {
		return nil
	}

	//If Content-Type isn't set, use the file's extension to find it, but
	//if the Content-Type is unset explicitly, do not sniff the type.
	ctype := resp.Header(aenum.ContentType)
	if ctype == "" {
		ctype = mime.TypeByExtension(filepath.Ext(name))
		if ctype == "" {
			// read a chunk to decide between utf-8 text and binary
			var buf [sniffLen]byte
			n, _ := io.ReadFull(content, buf[:])
			ctype = http.DetectContentType(buf[:n])
			content.Seek(0, io.SeekStart) // rewind to output whole file
		}
		resp.SetHeader(aenum.ContentType, ctype)
	}

	size, err := sizeFunc()
	if err != nil {
		resp.Write(500, err.Error)
		return err
	}

	// handle Content-Range header.
	sendSize := size
	var sendContent io.Reader = content
	if size >= 0 {
		ranges, err := parseRange(rangeReq, size)
		if err != nil {
			if err == errNoOverlap {
				resp.SetHeader("Content-Range", fmt.Sprintf("bytes */%d", size))
			}
			resp.Write(416, err.Error())
			return err
		}
		if sumRangesSize(ranges) > size {
			// The total number of bytes in all the ranges
			// is larger than the size of the file by
			// itself, so this is probably an attack, or a
			// dumb client. Ignore the range request.
			ranges = nil
		}
		switch {
		case len(ranges) == 1:
			// RFC 2616, Section 14.16:
			// "When an HTTP message includes the content of a single
			// range (for example, a response to a request for a
			// single range, or to a request for a set of ranges
			// that overlap without any holes), this content is
			// transmitted with a Content-Range header, and a
			// Content-Length header showing the number of bytes
			// actually transferred.
			// ...
			// A response to a request for a single range MUST NOT
			// be sent using the multipart/byteranges media type."
			ra := ranges[0]
			if _, err := content.Seek(ra.start, io.SeekStart); err != nil {
				resp.Write(416, err.Error())
				return err
			}
			sendSize = ra.length
			resp.code = 206
			resp.SetHeader("Content-Range", ra.contentRange(size))
		case len(ranges) > 1:
			sendSize = rangesMIMESize(ranges, ctype, size)
			resp.code = 206
			pr, pw := io.Pipe()
			mw := multipart.NewWriter(pw)
			resp.SetHeader(aenum.ContentType, "multipart/byteranges; boundary="+mw.Boundary())
			sendContent = pr
			defer pr.Close() // cause writing goroutine to fail and exit if CopyN doesn't finish.
			go func() {
				for _, ra := range ranges {
					part, err := mw.CreatePart(ra.mimeHeader(ctype, size))
					if err != nil {
						pw.CloseWithError(err)
						return
					}
					if _, err := content.Seek(ra.start, io.SeekStart); err != nil {
						pw.CloseWithError(err)
						return
					}
					if _, err := io.CopyN(part, content, ra.length); err != nil {
						pw.CloseWithError(err)
						return
					}
				}
				mw.Close()
				pw.Close()
			}()
		}
		resp.SetHeader("--Ranges", "bytes")
	}

	buf := make([]byte, sendSize)
	sendContent.Read(buf)

	if resp.Header("Cache-Control") == "" {
		resp.SetHeader("Cache-Control", "max-age=3600")
	}

	resp.WriteRaw(buf)

	return nil
}

// errSeeker is returned by ServeContent's sizeFunc when the content
// doesn't seek properly. The underlying Seeker's error text isn't
// included in the sizeFunc reply so it's not sent over HTTP to end
// users.
var errSeeker = errors.New("seeker can't seek")

// errNoOverlap is returned by serveContent's parseRange if first-byte-pos of
// all of the byte-range-spec values is greater than the content size.
var errNoOverlap = errors.New("invalid range: failed to overlap")

// httpRange specifies the byte range to be sent to the client.
type httpRange struct {
	start, length int64
}

func (r httpRange) contentRange(size int64) string {
	return fmt.Sprintf("bytes %d-%d/%d", r.start, r.start+r.length-1, size)
}

func (r httpRange) mimeHeader(contentType string, size int64) textproto.MIMEHeader {
	return textproto.MIMEHeader{
		aenum.ContentRange: {r.contentRange(size)},
		aenum.ContentType:  {contentType},
	}
}

// parseRange parses a Range header string as per RFC 2616.
// errNoOverlap is returned if none of the ranges overlap.
func parseRange(s string, size int64) ([]httpRange, error) {
	if s == "" {
		return nil, nil // header not present
	}
	const b = "bytes="
	if !strings.HasPrefix(s, b) {
		return nil, errors.New("invalid range")
	}
	var ranges []httpRange
	noOverlap := false
	for _, ra := range strings.Split(s[len(b):], ",") {
		ra = strings.TrimSpace(ra)
		if ra == "" {
			continue
		}
		i := strings.Index(ra, "-")
		if i < 0 {
			return nil, errors.New("invalid range")
		}
		start, end := strings.TrimSpace(ra[:i]), strings.TrimSpace(ra[i+1:])
		var r httpRange
		if start == "" {
			// If no start is specified, end specifies the
			// range start relative to the end of the file.
			i, err := strconv.ParseInt(end, 10, 64)
			if err != nil {
				return nil, errors.New("invalid range")
			}
			if i > size {
				i = size
			}
			r.start = size - i
			r.length = size - r.start
		} else {
			i, err := strconv.ParseInt(start, 10, 64)
			if err != nil || i < 0 {
				return nil, errors.New("invalid range")
			}
			if i >= size {
				// If the range begins after the size of the content,
				// then it does not overlap.
				noOverlap = true
				continue
			}
			r.start = i
			if end == "" {
				// If no end is specified, range extends to end of the file.
				r.length = size - r.start
			} else {
				i, err := strconv.ParseInt(end, 10, 64)
				if err != nil || r.start > i {
					return nil, errors.New("invalid range")
				}
				if i >= size {
					i = size - 1
				}
				r.length = i - r.start + 1
			}
		}
		ranges = append(ranges, r)
	}
	if noOverlap && len(ranges) == 0 {
		// The specified ranges did not overlap with the content.
		return nil, errNoOverlap
	}
	return ranges, nil
}
func sumRangesSize(ranges []httpRange) (size int64) {
	for _, ra := range ranges {
		size += ra.length
	}
	return
}

// countingWriter counts how many bytes have been written to it.
type countingWriter int64

func (w *countingWriter) Write(p []byte) (n int, err error) {
	*w += countingWriter(len(p))
	return len(p), nil
}

// rangesMIMESize returns the number of bytes it takes to encode the
// provided ranges as a multipart response.
func rangesMIMESize(ranges []httpRange, contentType string, contentSize int64) (encSize int64) {
	var w countingWriter
	mw := multipart.NewWriter(&w)
	for _, ra := range ranges {
		mw.CreatePart(ra.mimeHeader(contentType, contentSize))
		encSize += ra.length
	}
	mw.Close()
	encSize += int64(w)
	return
}
