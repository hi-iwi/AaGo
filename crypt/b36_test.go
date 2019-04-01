package crypt_test

import (
	"path"
	"testing"
	"time"

	"github.com/luexu/AaGo/crypt"
)

func isEnigmaOk(b byte) bool {
	return b != '?' && (b < '0' || (b > '9' && b < 'a') || b > 'z')
}
func TestEnigma(t *testing.T) {
	var (
		ln1 string
		b1  byte
	)

	for i := 0; i < len(crypt.Base36BinStr); i++ {
		ln1 = "< "
		for v := 0; v < 37; v++ {
			b1 = crypt.Enigma(v, crypt.Base36BinStr[i])
			if isEnigmaOk(b1) {
				t.Errorf("crypt.Enigma(%d, '%s') ====> %s", v, string(crypt.Base36BinStr[i]), string(b1))
				return
			}
			ln1 += string(b1)
		}
		t.Log(ln1)
	}
}

func TestEnigmaNoPs2(t *testing.T) {
	var (
		ln1 string
		b1  byte
	)

	ln1 = "< "
	for v := 0; v < 37; v++ {
		b1 = crypt.Enigma(v)
		if isEnigmaOk(b1) {
			t.Errorf("crypt.Enigma(%d) ====> %s", v, string(b1))
			return
		}
		ln1 += string(b1)
	}
	t.Log(ln1)
}

func TestEnigmaNoPsA(t *testing.T) {
	var (
		ln1 string
		b1  byte
	)

	ln1 = "< "
	for v := 0; v < 37; v++ {
		b1 = crypt.Enigma(v, 'y')
		if isEnigmaOk(b1) {
			t.Errorf("crypt.Enigma(%d, 'y') ====> %s", v, string(b1))
			return
		}
		ln1 += string(b1)
	}
	t.Log(ln1)
}

func TestEnigmaNoPsB(t *testing.T) {
	var (
		ln1 string
		b1  byte
	)

	ln1 = "< "
	for v := 0; v < 37; v++ {
		b1 = crypt.Enigma(v, true)
		if isEnigmaOk(b1) {
			t.Errorf("crypt.Enigma(%d, true) ====> %s", v, string(b1))
			return
		}
		ln1 += string(b1)
	}
	t.Log(ln1)
}

func TestEnigmaReverse(t *testing.T) {
	var (
		ln2 string
		b2  byte
	)

	for i := 0; i < len(crypt.Base36BinStr); i++ {
		ln2 = "> "
		for v := 0; v < 37; v++ {
			b2 = crypt.Enigma(v, crypt.Base36BinStr[i], true)
			if isEnigmaOk(b2) {
				t.Errorf("crypt.Enigma(%d, '%s', true) ====> %s", v, string(crypt.Base36BinStr[i]), string(b2))
				return
			}
			ln2 += string(b2)
		}
		t.Log(ln2)
	}
}

func TestB36(t *testing.T) {
	s := ""
	for i := 1; i < 10000; i++ {
		s += " " + crypt.B36(i)
	}
	t.Log(s)
}

func dstPath(workpath string) string {
	now := time.Now()
	y, mo, d := now.Date()
	month := int(mo)
	xm := month + (now.Second()%3)*12
	dir1 := crypt.B36(y-2018, 'y', true) + crypt.B36(xm, 'm')
	dir2 := crypt.B36(d, 'd', true) + crypt.B36(now.Hour(), 'h', true)
	fileMainName := crypt.Pad(crypt.B36(now.Minute(), 'm', true), 2) + crypt.Pad(crypt.B36(now.Second(), 's', true), 2) + crypt.Pad(crypt.B36(now.Nanosecond()/1e3, 'n', true), 4)

	fp := path.Join(workpath, dir1, dir2, fileMainName)
	return fp
	//return conf.App.Path.ImgDstPath(dirname, fp)
}

func TestX(t *testing.T) {
	t.Log(dstPath("LOVE"))
}
