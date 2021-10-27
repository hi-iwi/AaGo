package aenum

type EduLevel uint8      // 学历
type EduConclusion uint8 // 毕结业结论
const (
	GraduationConclusion EduConclusion = 1 // 毕业

)
const (
	NoEduLevel        EduLevel = 0
	BelowHighSchool   EduLevel = 1
	HighSchoolDiploma EduLevel = 2 // 高中学历

	NonFullTimeColleage EduLevel = 10 // 非全日制专科
	FullTimeColleage    EduLevel = 11 // 全日制专科

	NonFullTimeBachelor EduLevel = 20 // 非全日制本科（获得学士学位了）  undergraduate degree 只有毕业证，没有学位证
	FullTimeBachelor    EduLevel = 21 // 全日制本科
	TalentBachelor      EduLevel = 22 // 高等人才学士，全日制211/985院校
	EliteBachelor       EduLevel = 23 // 精英阶层，海外TOP100院校

	NonFullTimePostgraduate EduLevel = 30 // 非全日制研究生
	FullTimePostgraduate    EduLevel = 31 // 研究生
	TalentPostgraduate      EduLevel = 32 // 人才硕士 本科是全日制985/211院校，研究生也是
	ElitePostgraduate       EduLevel = 33 // 精英硕士 本科是TOP100院校，硕士是海外TOP100院校

	NonFullTimeDoctorate EduLevel = 40 // 非全日制博士学位
	FullTimeDoctorate    EduLevel = 41 // 博士
	TalentDoctorate      EduLevel = 42 // 人才博士，本科是全日制985/211院校，研究生也是
	EliteDoctorate       EduLevel = 43 // 精英博士，本科是TOP100院校，硕士是海外TOP100院校
)
