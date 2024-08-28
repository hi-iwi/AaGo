package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"regexp"
	"slices"
	"sort"
	"strings"
)

const tab = "    "

func readJsonFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	result := make([]byte, 0)
	var firstBraceHandled bool
	for scan.Scan() {
		b := bytes.Trim(scan.Bytes(), " ")
		// @TODO 暂时只支持单行 // 方式注释

		if len(b) == 0 || (len(b) > 1 && string(b[0:2]) == "//") {
			continue
		}
		if !firstBraceHandled {
			i := bytes.Index(b, []byte("{"))
			if i > -1 {
				b = b[i:]
				firstBraceHandled = true
			}
		}
		key, value, ok := bytes.Cut(b, []byte(":"))
		key = bytes.Trim(key, " ")
		if ok && (len(key) < 3 || (key[0] != '"' && key[0] != '\'')) {
			b = append([]byte{'"'}, key...)
			b = append(b, '"', ':')
			b = append(b, value...)
		}

		result = append(result, b...)
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`,[\s\r\n]*([}\]])`)
	result = re.ReplaceAll(result, []byte("$1"))

	return result, nil
}
func w(f *os.File, format string, args ...any) {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	_, err := f.WriteString(format)
	if err != nil {
		panic(err)
	}
}
func appendMimes(mimes []string, mime ...string) []string {
	for _, mi := range mime {
		var exists bool
		for _, m := range mimes {
			if mi == m {
				exists = true
			}
		}
		if !exists {
			mimes = append(mimes, mi)
		}
	}
	return mimes
}
func main() {
	const jsonFile = "../aenum/filetype.jsonp"

	b, err := readJsonFile(jsonFile)
	if err != nil {
		panic(err)
	}
	var raw map[string]map[string][]any
	err = json.Unmarshal(b, &raw)
	if err != nil {
		panic(err.Error() + "\n" + string(b))
	}
	mimes := make([]string, 0)
	enums := make([][2]any, 0)
	types := make(map[string]map[string][]string)
	for category, r := range raw {
		types[category] = make(map[string][]string)
		for k, arr := range r {
			if len(arr) < 3 {
				panic(fmt.Sprintf("invalid %s.%s %v", category, k, arr))
			}
			id := int(arr[0].(float64))

			if len(enums) == 0 {
				enums = append(enums, [2]any{k, id})
			} else {
				var i int
				var prev [2]any
				var found bool
				for i, prev = range enums {
					if prev[1].(int) > id {
						found = true
						break
					}
				}
				if !found {
					i++
				}
				enums = slices.Insert(enums, i, [2]any{k, id})
			}
			ext := arr[1].(string)
			standardMime := arr[2].(string)
			types[category][k] = []string{ext, standardMime}
			mimes = appendMimes(mimes, ext, standardMime)
			if len(arr) > 3 {
				others := arr[3].([]any)
				for _, m := range others {
					types[category][k] = append(types[category][k], m.(string))
					mimes = appendMimes(mimes, m.(string))
				}
			}

		}
	}

	buildFileTypeGo(enums, types)
	buildFileTypeJS(enums, types, mimes)
}

