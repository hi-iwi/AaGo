package aenum

import (
	"strconv"
	"strings"
)

type CountryId uint16

const (
	Canada                       CountryId = 50001
	America                      CountryId = 1
	Kazakhstan                   CountryId = 50007
	Russia                       CountryId = 7
	Egypt                        CountryId = 20
	SouthAfrica                  CountryId = 27
	Greece                       CountryId = 30
	Netherlands                  CountryId = 31
	Belgium                      CountryId = 32
	France                       CountryId = 33
	Spain                        CountryId = 34
	Hungary                      CountryId = 36
	Italy                        CountryId = 39
	Romania                      CountryId = 40
	Switzerland                  CountryId = 41
	Austria                      CountryId = 43
	UnitedKingdom                CountryId = 44
	Denmark                      CountryId = 45
	Sweden                       CountryId = 46
	Norway                       CountryId = 47
	SvalbardAndJanMayen          CountryId = 50047
	Poland                       CountryId = 48
	Germany                      CountryId = 49
	Peru                         CountryId = 51
	Mexico                       CountryId = 52
	Cuba                         CountryId = 53
	Argentina                    CountryId = 54
	Brazil                       CountryId = 55
	Chile                        CountryId = 56
	Colombia                     CountryId = 57
	Venezuela                    CountryId = 58
	Malaysia                     CountryId = 60
	Australia                    CountryId = 61
	ChristmasIsland              CountryId = 50061
	CocosIslands                 CountryId = 61
	Indonesia                    CountryId = 62
	Philippines                  CountryId = 63
	NewZealand                   CountryId = 64
	Pitcairn                     CountryId = 50064
	Singapore                    CountryId = 65
	Thailand                     CountryId = 66
	Japan                        CountryId = 81
	SouthKorea                   CountryId = 82
	Vietnam                      CountryId = 84
	China                        CountryId = 86
	Turkey                       CountryId = 90
	India                        CountryId = 91
	Pakistan                     CountryId = 92
	Afghanistan                  CountryId = 93
	SriLanka                     CountryId = 94
	Myanmar                      CountryId = 95
	Iran                         CountryId = 98
	SouthSudan                   CountryId = 211
	Morocco                      CountryId = 212
	WesternSahara                CountryId = 50212
	Algeria                      CountryId = 213
	Tunisia                      CountryId = 216
	Libya                        CountryId = 218
	Gambia                       CountryId = 220
	Senegal                      CountryId = 221
	Mauritania                   CountryId = 222
	Mali                         CountryId = 223
	Guinea                       CountryId = 224
	IvoryCoast                   CountryId = 225
	BurkinaFaso                  CountryId = 226
	Niger                        CountryId = 227
	Togo                         CountryId = 228
	Benin                        CountryId = 229
	Mauritius                    CountryId = 230
	Liberia                      CountryId = 231
	SierraLeone                  CountryId = 232
	Ghana                        CountryId = 233
	Nigeria                      CountryId = 234
	Chad                         CountryId = 235
	CentralAfricanRepublic       CountryId = 236
	Cameroon                     CountryId = 237
	CapeVerde                    CountryId = 238
	SaoTomeAndPrincipe           CountryId = 239
	EquatorialGuinea             CountryId = 240
	Gabon                        CountryId = 241
	RepublicOfTheCongo           CountryId = 242
	DemocraticRepublicOfTheCongo CountryId = 243
	Angola                       CountryId = 244
	GuineaBissau                 CountryId = 245
	BritishIndianOceanTerritory  CountryId = 246
	Seychelles                   CountryId = 248
	Sudan                        CountryId = 249
	Rwanda                       CountryId = 250
	Ethiopia                     CountryId = 251
	Somalia                      CountryId = 252
	Djibouti                     CountryId = 253
	Kenya                        CountryId = 254
	Tanzania                     CountryId = 255
	Uganda                       CountryId = 256
	Burundi                      CountryId = 257
	Mozambique                   CountryId = 258
	Zambia                       CountryId = 260
	Madagascar                   CountryId = 261
	Mayotte                      CountryId = 262
	Reunion                      CountryId = 262
	Zimbabwe                     CountryId = 263
	Namibia                      CountryId = 264
	Malawi                       CountryId = 265
	Lesotho                      CountryId = 266
	Botswana                     CountryId = 267
	Swaziland                    CountryId = 268
	Comoros                      CountryId = 269
	SaintHelena                  CountryId = 290
	Eritrea                      CountryId = 291
	Aruba                        CountryId = 297
	FaroeIslands                 CountryId = 298
	Greenland                    CountryId = 299
	Gibraltar                    CountryId = 350
	Portugal                     CountryId = 351
	Luxembourg                   CountryId = 352
	Ireland                      CountryId = 353
	Iceland                      CountryId = 354
	Albania                      CountryId = 355
	Malta                        CountryId = 356
	Cyprus                       CountryId = 357
	Finland                      CountryId = 358
	Bulgaria                     CountryId = 359
	Lithuania                    CountryId = 370
	Latvia                       CountryId = 371
	Estonia                      CountryId = 372
	Moldova                      CountryId = 373
	Armenia                      CountryId = 374
	Belarus                      CountryId = 375
	Andorra                      CountryId = 376
	Monaco                       CountryId = 377
	SanMarino                    CountryId = 378
	Vatican                      CountryId = 379
	Ukraine                      CountryId = 380
	Serbia                       CountryId = 381
	Montenegro                   CountryId = 382
	Kosovo                       CountryId = 383
	Croatia                      CountryId = 385
	Slovenia                     CountryId = 386
	BosniaAndHerzegovina         CountryId = 387
	Macedonia                    CountryId = 389
	CzechRepublic                CountryId = 420
	Slovakia                     CountryId = 421
	Liechtenstein                CountryId = 423
	FalklandIslands              CountryId = 500
	Belize                       CountryId = 501
	Guatemala                    CountryId = 502
	ElSalvador                   CountryId = 503
	Honduras                     CountryId = 504
	Nicaragua                    CountryId = 505
	CostaRica                    CountryId = 506
	Panama                       CountryId = 507
	SaintPierreAndMiquelon       CountryId = 508
	Haiti                        CountryId = 509
	SaintBarthelemy              CountryId = 590
	SaintMartin                  CountryId = 590
	Bolivia                      CountryId = 591
	Guyana                       CountryId = 592
	Ecuador                      CountryId = 593
	Paraguay                     CountryId = 595
	Suriname                     CountryId = 597
	Uruguay                      CountryId = 598
	Curacao                      CountryId = 599
	NetherlandsAntilles          CountryId = 50599
	EastTimor                    CountryId = 670
	Antarctica                   CountryId = 672
	Brunei                       CountryId = 673
	Nauru                        CountryId = 674
	PapuaNewGuinea               CountryId = 675
	Tonga                        CountryId = 676
	SolomonIslands               CountryId = 677
	Vanuatu                      CountryId = 678
	Fiji                         CountryId = 679
	Palau                        CountryId = 680
	WallisAndFutuna              CountryId = 681
	CookIslands                  CountryId = 682
	Niue                         CountryId = 683
	Samoa                        CountryId = 685
	Kiribati                     CountryId = 686
	NewCaledonia                 CountryId = 687
	Tuvalu                       CountryId = 688
	FrenchPolynesia              CountryId = 689
	Tokelau                      CountryId = 690
	Micronesia                   CountryId = 691
	MarshallIslands              CountryId = 692
	NorthKorea                   CountryId = 850
	HongKong                     CountryId = 852
	Macau                        CountryId = 853
	Cambodia                     CountryId = 855
	Laos                         CountryId = 856
	Bangladesh                   CountryId = 880
	Taiwan                       CountryId = 886
	Maldives                     CountryId = 960
	Lebanon                      CountryId = 961
	Jordan                       CountryId = 962
	Syria                        CountryId = 963
	Iraq                         CountryId = 964
	Kuwait                       CountryId = 965
	SaudiArabia                  CountryId = 966
	Yemen                        CountryId = 967
	Oman                         CountryId = 968
	Palestine                    CountryId = 970
	UnitedArabEmirates           CountryId = 971
	Israel                       CountryId = 972
	Bahrain                      CountryId = 973
	Qatar                        CountryId = 974
	Bhutan                       CountryId = 975
	Mongolia                     CountryId = 976
	Nepal                        CountryId = 977
	Tajikistan                   CountryId = 992
	Turkmenistan                 CountryId = 993
	Azerbaijan                   CountryId = 994
	Georgia                      CountryId = 995
	Kyrgyzstan                   CountryId = 996
	Uzbekistan                   CountryId = 998
)

func ToCallingCode(cc CountryId) uint16 {
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
func TelWithCountryCode(cc CountryId, tel string) string {
	ccs := strconv.FormatUint(uint64(cc), 10)
	return ccs + ":" + tel
}
func ParseTelWithCountryCode(t string, defaultCountryCode ...CountryId) (cc CountryId, tel string) {
	arr := strings.Split(t, ":")
	if len(arr) == 1 {
		tel = arr[0]
		if len(defaultCountryCode) == 1 {
			cc = defaultCountryCode[0]
		}
	} else {
		tel = arr[1]
		d, _ := strconv.Atoi(arr[0])
		cc = CountryId(d)
	}
	return
}

// 转化成拨打模式
func TelWithCallingCode(cc CountryId, tel string) string {
	c := ToCallingCode(cc)
	cs := strconv.FormatUint(uint64(c), 10)
	return cs + tel
}
