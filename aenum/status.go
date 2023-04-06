package aenum

import (
	"strconv"
)

type Status int8 // status 字段禁止默认为0，防止意外情况

const (
	// 最简易的 开关
	InvalidPart = "invalid" // <0 关   off 是关键字。
	OkPart      = "ok"      // < MAXVALUE 开

	//NegTernaryPart  = "neg" // -1, LESS 0
	//ZeroTernaryPart = ""    // 0, LESS 1
	//PosTernaryPart  = ""    // 1, LESS MAXVALUE

	// 通过与否
	NotPassPart = "notpass" // < 1
	PassedPart  = "passed"  // <MAXVALUE

	// 无用户操作的状态，仅后台管理
	InvalidPartA = "invalid" // < Pending
	PendingPartA = "pending" // < 0
	ValidPartA   = "valid"   // < MAXVALUE

	PublicPartAs  = ValidPartA
	PendingPartAs = PendingPartA

	// 含用户操作的状态
	DeletedPartB   = "deleted"   // < -99 后期需要清理的数据
	NonPublicPartB = "nonpublic" // < Pending 仅用户自己可见
	PendingPartB   = "pending"   // < 0
	CreatedPartB   = "created"   // = 0 审核中，显示在公开列表（适用于白名单用户）
	PassedPartB    = "passed"    // >=1 公开列表显示

	PublicPartBs  = CreatedPartB + "," + PassedPartB                         // 公开列表
	VisiblePartBs = NonPublicPartB + "," + PendingPartB + "," + PublicPartBs // 仅用户可见 + 公开列表
	PendingPartBs = PendingPartB + "," + CreatedPartB
)

const (
	SysRevoked Status = -128 // 已注销，系统删除（可能审核失败）
	Expired    Status = -121 // 已失效/已过期，后面系统会自动删除
	Deleted    Status = -100 // 用户已删除，谁都不可见

	/********************* <=-100  属于需要删除的内容，方便分区删除整个分区 ************************/

	Failed      Status = -20 // censor failed 审核失败，让用户可以修改，仅用户可见  --> 无论之前设置什么，这里统一一个失败状态，修改的时候再设置
	Pending     Status = -10 // 审核中，用户保留修改权，可以重新提交审核
	Pending1    Status = -9  // censor 1 审核1通过 --> 进入阶段性审核流程，就会阻却用户修改权；发现异常，系统锁定也会进入这个状态
	Pending2    Status = -8
	Pending3    Status = -7
	Pending4    Status = -6
	Pending5    Status = -5
	Pending6    Status = -4
	Pending7    Status = -3
	Pending8    Status = -2
	Pending9    Status = -1
	Created     Status = 0   // 新创建，审核中，但是显示；用户保留修改权
	Passed      Status = 1   // 审核通过
	SysTakeover Status = 127 // 系统已锁定，用户不得再修改；PS：允许删除，可执行解绑软删除，或硬删除，具体情况具体对待

)

var (
	MeReadableRange = [2]Status{Failed, SysTakeover} // 会显示在我列表的区间
)

func NewStatus(sts int8) (Status, bool) {
	s := Status(sts)
	ok := s == SysRevoked || s == Expired || s == Deleted || s == Failed
	ok = ok || (s >= Pending && s <= Passed) || s == SysTakeover
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
func (s Status) IsOk() bool { return s >= Created }

// 是否审核通过
func (s Status) IsPassed() bool { return s >= Passed }

// Created 也是待审核状态
func (s Status) IsPending() bool { return s >= Pending && s <= Created }
func (s Status) IsFailed() bool  { return s <= Failed }

// After/Before 不包括； From/Within 包括
// FromPending 一般用户检测某个状态是否在审核中，或者已审核通过。
func (s Status) FromPending() bool { return s >= Pending }

// 用户自己是否可见
func (s Status) MeReadable() bool { return s >= MeReadableRange[0] && s <= MeReadableRange[1] }

// 用户是否可以修改、删除
func (s Status) Modifiable() bool { return s.In(Failed, Pending, Created, Passed) }

func (s Status) EasyPart() string {
	if s.IsOk() {
		return OkPart
	}
	return InvalidPart
}
func (s Status) Part() string {
	if s.IsPassed() {
		return PassedPart
	}
	return NotPassPart
}
func (s Status) PartA() string {
	if s < Pending {
		return InvalidPartA
	} else if s < 0 {
		return PendingPartA
	}
	return ValidPartA
}
func (s Status) PartB() string {
	if s < -99 {
		return DeletedPartB
	} else if s < Pending {
		return NonPublicPartB
	} else if s < 0 {
		return PendingPartB
	} else if s == 0 {
		return CreatedPartB
	}
	return PassedPartB
}
