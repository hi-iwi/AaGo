package aenum

import (
	"strconv"
	"strings"
)

type CountryCode uint16

const (
	Canada                       CountryCode = 50001
	America                      CountryCode = 1
	Kazakhstan                   CountryCode = 50007
	Russia                       CountryCode = 7
	Egypt                        CountryCode = 20
	SouthAfrica                  CountryCode = 27
	Greece                       CountryCode = 30
	Netherlands                  CountryCode = 31
	Belgium                      CountryCode = 32
	France                       CountryCode = 33
	Spain                        CountryCode = 34
	Hungary                      CountryCode = 36
	Italy                        CountryCode = 39
	Romania                      CountryCode = 40
	Switzerland                  CountryCode = 41
	Austria                      CountryCode = 43
	UnitedKingdom                CountryCode = 44
	Denmark                      CountryCode = 45
	Sweden                       CountryCode = 46
	Norway                       CountryCode = 47
	SvalbardAndJanMayen          CountryCode = 50047
	Poland                       CountryCode = 48
	Germany                      CountryCode = 49
	Peru                         CountryCode = 51
	Mexico                       CountryCode = 52
	Cuba                         CountryCode = 53
	Argentina                    CountryCode = 54
	Brazil                       CountryCode = 55
	Chile                        CountryCode = 56
	Colombia                     CountryCode = 57
	Venezuela                    CountryCode = 58
	Malaysia                     CountryCode = 60
	Australia                    CountryCode = 61
	ChristmasIsland              CountryCode = 50061
	CocosIslands                 CountryCode = 61
	Indonesia                    CountryCode = 62
	Philippines                  CountryCode = 63
	NewZealand                   CountryCode = 64
	Pitcairn                     CountryCode = 50064
	Singapore                    CountryCode = 65
	Thailand                     CountryCode = 66
	Japan                        CountryCode = 81
	SouthKorea                   CountryCode = 82
	Vietnam                      CountryCode = 84
	China                        CountryCode = 86
	Turkey                       CountryCode = 90
	India                        CountryCode = 91
	Pakistan                     CountryCode = 92
	Afghanistan                  CountryCode = 93
	SriLanka                     CountryCode = 94
	Myanmar                      CountryCode = 95
	Iran                         CountryCode = 98
	SouthSudan                   CountryCode = 211
	Morocco                      CountryCode = 212
	WesternSahara                CountryCode = 50212
	Algeria                      CountryCode = 213
	Tunisia                      CountryCode = 216
	Libya                        CountryCode = 218
	Gambia                       CountryCode = 220
	Senegal                      CountryCode = 221
	Mauritania                   CountryCode = 222
	Mali                         CountryCode = 223
	Guinea                       CountryCode = 224
	IvoryCoast                   CountryCode = 225
	BurkinaFaso                  CountryCode = 226
	Niger                        CountryCode = 227
	Togo                         CountryCode = 228
	Benin                        CountryCode = 229
	Mauritius                    CountryCode = 230
	Liberia                      CountryCode = 231
	SierraLeone                  CountryCode = 232
	Ghana                        CountryCode = 233
	Nigeria                      CountryCode = 234
	Chad                         CountryCode = 235
	CentralAfricanRepublic       CountryCode = 236
	Cameroon                     CountryCode = 237
	CapeVerde                    CountryCode = 238
	SaoTomeAndPrincipe           CountryCode = 239
	EquatorialGuinea             CountryCode = 240
	Gabon                        CountryCode = 241
	RepublicOfTheCongo           CountryCode = 242
	DemocraticRepublicOfTheCongo CountryCode = 243
	Angola                       CountryCode = 244
	GuineaBissau                 CountryCode = 245
	BritishIndianOceanTerritory  CountryCode = 246
	Seychelles                   CountryCode = 248
	Sudan                        CountryCode = 249
	Rwanda                       CountryCode = 250
	Ethiopia                     CountryCode = 251
	Somalia                      CountryCode = 252
	Djibouti                     CountryCode = 253
	Kenya                        CountryCode = 254
	Tanzania                     CountryCode = 255
	Uganda                       CountryCode = 256
	Burundi                      CountryCode = 257
	Mozambique                   CountryCode = 258
	Zambia                       CountryCode = 260
	Madagascar                   CountryCode = 261
	Mayotte                      CountryCode = 262
	Reunion                      CountryCode = 262
	Zimbabwe                     CountryCode = 263
	Namibia                      CountryCode = 264
	Malawi                       CountryCode = 265
	Lesotho                      CountryCode = 266
	Botswana                     CountryCode = 267
	Swaziland                    CountryCode = 268
	Comoros                      CountryCode = 269
	SaintHelena                  CountryCode = 290
	Eritrea                      CountryCode = 291
	Aruba                        CountryCode = 297
	FaroeIslands                 CountryCode = 298
	Greenland                    CountryCode = 299
	Gibraltar                    CountryCode = 350
	Portugal                     CountryCode = 351
	Luxembourg                   CountryCode = 352
	Ireland                      CountryCode = 353
	Iceland                      CountryCode = 354
	Albania                      CountryCode = 355
	Malta                        CountryCode = 356
	Cyprus                       CountryCode = 357
	Finland                      CountryCode = 358
	Bulgaria                     CountryCode = 359
	Lithuania                    CountryCode = 370
	Latvia                       CountryCode = 371
	Estonia                      CountryCode = 372
	Moldova                      CountryCode = 373
	Armenia                      CountryCode = 374
	Belarus                      CountryCode = 375
	Andorra                      CountryCode = 376
	Monaco                       CountryCode = 377
	SanMarino                    CountryCode = 378
	Vatican                      CountryCode = 379
	Ukraine                      CountryCode = 380
	Serbia                       CountryCode = 381
	Montenegro                   CountryCode = 382
	Kosovo                       CountryCode = 383
	Croatia                      CountryCode = 385
	Slovenia                     CountryCode = 386
	BosniaAndHerzegovina         CountryCode = 387
	Macedonia                    CountryCode = 389
	CzechRepublic                CountryCode = 420
	Slovakia                     CountryCode = 421
	Liechtenstein                CountryCode = 423
	FalklandIslands              CountryCode = 500
	Belize                       CountryCode = 501
	Guatemala                    CountryCode = 502
	ElSalvador                   CountryCode = 503
	Honduras                     CountryCode = 504
	Nicaragua                    CountryCode = 505
	CostaRica                    CountryCode = 506
	Panama                       CountryCode = 507
	SaintPierreAndMiquelon       CountryCode = 508
	Haiti                        CountryCode = 509
	SaintBarthelemy              CountryCode = 590
	SaintMartin                  CountryCode = 590
	Bolivia                      CountryCode = 591
	Guyana                       CountryCode = 592
	Ecuador                      CountryCode = 593
	Paraguay                     CountryCode = 595
	Suriname                     CountryCode = 597
	Uruguay                      CountryCode = 598
	Curacao                      CountryCode = 599
	NetherlandsAntilles          CountryCode = 50599
	EastTimor                    CountryCode = 670
	Antarctica                   CountryCode = 672
	Brunei                       CountryCode = 673
	Nauru                        CountryCode = 674
	PapuaNewGuinea               CountryCode = 675
	Tonga                        CountryCode = 676
	SolomonIslands               CountryCode = 677
	Vanuatu                      CountryCode = 678
	Fiji                         CountryCode = 679
	Palau                        CountryCode = 680
	WallisAndFutuna              CountryCode = 681
	CookIslands                  CountryCode = 682
	Niue                         CountryCode = 683
	Samoa                        CountryCode = 685
	Kiribati                     CountryCode = 686
	NewCaledonia                 CountryCode = 687
	Tuvalu                       CountryCode = 688
	FrenchPolynesia              CountryCode = 689
	Tokelau                      CountryCode = 690
	Micronesia                   CountryCode = 691
	MarshallIslands              CountryCode = 692
	NorthKorea                   CountryCode = 850
	HongKong                     CountryCode = 852
	Macau                        CountryCode = 853
	Cambodia                     CountryCode = 855
	Laos                         CountryCode = 856
	Bangladesh                   CountryCode = 880
	Taiwan                       CountryCode = 886
	Maldives                     CountryCode = 960
	Lebanon                      CountryCode = 961
	Jordan                       CountryCode = 962
	Syria                        CountryCode = 963
	Iraq                         CountryCode = 964
	Kuwait                       CountryCode = 965
	SaudiArabia                  CountryCode = 966
	Yemen                        CountryCode = 967
	Oman                         CountryCode = 968
	Palestine                    CountryCode = 970
	UnitedArabEmirates           CountryCode = 971
	Israel                       CountryCode = 972
	Bahrain                      CountryCode = 973
	Qatar                        CountryCode = 974
	Bhutan                       CountryCode = 975
	Mongolia                     CountryCode = 976
	Nepal                        CountryCode = 977
	Tajikistan                   CountryCode = 992
	Turkmenistan                 CountryCode = 993
	Azerbaijan                   CountryCode = 994
	Georgia                      CountryCode = 995
	Kyrgyzstan                   CountryCode = 996
	Uzbekistan                   CountryCode = 998
)

