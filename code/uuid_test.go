package code_test

import (
	"github.com/hi-iwi/AaGo/code"
	"testing"
)

func TestUint64ShortUUID(t *testing.T) {
	id := code.Uint64ShortUUID()
	_, ets, seq := code.ShortUuidToTime(id)

	testingID := code.ToUint64ShortUUID(ets, seq)
	if id != testingID {
		t.Errorf("code.ShortUuidToTime error: %d != %d", id, testingID)
	}
}
