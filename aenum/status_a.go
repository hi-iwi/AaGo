package aenum

//const (
//	// 最简易的 开关
//	InvalidPart = "invalid" // <0 关   off 是关键字。  aenum.StsInvalid("status")
//	OkPart      = "ok"      // < MAXVALUE 开           aenum.StsPublic("status")
//
//	//NegTernaryPart  = "neg" // -1, LESS 0
//	//ZeroTernaryPart = ""    // 0, LESS 1
//	//PosTernaryPart  = ""    // 1, LESS MAXVALUE
//
//	// 通过与否
//	NotPassPart = "notpass" // < 1      StsNotPassed()
//	PassedPart  = "passed"  // <MAXVALUE   StsPassed()
//
//	// 无用户操作的状态，仅后台管理
//	InvalidPartA = "invalid" // < Pending
//	PendingPartA = "pending" // < 0
//	ValidPartA   = "valid"   // < MAXVALUE
//
//	PublicPartAs  = ValidPartA              // aenum.StsPublic("status")
//	PendingPartAs = PendingPartA           // aenum.StsPending("status")
//
//	// 含用户操作的状态
//	DeletedPartB   = "deleted"   // < -99 后期需要清理的数据
//	NonPublicPartB = "nonpublic" // < Pending 仅用户自己可见
//	PendingPartB   = "pending"   // < 0
//	CreatedPartB   = "created"   // = 0 审核中，显示在公开列表（适用于白名单用户）
//	PassedPartB    = "passed"    // >=1 公开列表显示
//
//
//	PublicPartBs  = CreatedPartB + "," + PassedPartB // aenum.StsPublic("status")
//	PendingPartBs = PendingPartB + "," + CreatedPartB // aenum.StsPendingC("status")
//	VisPartBs     = PendingPartB + "," + PublicPartBs  //  aenum.StsVisible("status")
//	VisiblePartBs = NonPublicPartB + "," + VisPartBs // aenum.StsVis("status")
//
//)
//
//func (s Status) EasyPart() string {
//	if s.IsOk() {
//		return OkPart
//	}
//	return InvalidPart
//}
//func (s Status) Part() string {
//	if s.IsPassed() {
//		return PassedPart
//	}
//	return NotPassPart
//}
//func (s Status) PartA() string {
//	if s < Pending {
//		return InvalidPartA
//	} else if s < 0 {
//		return PendingPartA
//	}
//	return ValidPartA
//}
//func (s Status) PartB() string {
//	if s < -99 {
//		return DeletedPartB
//	} else if s < Pending {
//		return NonPublicPartB
//	} else if s < 0 {
//		return PendingPartB
//	} else if s == 0 {
//		return CreatedPartB
//	}
//	return PassedPartB
//}
