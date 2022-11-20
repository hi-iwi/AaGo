package aenum

import "strconv"

type Status int8 // content status

const (
	SysDeleted Status = -128 // 系统删除
	SysRevoke  Status = -120 // 审核通过后，紧急撤销，进入加强审核状态

	UserDeleted Status = -90 // 用户已删除

	TmpAuditFailedLowerLimit Status = UserDeleted // 审核中下限

	NotVisibleToSpecificUsersFailed  Status = -85 // 用户设置：选定的特定用户不展示
	OnlyVisibleToSpecificUsersFailed Status = -84 // 用户设置：只对特定用户开放后，审核失败
	OnlyVisibleToFolloweeFailed      Status = -83 // 用户设置：仅自己关注者可见后，审核失败
	OnlyVisibleToFansFailed          Status = -82 // 用户设置：仅粉丝可见后，审核失败
	OnlyVisibleToMeFailed            Status = -81 // 用户设置：仅自己可见后，审核失败
	AuditFailed                      Status = -80 // 审核失败，让用户可以修改，仅用户可见

	TmpAuditFailedUpperLimit Status = AuditFailed - 1 // 审核中上限

	TmpPendingLowerLimit Status = TmpAuditFailedUpperLimit // 审核中下限

	NotVisibleToSpecificUsers  Status = -49 // 用户设置：选定的特定用户不展示
	OnlyVisibleToSpecificUsers Status = -44 // 用户设置：只对特定用户开放，审核中
	OnlyVisibleToFollowee      Status = -43 // 用户设置：仅自己关注者可见，审核中
	OnlyVisibleToFans          Status = -42 // 用户设置：仅粉丝可见，审核中
	OnlyVisibleToMe            Status = -41 // 用户设置：仅自己可见，审核中
	PendingAudit               Status = -40 // 审核中，不展示；通过则为 Approved；不通过则为：AuditFailed

	TmpPendingUpperLimit Status = PendingAudit - 1 // 审核中上限

	// 临时基准线
	TmpApprovedLowerLimit Status = TmpPendingUpperLimit // 审核通过下限

	NotVisibleToSpecificUsersOk  Status = -9 // 用户设置：选定的特定用户不展示，审核通过
	OnlyVisibleToSpecificUsersOk Status = -4 // 用户设置：只对特定用户开放后，审核通过
	OnlyVisibleToFolloweeOk      Status = -3 // 用户设置：仅自己关注者可见后，审核通过
	OnlyVisibleToFansOk          Status = -2 // 用户设置：仅粉丝可见后，审核通过
	OnlyVisibleToMeOk            Status = -1 // 用户设置：仅自己可见后，审核通过

	Created  Status = 0 // 显示；未审核
	Approved Status = 1 // 审核通过，所有人可见、可评论

	OnlyMeCanComment            Status = 11 // 审核通过后，用户设置：仅自己可评论（任何人都可以看）
	OnlyFansCanComment          Status = 12 // 审核通过后，用户设置：仅粉丝可评论（任何人都可以看）
	OnlyFolloweeCanComment      Status = 13 // 审核通过后，用户设置：仅自己关注者可评论（任何人都可以看）
	OnlySpecificUsersCanComment Status = 14 // 审核通过后，用户设置：仅特定用户可评论（任何人都可以看）

	Marked                            Status = 20 // 审核通过后，用户进行标记，如置顶、加精 --> 提示，修改为置顶，将对所有人开放
	MarkedOnlyMeCanComment            Status = 21 // 审核通过后，用户进行标记，如置顶、加精 --> 提示，修改为置顶，将对所有人开放
	MarkedOnlyFansCanComment          Status = 22 // 审核通过后，用户进行标记，如置顶、加精 --> 提示，修改为置顶，将对所有人开放
	MarkedOnlyFolloweeCanComment      Status = 23 // 审核通过后，用户进行标记，如置顶、加精 --> 提示，修改为置顶，将对所有人开放
	MarkedOnlySpecificUsersCanComment Status = 24 // 审核通过后，用户进行标记，如置顶、加精 --> 提示，修改为置顶，将对所有人开放

	SysLockedComment Status = 120 // 系统已锁定，所有人禁止评论（仅作者可以评论、删除评论）
	SysLocked        Status = 127 // 系统已锁定，用户不得再修改、评论
)

// 是否审核通过
func (s Status) IsApprovedStatus() bool {
	return s != 0 && s > TmpApprovedLowerLimit
}
func (s Status) IsPenddingStatus() bool {
	return s > TmpPendingLowerLimit && s < TmpPendingUpperLimit
}
func (s Status) IsFailedStatus() bool {
	return s > TmpAuditFailedLowerLimit && s < TmpAuditFailedUpperLimit
}
func PenddingStatusStmt(field string) string {
	a := field + ">" + strconv.FormatInt(int64(TmpPendingLowerLimit), 10)
	b := field + "<" + strconv.FormatInt(int64(TmpPendingUpperLimit), 10)
	return a + " && " + b
}
func FailedStatusStmt(field string) string {
	a := field + ">=" + strconv.FormatInt(int64(TmpAuditFailedLowerLimit), 10)
	b := field + "<=" + strconv.FormatInt(int64(TmpAuditFailedUpperLimit), 10)
	return a + " && " + b
}

// 自己可见的声明语句
func MeVisibleStatusStmt(field string) string {
	return field + ">" + strconv.FormatInt(int64(TmpApprovedLowerLimit), 10)
}

// 自定义可见；如用户与作者互相关注，而且也被其分入特定用户组，那么可见语句为：
// s.VisibleStatusStmt("status", OnlyVisibleToFansOk, OnlyVisibleToFolloweeOk, OnlyVisibleToSpecificUsersOk)
func VisibleStatusStmt(field string, sts ...Status) string {
	stmt := field + ">-1" // 0 是可见的
	if len(sts) > 0 {
		for _, st := range sts {
			stmt += " OR " + field + "=" + strconv.FormatInt(int64(st), 10)
		}
	}
	return stmt
}
