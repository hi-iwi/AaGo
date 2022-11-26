package aenum

import "strconv"

type StatusR Status // resource status
const (
	SysDeletedR = StatusR(SysDeleted) // 系统删除
	SysRevokeR  = StatusR(SysRevoke)  // 审核通过后，紧急撤销，进入加强审核状态

	UserDeletedR = StatusR(UserDeleted) // 用户已删除

	_auditFailedLowerLimitR = StatusR(_auditFailedLowerLimit) // 审核中下限

	AuditFailedR = StatusR(AuditFailed) // 审核失败，让用户可以修改，仅用户可见

	_auditFailedUpperLimitR = StatusR(_auditFailedUpperLimit) // 审核中上限

	_pendingLowerLimitR = StatusR(_pendingLowerLimit) // 审核中下限

	PendingAuditR = StatusR(PendingAudit) // 审核中，不展示；通过则为 Approved；不通过则为：AuditFailed

	_pendingUpperLimitR = StatusR(_pendingUpperLimit) // 审核中上限

	// 临时基准线
	_approvedLowerLimitR = StatusR(_approvedLowerLimit) // 审核通过下限

	CreatedR  = StatusR(Created)  // 显示；未审核
	ApprovedR = StatusR(Approved) // 审核通过，所有人可见、可评论

	SysLockedR = StatusR(SysLocked) // 系统已锁定，用户不得再修改、评论
)

// 是否审核通过
func (s StatusR) IsApproved() bool {
	return s != 0 && s > _approvedLowerLimitR
}
func (s StatusR) IsPendding() bool {
	return s > _pendingLowerLimitR && s < _pendingUpperLimitR
}
func (s StatusR) IsFailed() bool {
	return s > _auditFailedLowerLimitR && s < _auditFailedUpperLimitR
}
func (s StatusR) MeVisible() bool {
	return s > _auditFailedLowerLimitR
}
func StmtPenddingR(field string) string {
	a := field + ">" + strconv.FormatInt(int64(_pendingLowerLimitR), 10)
	b := field + "<" + strconv.FormatInt(int64(_pendingUpperLimitR), 10)
	return a + " && " + b
}
func StmtFailedR(field string) string {
	a := field + ">=" + strconv.FormatInt(int64(_auditFailedLowerLimitR), 10)
	b := field + "<=" + strconv.FormatInt(int64(_auditFailedUpperLimitR), 10)
	return a + " && " + b
}

// 自己可见的声明语句
func StmtMeVisibleR(field string) string {
	return field + ">" + strconv.FormatInt(int64(_approvedLowerLimitR), 10)
}

// 自定义可见；如用户与作者互相关注，而且也被其分入特定用户组，那么可见语句为：
// s.VisibleStatusRStmt("status", OnlyVisibleToFansOk, OnlyVisibleToFolloweeOk, OnlyVisibleToSpecificUsersOk)
func StmtVisibleR(field string, sts ...StatusR) string {
	stmt := field + ">" + strconv.FormatInt(int64(_approvedLowerLimitR), 10)
	if len(sts) > 0 {
		for _, st := range sts {
			stmt += " OR " + field + "=" + strconv.FormatInt(int64(st), 10)
		}
	}
	return stmt
}
