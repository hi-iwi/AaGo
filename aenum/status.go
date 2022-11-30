package aenum

import (
	"strconv"
)

type Status int8

const (
	SysRevoked Status = -128 // 已注销，系统删除（可能审核失败）
	SysExpired Status = -121 // 已失效/已过期
	Deleted    Status = -100 // 用户已删除，谁都不可见
	Failed     Status = -20  // 审核失败，让用户可以修改，仅用户可见  --> 无论之前设置什么，这里统一一个失败状态，修改的时候再设置
	Pending1   Status = -10  // 审核1通过 --> 进入阶段性审核流程，就会阻却用户修改权
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
	PendingRange    = [2]Status{Pending1, Created} // 等待审核的区间
	MeReadableRange = [2]Status{Failed, SysLocked} // 会显示在我列表的区间
)
var StatusNames = map[Language]map[Status]string{
	ZhCN: {
		SysRevoked: "已注销",
		SysExpired: "已失效",
		Deleted:    "已删除",
		Failed:     "未通过审核",
		Pending1:   "一审",
		Pending2:   "二审",
		Pending3:   "三审",
		Pending4:   "四审",
		Pending5:   "五审",
		Pending6:   "六审",
		Pending7:   "七审",
		Pending8:   "八审",
		Pending9:   "九审",
		Pending:    "审核中",
		Created:    "新创建",
		Passed:     "审核通过",
		SysLocked:  "已锁定",
	},
	EnUS: {
		SysRevoked: "revoked",
		SysExpired: "expired",
		Deleted:    "deleted",
		Failed:     "failed",
		Pending1:   "passing 1",
		Pending2:   "passing 2",
		Pending3:   "passing 3",
		Pending4:   "passing 4",
		Pending5:   "passing 5",
		Pending6:   "passing 6",
		Pending7:   "passing 7",
		Pending8:   "passing 8",
		Pending9:   "passing 9",
		Pending:    "pending",
		Created:    "created",
		Passed:     "passed",
		SysLocked:  "locked",
	},
}

func NewStatus(sts int8) (Status, bool) {
	s := Status(sts)
	ok := s == SysRevoked || s == SysExpired || s == Deleted || s == Failed
	ok = ok || (s >= Pending1 && s <= Passed) || s == SysLocked
	return s, ok
}
func (s Status) Name(lang Language) string {
	if names, ok := StatusNames[lang]; ok {
		if name, ok := names[s]; ok {
			return name
		}
	}
	return strconv.FormatInt(int64(s), 10)
}
func (s Status) NameZh() string {
	return s.Name(ZhCN)
}
func (s Status) NameEn() string {
	return s.Name(EnUS)
}

func (s Status) IsPending() bool {
	return s >= PendingRange[0] && s <= PendingRange[1]
}

// 是否审核通过
func (s Status) IsPassed() bool {
	return s >= Passed
}

// 用户自己是否可见
func (s Status) MeReadable() bool {
	return s >= MeReadableRange[0] && s <= MeReadableRange[1]
}

// 用户是否可以修改、删除
func (s Status) Modifiable() bool {
	return s == Failed || s == Pending || s == Created || s == Passed
}
