package code_test

import (
	"testing"
	"time"
)

func TestObfuscateNumber(t *testing.T) {
	id := uint64(20277781505)
	eu := ObfuscateNumber(id, 'b', ObfuscateBase)
	deid := DeobfuscateNumber(eu, 'b', ObfuscateBase)
	if id != deid {
		t.Errorf("code.ObfuscateNumber(%d) ==> %s, code.DeobfuscateNumber(%s) ==> %d", id, eu, eu, deid)
	}

	id = uint64(1)
	eu = ObfuscateNumber(id, '6', ObfuscateBase)
	deid = DeobfuscateNumber(eu, '6', ObfuscateBase)
	if id != deid {
		t.Errorf("code.ObfuscateNumber(%d) ==> %s, code.DeobfuscateNumber(%s) ==> %d", id, eu, eu, deid)
	}

	ts := time.Now().Unix()
	ets := ObfuscateNumber(uint64(ts), 0, ObfuscateBase)
	dets := DeobfuscateNumber(ets, 0, ObfuscateBase)
	if ts != int64(dets) {
		t.Errorf("code.ObfuscateNumber(%d) ==> %s, code.DeobfuscateNumber(%s) ==> %d", ts, ets, ets, dets)
	}

}
func TestObfuscateBytes(t *testing.T) {
	str := "base36"
	obStr, _ := ObfuscateBytes([]byte(str), 'a', ObfuscateBase)
	deStr, _ := DeobfuscateBytes([]byte(obStr), 'a', ObfuscateBase)
	if str != string(deStr) {
		t.Errorf("code.ObfuscateBytes(%s) ==> %s, code.DeobfuscateBytes(%s) ==> %s", str, obStr, obStr, deStr)
	}
	var obfuscateBase = []byte{'6', 'h', 'n', '1', '3', 'z', 's', 'm', 'c', 'd', 'o', 'f', 'i', 'j', '2', 'l', 'q', '0', '4', 'y', 'v', '5', '7', 'b', '9', 'a', 'x', 'g', 'w', 'k', 'u', '8', 't', 'p', 'e', 'r', '.', 'L'}
	str = "Luexu.com"
	obStr, _ = ObfuscateBytes([]byte(str), 'x', obfuscateBase)
	deStr, _ = DeobfuscateBytes([]byte(obStr), 'x', obfuscateBase)
	if str != string(deStr) {
		t.Errorf("code.ObfuscateBytes(%s) ==> %s, code.DeobfuscateBytes(%s) ==> %s", str, obStr, obStr, deStr)
	}
}
