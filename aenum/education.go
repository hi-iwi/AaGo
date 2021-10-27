package aenum

type EduRating uint8     // 学院评级
type EduLevel uint8      // 学历
type EduConclusion uint8 // 毕结业结论

const (
	NormalUniversity EduRating = 0 // 一般院校
	GoodUniversity   EduRating = 1 // 好大学，双一流大学
	TalentUniversity EduRating = 2 // 人才院校，一般国内985
	EliteUniversity  EduRating = 3 // 精英院校，海外TOP50
)
const (
	GraduationConclusion EduConclusion = 1 // 毕业

)
const (
	NoEduLevel        EduLevel = 0
	BelowHighSchool   EduLevel = 1
	HighSchoolDiploma EduLevel = 2 // 高中学历

	ColleageStudent     EduLevel = 10 // 全日制专科 在读学生
	NonFullTimeColleage EduLevel = 11 // 非全日制专科
	FullTimeColleage    EduLevel = 12 // 全日制专科

	BachelorStudent     EduLevel = 20 // 在读全日制本科（非985院校）
	NonFullTimeBachelor EduLevel = 21 // 非全日制本科（获得学士学位了）  undergraduate degree 只有毕业证，没有学位证
	FullTimeBachelor    EduLevel = 22 // 全日制本科
	TalentBachelor      EduLevel = 23 // 高等人才学士，全日制985院校
	EliteBachelor       EduLevel = 24 // 精英阶层，海外TOP50院校

	NonFullTimePostgraduate EduLevel = 30 // 非全日制研究生
	FullTimePostgraduate    EduLevel = 31 // 研究生
	TalentPostgraduate      EduLevel = 32 // 人才硕士 本科是全日制985院校，研究生也是
	ElitePostgraduate       EduLevel = 33 // 精英硕士 本科是TOP50院校，硕士是海外TOP50院校

	NonFullTimeDoctorate EduLevel = 40 // 非全日制博士学位
	FullTimeDoctorate    EduLevel = 41 // 博士
	TalentDoctorate      EduLevel = 42 // 人才博士，本科是全日制985院校，研究生也是
	EliteDoctorate       EduLevel = 43 // 精英博士，本科是TOP50院校，硕士是海外TOP50院校
)
