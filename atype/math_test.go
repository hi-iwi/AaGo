package atype_test

import (
	"github.com/hi-iwi/AaGo/atype"
	"math"
	"math/rand"
	"testing"
)

func TestPercent(t *testing.T) {
	m := atype.YuanX(1000.0)
	if m.MulRound(atype.HundredPercent(0.6)) != atype.YuanX(6) {
		t.Errorf("atype.MoneyYuan(1000).MulPercent(0.6) != atype.YuanX(6)")
	}
	if m.MulRound(atype.HundredPercent(30)) != atype.YuanX(300) {
		t.Errorf("atype.MoneyYuan(1000).MulPercent(30) != atype.YuanX(300)")
	}
}
func TestDecimal(t *testing.T) {
	const n = 100
	for i := 0; i < n; i++ {
		a := math.Floor(rand.Float64()*10000*10000) / 10000 // 保留4位小数
		b := math.Floor(rand.Float64()*n*10000) / 10000     // 保留4位小数
		c := a * b
		x := atype.DecimalUnit(a).MulRound(atype.DecimalUnit(b))
		if x != atype.DecimalUnit(c) {
			t.Errorf("%f*%f=%f   %d", a, 10000.0, a*10000.0, int(a*10000.0))
			t.Errorf("%f*%f=%f   %d", b, 10000.0, b*10000.0, int(b*10000.0))
			t.Errorf("atype.Real %f*%f=%f  != %f (%d*%d=%d) error", a, b, c, x.Real(), atype.DecimalUnit(a), atype.DecimalUnit(b), x)
		}
	}

}
func TestMoney(t *testing.T) {

	//if atype.Decimal(0).FmtMantissa() != ".00" {
	//	t.Errorf("atype.Decimal().FmtMantissa() : %s != .00", atype.Decimal(0).FmtMantissa())
	//}
	//
	//m := atype.YuanX(188.8)
	//if m.Format(nil) != "188.8" {
	//	t.Errorf("atype.MoneyYuan(188.8).Format() : %s != 188.8", m.Format(nil))
	//}
	//m = atype.YuanX(199.9)
	//if m.Fmt() != "199.90" {
	//	t.Errorf("atype.YuanX(199.9).Fmt() : %s != 199.90", m.Fmt())
	//}
	//p := 7 * atype.Thousandth
	//s := atype.YuanX(2360).MulRound(p).Fmt()
	//if s != "16.52" {
	//	t.Errorf("2360*0.7%% : %s != 16.52", s)
	//}

	b := atype.Money(-234242342340503)

	//if b.FormatWhole(nil) != "-23424234234" {
	//	t.Errorf("money (%d).FormatWhole() ==> string(%s)", b, b.FormatWhole(nil))
	//}
	if b.Format(&atype.DecimalFormat{Scale: 1}) != "-23424234234.0" {
		t.Errorf("money (%d).Format(1) ==> string(%s)", b, b.Format(&atype.DecimalFormat{Scale: 1}))
	}
	//if b.Format(&atype.DecimalFormat{Scale: 2}) != "-23424234234.05" {
	//	t.Errorf("money (%d).Format(2) ==> string(%s)", b, b.Format(&atype.DecimalFormat{Scale: 2}))
	//}
	//if b.Format(&atype.DecimalFormat{Scale: 3}) != "-23424234234.050" {
	//	t.Errorf("money (%d).Format(3) ==> string(%s)", b, b.Format(&atype.DecimalFormat{Scale: 3}))
	//}
	//if b.Format(&atype.DecimalFormat{Scale: 4}) != "-23424234234.0503" {
	//	t.Errorf("money (%d).Format(4) ==> string(%s)", b, b.Format(&atype.DecimalFormat{Scale: 4}))
	//}

}
