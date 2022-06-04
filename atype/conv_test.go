package atype_test

import (
	"github.com/hi-iwi/AaGo/atype"
	"testing"
)

func TestNotEmpty(t *testing.T) {
	if atype.NotEmpty(false) != false {
		t.Errorf("atype.NotEmpty(false) == %s", atype.String(atype.NotEmpty(false)))
	}
	if atype.NotEmpty(true) != true {
		t.Errorf("atype.NotEmpty(true) == %s", atype.String(atype.NotEmpty(true)))
	}
	if atype.NotEmpty(byte(0)) != false {
		t.Errorf("atype.NotEmpty(byte(0)) == %s", atype.String(atype.NotEmpty(byte(0))))
	} else {
		t.Logf("[warn] atype.NotEmpty(byte(0)) == %s", atype.String(atype.NotEmpty(byte(0))))

	}
	if atype.NotEmpty(byte('0')) != true {
		t.Errorf("atype.NotEmpty(byte('0')) == %s", atype.String(atype.NotEmpty(byte('0'))))
	} else {
		t.Logf("[warn] atype.NotEmpty(byte('0')) == %s", atype.String(atype.NotEmpty(byte('0'))))
	}
	if atype.NotEmpty(rune(0)) != false {
		t.Errorf("atype.NotEmpty(rune(0)) == %s",  atype.String(atype.NotEmpty(rune(0))))
	}
	if atype.NotEmpty("") != false {
		t.Errorf("atype.NotEmpty(\"\") == %s",  atype.String(atype.NotEmpty("")))
	}
	if atype.NotEmpty([]byte{}) != false {
		t.Errorf("atype.NotEmpty([]byte{}) == %s",  atype.String(atype.NotEmpty([]byte{})))
	}
	if atype.NotEmpty([]byte{'0'}) != true {
		t.Errorf("atype.NotEmpty([]byte{'0'}) == %s",  atype.String(atype.NotEmpty([]byte{'0'})))
	}

	if atype.NotEmpty(-1) != true {
		t.Errorf("atype.NotEmpty(-1) == %s",  atype.String(atype.NotEmpty(-1)))
	}
	if atype.NotEmpty(0) != false {
		t.Errorf("atype.NotEmpty(0) == %s",  atype.String(atype.NotEmpty(0)))
	}
	if atype.NotEmpty(1) != true {
		t.Errorf("atype.NotEmpty(1) == %s",  atype.String(atype.NotEmpty(1)))
	}
}
func TestString(t *testing.T) {
	if atype.String(false) != "false" {
		t.Errorf("bool(false) ==> string(%s)",  atype.String(false))
	}
	if atype.String(true) != "true" {
		t.Errorf("bool(true) ==> string(%s)",  atype.String(true))
	}

	// byte is a built-in alias of uint8, Name('A') returns "97"

	if atype.String('A') != "65" {
		t.Errorf("A ==> string(%s)",  atype.String('A'))
	}

	if atype.String(byte('A')) != "65" {
		t.Errorf("byte(A) ==> string(%s)",  atype.String(byte('A')))
	}

	if atype.String(atype.Abyte('A')) != "A" {
		t.Errorf("atype.Abyte(A) ==> string(%s)",  atype.String(atype.Abyte('A')))
	}

	if atype.String([]byte{'A', 'a'}) != "App" {
		t.Errorf("[]byte(App) ==> string(%s)",  atype.String([]byte{'A', 'a'}))
	}

	if atype.String("Iwi") != "Iwi" {
		t.Errorf("string(Iwi) ==> string(%s)",  atype.String("Iwi"))
	}

	if atype.String(int8(100)) != "100" {
		t.Errorf("int8(100) ==> string(%s)",  atype.String(int8(100)))
	}

	if atype.String(int16(100)) != "100" {
		t.Errorf("int16(100) ==> string(%s)",  atype.String(int16(100)))
	}

	if atype.String(int32(100)) != "100" {
		t.Errorf("int32(100) ==> string(%s)",  atype.String(int32(100)))
	}
	if atype.String(100) != "100" {
		t.Errorf("int(100) ==> string(%s)",  atype.String(100))
	}

	if atype.String(int64(100)) != "100" {
		t.Errorf("int64(100) ==> string(%s)",  atype.String(int64(100)))
	}

	if atype.String(float32(100.0)) != "100" {
		t.Errorf("float32(100.0) ==> string(%s)",  atype.String(float32(100.0)))
	} else {
		t.Logf("[warn] float32(100.0) ==> string(%s)",  atype.String(float32(100.0)))
	}
	if atype.String(float64(100.0)) != "100" {
		t.Errorf("float64(100.0) ==> string(%s)",  atype.String(float64(100.0)))
	} else {
		t.Logf("[warn] float32(100.0) ==> string(%s)",  atype.String(float32(100.0)))
	}

	b := 234242342342423.3
	if atype.String(b) != "234242342342423.3" {
		t.Errorf("float64(%f) ==> string(%s)", b,  atype.String(b))
	}

}
func TestatypeConv(t *testing.T) {

}
