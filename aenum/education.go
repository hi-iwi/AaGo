package aenum

type EduLevel uint8 // 学历
type EduConclusion uint8  // 毕结业结论
const (
	GraduationConclusion EduConclusion = 1 // 毕业

)
const (
	NoEduLevel        EduLevel = 0
	BelowHighSchool   EduLevel = 1
	HighSchoolDiploma EduLevel = 2 // 高中学历

	ColleageDegree         EduLevel = 50 // 专科
	FullTimeColleageDegree EduLevel = 51 // 全日制专科

	BachelorDegree         EduLevel = 80 // 本科
	FullTimeBachelorDegree EduLevel = 81 // 全日制本科
	EliteBachelorDegree    EduLevel = 82 // 全日制211/985院校，精英院校

	PostgraduateDegree EduLevel = 100 // 研究生
	DoctorateDegree    EduLevel = 101 // 博士
	Postdoctoral       EduLevel = 102 // 博士后
)
