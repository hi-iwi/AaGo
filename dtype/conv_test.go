package dtype_test

import (
	"testing"
)

func TestNotEmpty(t *testing.T) {
	if NotEmpty(false) != false {
		t.Errorf("dtype.NotEmpty(false) == %s", String(NotEmpty(false)))
	}
	if NotEmpty(true) != true {
		t.Errorf("dtype.NotEmpty(true) == %s", String(NotEmpty(true)))
	}
	if NotEmpty(byte(0)) != false {
		t.Errorf("dtype.NotEmpty(byte(0)) == %s", String(NotEmpty(byte(0))))
	} else {
		t.Logf("[warn] dtype.NotEmpty(byte(0)) == %s", String(NotEmpty(byte(0))))

	}
	if NotEmpty(byte('0')) != true {
		t.Errorf("dtype.NotEmpty(byte('0')) == %s", String(NotEmpty(byte('0'))))
	} else {
		t.Logf("[warn] dtype.NotEmpty(byte('0')) == %s", String(NotEmpty(byte('0'))))
	}
	if NotEmpty(rune(0)) != false {
		t.Errorf("dtype.NotEmpty(rune(0)) == %s", String(NotEmpty(rune(0))))
	}
	if NotEmpty("") != false {
		t.Errorf("dtype.NotEmpty(\"\") == %s", String(NotEmpty("")))
	}
	if NotEmpty([]byte{}) != false {
		t.Errorf("dtype.NotEmpty([]byte{}) == %s", String(NotEmpty([]byte{})))
	}
	if NotEmpty([]byte{'0'}) != true {
		t.Errorf("dtype.NotEmpty([]byte{'0'}) == %s", String(NotEmpty([]byte{'0'})))
	}

	if NotEmpty(-1) != true {
		t.Errorf("dtype.NotEmpty(-1) == %s", String(NotEmpty(-1)))
	}
	if NotEmpty(0) != false {
		t.Errorf("dtype.NotEmpty(0) == %s", String(NotEmpty(0)))
	}
	if NotEmpty(1) != true {
		t.Errorf("dtype.NotEmpty(1) == %s", String(NotEmpty(1)))
	}
}
func TestString(t *testing.T) {
	if String(false) != "false" {
		t.Errorf("bool(false) ==> string(%s)", String(false))
	}
	if String(true) != "true" {
		t.Errorf("bool(true) ==> string(%s)", String(true))
	}

	// byte is a built-in alias of uint8, Name('A') returns "97"

	if String('A') != "65" {
		t.Errorf("A ==> string(%s)", String('A'))
	}

	if String(byte('A')) != "65" {
		t.Errorf("byte(A) ==> string(%s)", String(byte('A')))
	}

	if String(Dbyte('A')) != "A" {
		t.Errorf("dtype.Dbyte(A) ==> string(%s)", String(Dbyte('A')))
	}

	if String([]byte{'A', 'a'}) != "Aa" {
		t.Errorf("[]byte(Aa) ==> string(%s)", String([]byte{'A', 'a'}))
	}

	if String("Aario") != "Aario" {
		t.Errorf("string(Aario) ==> string(%s)", String("Aario"))
	}

	if String(int8(100)) != "100" {
		t.Errorf("int8(100) ==> string(%s)", String(int8(100)))
	}

	if String(int16(100)) != "100" {
		t.Errorf("int16(100) ==> string(%s)", String(int16(100)))
	}

	if String(int32(100)) != "100" {
		t.Errorf("int32(100) ==> string(%s)", String(int32(100)))
	}
	if String(100) != "100" {
		t.Errorf("int(100) ==> string(%s)", String(100))
	}

	if String(int64(100)) != "100" {
		t.Errorf("int64(100) ==> string(%s)", String(int64(100)))
	}

	if String(float32(100.0)) != "100" {
		t.Errorf("float32(100.0) ==> string(%s)", String(float32(100.0)))
	} else {
		t.Logf("[warn] float32(100.0) ==> string(%s)", String(float32(100.0)))
	}
	if String(float64(100.0)) != "100" {
		t.Errorf("float64(100.0) ==> string(%s)", String(float64(100.0)))
	} else {
		t.Logf("[warn] float32(100.0) ==> string(%s)", String(float32(100.0)))
	}

	b := 234242342342423.3
	if String(b) != "234242342342423.3" {
		t.Errorf("float64(%f) ==> string(%s)", b, String(b))
	}

}
func TestDtypeConv(t *testing.T) {

}
