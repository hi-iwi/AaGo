package aenum

type EducationDegree uint8

const (
	NoListedEducationDegree EducationDegree = 0
	BelowHighSchool         EducationDegree = 1
	HighSchoolDiploma       EducationDegree = 2 // 高中学历

	ColleageDegree         EducationDegree = 50 // 专科
	FullTimeColleageDegree EducationDegree = 51 // 全日制专科

	BachelorDegree         EducationDegree = 80 // 本科
	FullTimeBachelorDegree EducationDegree = 81 // 全日制本科
	EliteBachelorDegree    EducationDegree = 82 // 全日制211/985院校，精英院校

	PostgraduateDegree EducationDegree = 100 // 研究生
	DoctorateDegree    EducationDegree = 101 // 博士
	Postdoctoral       EducationDegree = 102 // 博士后
)
