package code_test

import (
	"testing"
)

func TestUint64ShortUUID(t *testing.T) {
	id := Uint64ShortUUID()
	_, ets, seq := ShortUuidToTime(id)

	testingID := ToUint64ShortUUID(ets, seq)
	if id != testingID {
		t.Errorf("code.ShortUuidToTime error: %d != %d", id, testingID)
	}
}
