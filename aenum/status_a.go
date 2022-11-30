package aenum

//type StatusA int8
//
//const (
//	OnlyVisibleToMe            Status = -6 // 用户设置：仅自己可见，审核中
//	OnlyVisibleToSpecificUsers Status = -5 // 用户设置：只对特定用户开放，审核中
//	NotVisibleToSpecificUsers  Status = -6 // 用户设置：选定的特定用户不展示
//	OnlyVisibleToFollowee      Status = -3 // 用户设置：仅自己关注者可见，审核中
//	OnlyVisibleToFans          Status = -2 // 用户设置：仅粉丝可见，审核中
//
//	OnlyMeCanComment            Status = -5 // 用户设置：仅自己可评论（任何人都可以看），审核中
//	OnlySpecificUsersCanComment Status = -4 // 用户设置：仅特定用户可评论（任何人都可以看），审核中
//	OnlyFolloweeCanComment      Status = -3 // 用户设置：仅自己关注者可评论（任何人都可以看），审核中
//	OnlyFansCanComment          Status = -2 // 用户设置：仅粉丝可评论（任何人都可以看），审核中
//
//	OnlyFansCanCommentOk          Status = 2 // 审核通过后，用户设置：仅粉丝可评论（任何人都可以看）
//	OnlyFolloweeCanCommentOk      Status = 3 // 审核通过后，用户设置：仅自己关注者可评论（任何人都可以看）
//	OnlySpecificUsersCanCommentOk Status = 4 // 审核通过后，用户设置：仅特定用户可评论（任何人都可以看）
//	OnlyMeCanCommentOk            Status = 5 // 审核通过后，用户设置：仅自己可评论（任何人都可以看） --> 仅自己可以评论
//
//	OnlyVisibleToMeOk            Status = 21 // 审核通过, 仅自己可见
//	OnlyVisibleToSpecificUsersOk Status = 22 // 审核通过, 只对特定用户开放
//	NotVisibleToSpecificUsersOk  Status = 23 // 审核通过, 选定的特定用户不展示
//	OnlyVisibleToFolloweeOk      Status = 24 // 审核通过, 仅自己关注者可见
//	OnlyVisibleToFansOk          Status = 25 // 审核通过, 仅粉丝可见
//
//	Marked                            Status = 60 // 审核通过后，用户进行标记，如置顶、加精 --> 提示，修改为置顶，将对所有人开放
//	MarkedOnlyMeCanComment            Status = 61 // 审核通过后，用户进行标记，如置顶、加精 --> 提示，修改为置顶，将对所有人开放
//	MarkedOnlyFansCanComment          Status = 62 // 审核通过后，用户进行标记，如置顶、加精 --> 提示，修改为置顶，将对所有人开放
//	MarkedOnlyFolloweeCanComment      Status = 63 // 审核通过后，用户进行标记，如置顶、加精 --> 提示，修改为置顶，将对所有人开放
//	MarkedOnlySpecificUsersCanComment Status = 64 // 审核通过后，用户进行标记，如置顶、加精 --> 提示，修改为置顶，将对所有人开放
//
//	SysLockedComment Status = 120 // 系统已锁定，所有人禁止评论（仅作者可以评论、删除评论）
//	SysLockedModify  Status = 126 // 系统锁定，禁止修改（允许评论）
//)
//