package aenum

import "strconv"

type Ethn uint8 // 全世界大概2000个民族；1-99 为中国的民族
// type Ethnicity uint16 // 全世界民族

const (
	Han       Ethn = 1
	Achang    Ethn = 2
	Bai       Ethn = 3
	Baoan     Ethn = 4
	Bulang    Ethn = 5
	Buyi      Ethn = 6
	Chaoxian  Ethn = 7
	Dawoer    Ethn = 8
	Dai       Ethn = 9
	Deang     Ethn = 10
	Dong      Ethn = 11
	Dongxiang Ethn = 12
	Dulong    Ethn = 13
	Elunchun  Ethn = 14
	Eluosi    Ethn = 15
	Ewenke    Ethn = 16
	Gaoshan   Ethn = 17
	Gelao     Ethn = 18
	Hani      Ethn = 19
	Hasake    Ethn = 20
	Hezhe     Ethn = 21
	Hui       Ethn = 22
	Jinuo     Ethn = 23
	Jing      Ethn = 24
	Jingpo    Ethn = 25
	Keerkezi  Ethn = 26
	Lahu      Ethn = 27
	Lizu      Ethn = 28
	Lisu      Ethn = 29
	Luoba     Ethn = 30
	Manzu     Ethn = 31
	Maonan    Ethn = 32
	Menba     Ethn = 33
	Menggu    Ethn = 34
	Miao      Ethn = 35
	Mulao     Ethn = 36
	Naxi      Ethn = 37
	Nuzu      Ethn = 38
	Pumi      Ethn = 39
	Qiang     Ethn = 40
	Sala      Ethn = 41
	Shezu     Ethn = 42
	Shuizu    Ethn = 43
	Tajike    Ethn = 44
	Tataer    Ethn = 45
	Tuzu      Ethn = 46
	Tujia     Ethn = 47
	Wazu      Ethn = 48
	Weiwuer   Ethn = 49
	Wuzibieke Ethn = 50
	Xibo      Ethn = 51
	Yaozu     Ethn = 52
	Yizu      Ethn = 53
	Yugu      Ethn = 54
	Zangzu    Ethn = 55
	Zhuangzu  Ethn = 56
	OtherEthn Ethn = 99 // 其他中国民族
)

var EthnNames = map[Ethn]string{
	Han:       "汉族",
	Achang:    "阿昌族",
	Bai:       "白族",
	Baoan:     "保安族",
	Bulang:    "布朗族",
	Buyi:      "布依族",
	Chaoxian:  "朝鲜族",
	Dawoer:    "达斡尔族",
	Dai:       "傣族",
	Deang:     "德昂族",
	Dong:      "侗族",
	Dongxiang: "东乡族",
	Dulong:    "独龙族",
	Elunchun:  "鄂伦春族",
	Eluosi:    "俄罗斯族",
	Ewenke:    "鄂温克族",
	Gaoshan:   "高山族",
	Gelao:     "仡佬族",
	Hani:      "哈尼族",
	Hasake:    "哈萨克族",
	Hezhe:     "赫哲族",
	Hui:       "回族",
	Jinuo:     "基诺族",
	Jing:      "京族",
	Jingpo:    "景颇族",
	Keerkezi:  "柯尔克孜族",
	Lahu:      "拉祜族",
	Lizu:      "黎族",
	Lisu:      "傈僳族",
	Luoba:     "珞巴族",
	Manzu:     "满族",
	Maonan:    "毛南族",
	Menba:     "门巴族",
	Menggu:    "蒙古族",
	Miao:      "苗族",
	Mulao:     "仫佬族",
	Naxi:      "纳西族",
	Nuzu:      "怒族",
	Pumi:      "普米族",
	Qiang:     "羌族",
	Sala:      "撒拉族",
	Shezu:     "畲族",
	Shuizu:    "水族",
	Tajike:    "塔吉克族",
	Tataer:    "塔塔尔族",
	Tuzu:      "土族",
	Tujia:     "土家族",
	Wazu:      "佤族",
	Weiwuer:   "维吾尔族",
	Wuzibieke: "乌兹别克族",
	Xibo:      "锡伯族",
	Yaozu:     "瑶族",
	Yizu:      "彝族",
	Yugu:      "裕固族",
	Zangzu:    "藏族",
	Zhuangzu:  "壮族",
	OtherEthn: "其他民族",
}

func NewEthn(ethn uint8) (Ethn, bool) {
	et := Ethn(ethn)
	ok := (et > 0 && et < Zhuangzu) || et == OtherEthn
	return et, ok
}
func ToEthn(s string) (Ethn, bool) {
	for ethn, name := range EthnNames {
		if s == name || s+"族" == name {
			return ethn, true
		}
	}
	return 0, false
}

func (s Ethn) Uint8() uint8    { return uint8(s) }
func (s Ethn) String() string  { return strconv.Itoa(int(s)) }
func (s Ethn) Is(x uint8) bool { return s.Uint8() == x }
func (s Ethn) In(a ...Ethn) bool {
	for _, sts := range a {
		if sts == s {
			return true
		}
	}
	return false
}
func (s Ethn) Name() string {
	name, _ := EthnNames[s]
	return name
}
