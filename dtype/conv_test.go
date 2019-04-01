package dtype_test

import (
	"testing"

	"github.com/luexu/AaGo/dtype"
)

func TestNotEmpty(t *testing.T) {
	if dtype.NotEmpty(false) != false {
		t.Errorf("dtype.NotEmpty(false) == %s", dtype.String(dtype.NotEmpty(false)))
	}
	if dtype.NotEmpty(true) != true {
		t.Errorf("dtype.NotEmpty(true) == %s", dtype.String(dtype.NotEmpty(true)))
	}
	if dtype.NotEmpty(byte(0)) != false {
		t.Errorf("dtype.NotEmpty(byte(0)) == %s", dtype.String(dtype.NotEmpty(byte(0))))
	} else {
		t.Logf("[warn] dtype.NotEmpty(byte(0)) == %s", dtype.String(dtype.NotEmpty(byte(0))))

	}
	if dtype.NotEmpty(byte('0')) != true {
		t.Errorf("dtype.NotEmpty(byte('0')) == %s", dtype.String(dtype.NotEmpty(byte('0'))))
	} else {
		t.Logf("[warn] dtype.NotEmpty(byte('0')) == %s", dtype.String(dtype.NotEmpty(byte('0'))))
	}
	if dtype.NotEmpty(rune(0)) != false {
		t.Errorf("dtype.NotEmpty(rune(0)) == %s", dtype.String(dtype.NotEmpty(rune(0))))
	}
	if dtype.NotEmpty("") != false {
		t.Errorf("dtype.NotEmpty(\"\") == %s", dtype.String(dtype.NotEmpty("")))
	}
	if dtype.NotEmpty([]byte{}) != false {
		t.Errorf("dtype.NotEmpty([]byte{}) == %s", dtype.String(dtype.NotEmpty([]byte{})))
	}
	if dtype.NotEmpty([]byte{'0'}) != true {
		t.Errorf("dtype.NotEmpty([]byte{'0'}) == %s", dtype.String(dtype.NotEmpty([]byte{'0'})))
	}

	if dtype.NotEmpty(-1) != true {
		t.Errorf("dtype.NotEmpty(-1) == %s", dtype.String(dtype.NotEmpty(-1)))
	}
	if dtype.NotEmpty(0) != false {
		t.Errorf("dtype.NotEmpty(0) == %s", dtype.String(dtype.NotEmpty(0)))
	}
	if dtype.NotEmpty(1) != true {
		t.Errorf("dtype.NotEmpty(1) == %s", dtype.String(dtype.NotEmpty(1)))
	}
}
func TestString(t *testing.T) {
	if dtype.String(false) != "false" {
		t.Errorf("bool(false) ==> string(%s)", dtype.String(false))
	}
	if dtype.String(true) != "true" {
		t.Errorf("bool(true) ==> string(%s)", dtype.String(true))
	}

	// 注意：'A' == uint8(65)    byte('A') == uint8(65)

	if dtype.String('A') != "65" {
		t.Errorf("A ==> string(%s)", dtype.String('A'))
	} else {
		t.Logf("[warn] A is int32(A) ==> string(%s)", dtype.String('A'))
	}

	if dtype.String(byte('A')) != "A" {
		t.Errorf("byte(A) ==> string(%s)", dtype.String(byte('A')))
	}
	if dtype.String([]byte{'A', 'a'}) != "Aa" {
		t.Errorf("[]byte(Aa) ==> string(%s)", dtype.String([]byte{'A', 'a'}))
	}

	if dtype.String("Aario") != "Aario" {
		t.Errorf("string(Aario) ==> string(%s)", dtype.String("Aario"))
	}

	if dtype.String(int8(100)) != "100" {
		t.Errorf("int8(100) ==> string(%s)", dtype.String(int8(100)))
	}

	if dtype.String(int16(100)) != "100" {
		t.Errorf("int16(100) ==> string(%s)", dtype.String(int16(100)))
	}

	if dtype.String(int32(100)) != "100" {
		t.Errorf("int32(100) ==> string(%s)", dtype.String(int32(100)))
	}
	if dtype.String(100) != "100" {
		t.Errorf("int(100) ==> string(%s)", dtype.String(100))
	}

	if dtype.String(int64(100)) != "100" {
		t.Errorf("int64(100) ==> string(%s)", dtype.String(int64(100)))
	}

	if dtype.String(float32(100.0)) != "100" {
		t.Errorf("float32(100.0) ==> string(%s)", dtype.String(float32(100.0)))
	} else {
		t.Logf("[warn] float32(100.0) ==> string(%s)", dtype.String(float32(100.0)))
	}
	if dtype.String(float64(100.0)) != "100" {
		t.Errorf("float64(100.0) ==> string(%s)", dtype.String(float64(100.0)))
	} else {
		t.Logf("[warn] float32(100.0) ==> string(%s)", dtype.String(float32(100.0)))
	}

	b := 234242342342423.3
	if dtype.String(b) != "234242342342423.3" {
		t.Errorf("float64(%f) ==> string(%s)", b, dtype.String(b))
	}

}
func TestDtypeConv(t *testing.T) {

}