func buildFileTypeGo(enums [][2]any, types map[string]map[string][]string) {
	const dstFile = "../aenum/filetype_readonly.go"
	f, err := os.OpenFile(dstFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w(f, "package aenum\n")
	w(f, `import "strings"`+"\n\n")

	//w(f, "type FileType uint16\n")
	w(f, "const (\n")
	w(f, "%sUnknownType FileType = 0\n", tab)
	for _, enum := range enums {
		w(f, fmt.Sprintf("%s%-11s FileType = %v\n", tab, enum[0], enum[1]))
	}
	w(f, ")\n")

	for t, d := range types {
		w(f, fmt.Sprintf("var %sTypes = map[FileType][]string{\n", t))
		ks := make([]string, 0, len(d))
		for k, _ := range d {
			ks = append(ks, k)
		}
		sort.Strings(ks)
		for _, k := range ks {
			w(f, fmt.Sprintf("%s%-11s : {", tab, k))
			for i, m := range d[k] {
				if i > 0 {
					w(f, ", ")
				}
				w(f, `"`+m+`"`)
			}
			w(f, "},\n")
		}

		w(f, "}\n")
	}
	// NewXxxType()
	for t, _ := range types {
		w(f, fmt.Sprintf("func New%sType(mime string) (FileType, bool) {return ParseFileType(mime, %sTypes)}\n", t, t))
	}

	// Content Type
	w(f, "func (t FileType) ContentType() string {\n")
	for t, _ := range types {
		w(f, fmt.Sprintf("%sif d, ok := %sTypes[t]; ok {return d[1]}\n", tab, t))
	}
	w(f, fmt.Sprintf("%sreturn \"\"\n", tab))
	w(f, "}\n")

	// ext
	w(f, "func (t FileType) Ext() string {\n")
	for t, _ := range types {
		w(f, fmt.Sprintf("%sif d, ok := %sTypes[t]; ok {return d[0]}\n", tab, t))
	}
	w(f, fmt.Sprintf("%sreturn \"\"\n", tab))
	w(f, "}\n")

	// filename
	w(f, `func (t FileType) Name() string {return strings.TrimPrefix(t.Ext(), ".")}`+"\n")
}
func buildFileTypeJS(enums [][2]any, types map[string]map[string][]string, mimes []string) {
	var dstFile = "./f_oss_filetype_readonly.js"
	const aaJsFile = "../../../../project/xixi/deploy/asset_src/lib_dev/aa-js/src/f_oss_filetype_readonly.js"
	fi, err := os.Stat(path.Dir(aaJsFile))
	if err == nil && fi.IsDir() {
		dstFile = aaJsFile
	}

	f, err := os.OpenFile(dstFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	w(f, "/** @note this is an auto-generated file, do not modify it! */\n\n")
	w(f, "/** @typedef {")
	w(f, `"`+strings.Join(mimes, `"|"`)+`"`)
	w(f, "} AaFileTypeMime */\n\n")

	w(f, "class AaFileType {\n")
	w(f, "%s/** @enum */\n", tab)
	w(f, "%sstatic Enum={\n", tab)
	w(f, "%s%sUnknownType: 0,\n", tab, tab)
	for _, enum := range enums {
		w(f, fmt.Sprintf("%s%s%-11s : %v,\n", tab, tab, enum[0], enum[1]))
	}
	w(f, "%s}\n", tab)

	w(f, "%sstatic Mimes = {\n", tab)
	for t, d := range types {
		w(f, fmt.Sprintf("%s%s%s : {\n", tab, tab, t))
		ks := make([]string, 0, len(d))
		for k, _ := range d {
			ks = append(ks, k)
		}
		sort.Strings(ks)
		for _, k := range ks {
			w(f, fmt.Sprintf("%s%s%s%-11s : [", tab, tab, tab, k))
			for i, m := range d[k] {
				if i > 0 {
					w(f, ", ")
				}
				w(f, `"`+m+`"`)
			}
			w(f, "],\n")
		}

		w(f, "%s%s},\n", tab, tab)
	}
	w(f, "%s}\n", tab)
	w(f, "%scontentType\n", tab)
	w(f, "%sext\n", tab)
	w(f, "%smimeType\n", tab)
	w(f, "%svalue\n", tab)
	// constructor
	w(f, "\n%s/**\n", tab)
	w(f, "%s * @param {AaFileTypeMime|number} mime\n", tab)
	w(f, "%s */\n", tab)
	w(f, "%sconstructor(mime){", tab)
	w(f, `
		this.value = AaFileType.Enum.UnknownType
		for(const [type, cv] of Object.entries(AaFileType.Mimes)){
			for(const [v,mimes] of Object.entries(cv)){
				 if(mime ===  AaFileType[v] || mimes.includes(mime)){
					this.contentType = mimes[1]
					this.ext = mimes[0]
					this.mimeType = type
					this.value = AaFileType[v]
					return
				}
			}
		}
`)
	w(f, "%s}\n", tab)
	for t, _ := range types {
		w(f, `%sis%s(){return this.mimeType === "%s"}`+"\n", tab, t, t)
	}
	w(f, "%stoJSON(){return this.value}\n", tab)
	w(f, "%svalueOf(){return this.value}\n", tab)
	w(f, "}\n")
}