func ToCallingCode(cc CountryCode) uint16 {
	switch cc {
	case Canada:
		return 1
	case Kazakhstan:
		return 7
	case SvalbardAndJanMayen:
		return 47
	case ChristmasIsland:
		return 61
	case Pitcairn:
		return 64
	case WesternSahara:
		return 212
	case NetherlandsAntilles:
		return 599
	}
	return uint16(cc)
}

// TelWithCountryCode 电话号码转化为存储模式
func TelWithCountryCode(cc CountryCode, tel string) string {
	ccs := strconv.FormatUint(uint64(cc), 10)
	return ccs + ":" + tel
}
func ParseTelWithCountryCode(t string, defaultCountryCode ...CountryCode) (cc CountryCode, tel string) {
	arr := strings.Split(t, ":")
	if len(arr) == 1 {
		tel = arr[0]
		if len(defaultCountryCode) == 1 {
			cc = defaultCountryCode[0]
		}
	} else {
		tel = arr[1]
		d, _ := strconv.Atoi(arr[0])
		cc = CountryCode(d)
	}
	return
}

// 转化成拨打模式
func TelWithCallingCode(cc CountryCode, tel string) string {
	c := ToCallingCode(cc)
	cs := strconv.FormatUint(uint64(c), 10)
	return cs + tel
}
