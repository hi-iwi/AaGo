package aenum

import "strconv"

type StatusR int8 // resource status
const (
	SysDeletedR Status = SysDeleted // 系统删除
	SysRevokeR  Status = SysRevoke  // 审核通过后，紧急撤销，进入加强审核状态

	UserDeletedR Status = UserDeleted // 用户已删除

	AuditFailedR Status = AuditFailed // 审核失败，让用户可以修改，仅用户可见

	PendingAuditR Status = PendingAudit // 审核中，不展示；通过则为 Approved；不通过则为：AuditFailed

	CreatedR  Status = Created  // 显示；未审核
	ApprovedR Status = Approved // 审核通过，所有人可见、可评论

	SysLockedR Status = SysLocked // 系统已锁定，用户不得再修改、评论
)

// 自定义可见；如用户与作者互相关注，而且也被其分入特定用户组，那么可见语句为：
// s.VisibleStatusRStmt("status", OnlyVisibleToFansOk, OnlyVisibleToFolloweeOk, OnlyVisibleToSpecificUsersOk)
func VisibleStatusRStmt(field string, sts ...StatusR) string {
	stmt := field + ">-1" // 0 是可见的
	if len(sts) > 0 {
		for _, st := range sts {
			stmt += " OR " + field + "=" + strconv.FormatInt(int64(st), 10)
		}
	}
	return stmt
}
