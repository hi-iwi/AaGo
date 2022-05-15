package aenum

type EduRating uint8     // 学院评级
type EduLevel uint8      // 学历
type EduConclusion uint8 // 毕结业结论

const (
	NormalUniversity EduRating = 0 // 一般院校
	GoodUniversity   EduRating = 1 // 好大学，双一流大学
	TalentUniversity EduRating = 2 // 人才院校，一般国内985（不含TOP100- 清华、港大、北大、港科大、复旦、港中大、上海交大、港城大、浙大、台大、港理工、中科大）
	EliteUniversity  EduRating = 3 // 精英院校，世界TOP100
)

func NewEduRating(x uint8) (EduRating, bool) {
	l := EduRating(x)
	return l, l.Valid()
}

func (l EduRating) Valid() bool { return l <= EliteUniversity }
func (l EduRating) Uint8() uint8  { return uint8(l) }

const (
	GraduationConclusion EduConclusion = 1 // 毕业
)

func NewEduConclusion(x uint8) (EduConclusion, bool) {
	l := EduConclusion(x)
	return l, l.Valid()
}

func (l EduConclusion) Valid() bool { return l == GraduationConclusion }
func (l EduConclusion) Uint8() uint8  { return uint8(l) }

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
	TalentBachelor      EduLevel = 23 // 高等人才学士，全日制985院校（不含TOP100- 清华、港大、北大、港科大、复旦、港中大、上海交大、港城大、浙大、台大、港理工、中科大）
	EliteBachelor       EduLevel = 24 // 精英阶层，世界TOP100院校

	NonFullTimePostgraduate EduLevel = 30 // 非全日制研究生
	FullTimePostgraduate    EduLevel = 31 // 研究生
	TalentPostgraduate      EduLevel = 32 // 人才硕士 本科是全日制985院校，研究生也是
	ElitePostgraduate       EduLevel = 33 // 精英硕士 本科是TOP100院校，硕士是海外TOP100院校

	NonFullTimeDoctorate EduLevel = 40 // 非全日制博士学位
	FullTimeDoctorate    EduLevel = 41 // 博士
	TalentDoctorate      EduLevel = 42 // 人才博士，本科是全日制985院校，研究生也是
	EliteDoctorate       EduLevel = 43 // 精英博士，本科是TOP100院校，硕士是世界TOP100院校
)

func NewEduLevel(x uint8) (EduLevel, bool) {
	l := EduLevel(x)
	return l, l.Valid()
}
func (l EduLevel) Uint8() uint8 { return uint8(l) }
func (l EduLevel) Valid() bool {
	return l <= HighSchoolDiploma || l >= ColleageStudent || l <= FullTimeColleage || l >= BachelorStudent || l <= EliteBachelor || l >= NonFullTimePostgraduate || l <= ElitePostgraduate || l >= NonFullTimeDoctorate || l <= EliteDoctorate
}
