package crypt

import (
	"testing"
)

func TestSha1(t *testing.T) {
	h := "c2a9ca9e17605bdad0d1d7c6985a69e6e5aa34f0"
	if Sha1("Aario") != h {
		t.Errorf("crypt.Sha1(Aario) == %s is wrong", Sha1("Aario"))
	}
	if Sha1("Aario", 4) != "c2a9" {
		t.Errorf("crypt.Sha1(Aario, 4) == %s is wrong", Sha1("Aario", 4))
	}
	if Sha1("Aario", len(h)) != h {
		t.Errorf("crypt.Sha1(Aario, len(h)) == %s is wrong", Sha1("Aario"))
	}

	if Sha1("Aario", len(h)-1) != h[0:(len(h)-1)] {
		t.Errorf("crypt.Sha1(Aario, len(h) - 1) == %s is wrong", Sha1("Aario", len(h)-1))
	}

	if Sha1("Aario", len(h)+3) != "000"+h {
		t.Errorf("crypt.Sha1(Aario, len(h)+3) == %s is wrong", Sha1("Aario", len(h)+3))
	}
}
