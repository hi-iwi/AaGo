package aenum

import "strconv"

type Status int8 // content status

type DataParsing int8 // 解析远程数据，储存远程数据记录时用到

const (
	Deleted     Status = -128 //已删除
	AuditFailed Status = -127 // 审核失败，让用户可以修改，仅用户可见

	OnlyVisibleToFansFailed          Status = -85 // 审核失败，用户设置：仅粉丝可见
	OnlyVisibleToFolloweeFailed      Status = -84 // 审核失败，用户设置：仅自己关注者可见
	OnlyVisibleToSpecificUsersFailed Status = -83 // 审核失败，用户设置：只对特定用户开放
	OnlyVisibleToMeFailed            Status = -82 // 审核失败+，用户设置：仅自己可见

	OnlyVisibleToFans          Status = -44 // 未审核，用户设置：仅粉丝可见
	OnlyVisibleToFollowee      Status = -43 // 未审核，用户设置：仅自己关注者可见
	OnlyVisibleToSpecificUsers Status = -42 // 未审核，用户设置：只对特定用户开放
	OnlyVisibleToMe            Status = -41 // 未审核，用户设置：仅自己可见
	PendingAudit               Status = -40 // 未审核，仅自己可见

	OnlyVisibleToFansOk          Status = -4 // 审核通过，用户设置：仅粉丝可见
	OnlyVisibleToFolloweeOk      Status = -3 // 审核通过，用户设置：仅自己关注者可见
	OnlyVisibleToSpecificUsersOk Status = -2 // 审核通过，用户设置：只对特定用户开放
	OnlyVisibleToMeOk            Status = -1 // 审核通过，用户设置：仅自己可见

	Created  Status = 0   // 显示
	Approved Status = 1   // 审核通过，展示
	Topping  Status = 127 // 审核通过，用户设置：置顶展示
)

// 是否审核通过
func (s Status) IsApprovedStatus() bool {
	return s != 0 && s >= OnlyVisibleToFansOk
}
func (s Status) IsPenddingStatus() bool {
	return s >= OnlyVisibleToFans && s <= PendingAudit
}
func (s Status) IsFailedStatus() bool {
	return s > Deleted && s <= OnlyVisibleToMeFailed
}
func (s Status) PenddingStatement(field string) string {
	a := field + ">=" + strconv.FormatInt(int64(OnlyVisibleToFans), 10)
	b := field + "<=" + strconv.FormatInt(int64(PendingAudit), 10)
	return a + " && " + b
}
func (s Status) FailedStatement(field string) string {
	a := field + ">=" + strconv.FormatInt(int64(Deleted), 10)
	b := field + "<=" + strconv.FormatInt(int64(OnlyVisibleToMeFailed), 10)
	return a + " && " + b
}

// 自己可见的声明语句
func (s Status) MeVisibleStatement(field string) string {
	return field + ">" + strconv.FormatInt(int64(Deleted), 10)
}

// 自定义可见；如用户与作者互相关注，而且也被其分入特定用户组，那么可见语句为：
// s.VisibleStatement("status", OnlyVisibleToFansOk, OnlyVisibleToFolloweeOk, OnlyVisibleToSpecificUsersOk)
func (s Status) VisibleStatement(field string, sts ...Status) string {
	stmt := field + ">0"
	if len(sts) > 0 {
		for _, st := range sts {
			stmt += " OR " + field + "=" + strconv.FormatInt(int64(st), 10)
		}
	}
	return stmt
}

const (
	DataParsingCheckFailed DataParsing = -2 // 数据签名核对错误、字段核对错误
	DataParsingFailed      DataParsing = -1 // 数据解析失败
	DataParsingBizFailed   DataParsing = 0  // 数据解析成功了，但是业务结果返回失败
	DataParsingBizOK       DataParsing = 1  // 数据解析成功了，并且业务结果返回成功
)

const (
	HttpStatusAccountConflict = 490 // 授权登录成功，但是该授权账户绑定过的UID，和当前登录的UID不一致。
	HttpStatusAccountUnlinked = 491 // 已经授权登陆，但是需要绑定手机号/账号
)
