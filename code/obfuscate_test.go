package code_test

import (
	"github.com/hi-iwi/AaGo/code"
	"testing"
	"time"
)

func TestObfuscateNumber(t *testing.T) {
	id := uint64(20277781505)
	eu := code.ObfuscateNumber(id, 'b', code.ObfuscateBase)
	deid := code.DeobfuscateNumber(eu, 'b', code.ObfuscateBase)
	if id != deid {
		t.Errorf("code.ObfuscateNumber(%d) ==> %s, code.DeobfuscateNumber(%s) ==> %d", id, eu, eu, deid)
	}

	id = uint64(1)
	eu = code.ObfuscateNumber(id, '6', code.ObfuscateBase)
	deid = code.DeobfuscateNumber(eu, '6', code.ObfuscateBase)
	if id != deid {
		t.Errorf("code.ObfuscateNumber(%d) ==> %s, code.DeobfuscateNumber(%s) ==> %d", id, eu, eu, deid)
	}

	ts := time.Now().Unix()
	ets := code.ObfuscateNumber(uint64(ts), 0, code.ObfuscateBase)
	dets := code.DeobfuscateNumber(ets, 0, code.ObfuscateBase)
	if ts != int64(dets) {
		t.Errorf("code.ObfuscateNumber(%d) ==> %s, code.DeobfuscateNumber(%s) ==> %d", ts, ets, ets, dets)
	}

}
func TestObfuscateBytes(t *testing.T) {
	str := "base36"
	obStr, _ := code.ObfuscateBytes([]byte(str), 'a', code.ObfuscateBase)
	deStr, _ := code.DeobfuscateBytes([]byte(obStr), 'a', code.ObfuscateBase)
	if str != string(deStr) {
		t.Errorf("code.ObfuscateBytes(%s) ==> %s, code.DeobfuscateBytes(%s) ==> %s", str, obStr, obStr, deStr)
	}
	var obfuscateBase = []byte{'6', 'h', 'n', '1', '3', 'z', 's', 'm', 'c', 'd', 'o', 'f', 'i', 'j', '2', 'l', 'q', '0', '4', 'y', 'v', '5', '7', 'b', '9', 'a', 'x', 'g', 'w', 'k', 'u', '8', 't', 'p', 'e', 'r', '.', 'L'}
	str = "Luexu.com"
	obStr, _ = code.ObfuscateBytes([]byte(str), 'x', obfuscateBase)
	deStr, _ = code.DeobfuscateBytes([]byte(obStr), 'x', obfuscateBase)
	if str != string(deStr) {
		t.Errorf("code.ObfuscateBytes(%s) ==> %s, code.DeobfuscateBytes(%s) ==> %s", str, obStr, obStr, deStr)
	}
}
