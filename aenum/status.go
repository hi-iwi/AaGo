package aenum

import "strconv"

type Status int8

const (
	SysRevoked Status = -128 // 已注销，系统删除（可能审核失败）
	Expired    Status = -121 // 已失效/已过期
	Deleted    Status = -100 // 用户已删除，谁都不可见
	Failed     Status = -20  // censor failed 审核失败，让用户可以修改，仅用户可见  --> 无论之前设置什么，这里统一一个失败状态，修改的时候再设置
	Pending1   Status = -10  // censor 1 审核1通过 --> 进入阶段性审核流程，就会阻却用户修改权
	Pending2   Status = -9
	Pending3   Status = -8
	Pending4   Status = -7
	Pending5   Status = -6
	Pending6   Status = -5
	Pending7   Status = -4
	Pending8   Status = -3
	Pending9   Status = -2
	Pending    Status = -1  // 审核中，不展示
	Created    Status = 0   // 新创建，审核中，但是显示
	Passed     Status = 1   // 审核通过
	SysLocked  Status = 127 // 系统已锁定，用户不得再修改；PS：允许删除，可执行解绑软删除，或硬删除，具体情况具体对待

)

var (
	PassedRange     = [2]Status{Passed, SysLocked}
	PendingRange    = [2]Status{Pending1, Created} // 等待审核的区间
	MeReadableRange = [2]Status{Failed, SysLocked} // 会显示在我列表的区间
)

func NewStatus(sts int8) (Status, bool) {
	s := Status(sts)
	ok := s == SysRevoked || s == Expired || s == Deleted || s == Failed
	ok = ok || (s >= Pending1 && s <= Passed) || s == SysLocked
	return s, ok
}
func (s Status) Int8() int8     { return int8(s) }
func (s Status) String() string { return strconv.Itoa(int(s)) }
func (s Status) Is(x int8) bool { return s.Int8() == x }
func (s Status) In(a ...Status) bool {
	for _, sts := range a {
		if sts == s {
			return true
		}
	}
	return false
}

// 显示与不显示的零界点
func (s Status) IsOk() bool {
	return s >= Created
}
func (s Status) IsPending() bool {
	return s >= PendingRange[0] && s <= PendingRange[1]
}

// 是否审核通过
func (s Status) IsPassed() bool {
	return s >= PassedRange[0] && s <= PassedRange[1]
}

// 用户自己是否可见
func (s Status) MeReadable() bool {
	return s >= MeReadableRange[0] && s <= MeReadableRange[1]
}

// 用户是否可以修改、删除
func (s Status) Modifiable() bool {
	return s.In(Failed, Pending, Created, Passed)
}
