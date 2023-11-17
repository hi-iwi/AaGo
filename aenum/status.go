package aenum

import (
	"strconv"
)

type Status int8 // status 字段禁止默认为0，防止意外情况

const (
	SysRevoked Status = -128 // 已注销，系统删除（可能审核失败）
	Expired    Status = -121 // 已失效/已过期，后面系统会自动删除
	NotExists  Status = -114 // 数据不存在 --> 一般用于初始化一个字段，不常用
	Deleted    Status = -100 // 用户已删除，谁都不可见
	/********************* <=-100  属于需要删除的内容，方便分区删除整个分区 ************************/
	Supervised  Status = -90 // 系统检测到异常，自动锁定帐户。等待异常解除。
	Deleting    Status = -89 // 删除中；冷静期后，才进入 Deleted
	Suspend     Status = -80 // 暂停/暂时隐藏，由用户设置
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
	deletedStr      = strconv.Itoa(int(Deleted))
	pendingLimit    = strconv.Itoa(int(Pending - 1))
	MeReadableRange = [2]Status{Failed, SysTakeover} // 会显示在我列表的区间
)

func NewStatus(sts int8) (Status, bool) {
	s := Status(sts)
	ok := s == SysRevoked || s == Expired || s == NotExists || s == Deleted || s == Failed
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
func (s Status) IsOk() bool  { return s >= Created }
func (s Status) NotOk() bool { return s < Created }

// 是否审核通过
func (s Status) IsPassed() bool  { return s >= Passed }
func (s Status) NotPassed() bool { return s < Passed }

// Created 也是待审核状态
func (s Status) IsPending() bool { return s >= Pending && s <= Created }
func (s Status) IsFailed() bool  { return s <= Failed }
func (s Status) IsDiscard() bool { return s <= Deleting }
func (s Status) NotExits() bool  { return s == NotExists }

// After/Before 不包括； From/Within 包括
// FromPending 一般用户检测某个状态是否在审核中，或者已审核通过。
func (s Status) FromPending() bool { return s >= Pending }

// 用户自己是否可见
func (s Status) MeReadable() bool { return s >= MeReadableRange[0] && s <= MeReadableRange[1] }

// 用户是否可以修改、删除
func (s Status) Modifiable() bool { return s.In(Failed, Pending, Created, Passed) }

func StsInvalid(k string) string    { return k + "<0" }
func StsPublic(k string) string     { return k + ">-1" }
func StsPassed(k string) string     { return k + ">0" }
func StsNotPassed(k string) string  { return k + "<1" }
func StsNotDiscard(k string) string { return k + ">-90" }

// 公开可见的状态，即审核中、创建、审核通过
func StsVisible(k string) string { return k + ">" + pendingLimit }

// 纯审核中
func StsPending(k string) string {
	return k + ">" + pendingLimit + " AND " + k + "<0"
}

// 审核中、创建的状态
func StsPendingC(k string) string {
	return k + ">" + pendingLimit + " AND " + k + "<1"
}

// 自己可见
func StsVis(k string) string { return k + ">" + deletedStr }
