package aenum

import (
	"strconv"
)

type Continent uint8
type Country uint16

const (
	Asia         Continent = 1
	Europe       Continent = 2
	NorthAmerica Continent = 3
	SouthAmerica Continent = 4
	Oceania      Continent = 5
	Africa       Continent = 6
	Antarctica   Continent = 7
)

const (
	Canada        Country = 50001
	America       Country = 1
	Kazakhstan    Country = 997 // 2021年，哈萨克斯坦国际区号变更为 997
	Russia        Country = 7
	Egypt         Country = 20
	SouthAfrica   Country = 27
	Greece        Country = 30
	Netherlands   Country = 31
	Belgium       Country = 32
	France        Country = 33
	Spain         Country = 34
	Hungary       Country = 36
	Italy         Country = 39
	Romania       Country = 40
	Switzerland   Country = 41
	Austria       Country = 43
	UnitedKingdom Country = 44
	Denmark       Country = 45
	Sweden        Country = 46
	Norway        Country = 47
	Poland        Country = 48
	Germany       Country = 49
	Peru          Country = 51
	Mexico        Country = 52
	Cuba          Country = 53
	Argentina     Country = 54
	Brazil        Country = 55
	Chile         Country = 56
	Colombia      Country = 57
	Venezuela     Country = 58
	Malaysia      Country = 60
	Australia     Country = 61
	Indonesia     Country = 62
	Philippines   Country = 63
	NewZealand    Country = 64
	//Pitcairn                     Country = 50064
	Singapore                    Country = 65
	Thailand                     Country = 66
	Japan                        Country = 81
	SouthKorea                   Country = 82
	Vietnam                      Country = 84
	China                        Country = 86
	Turkey                       Country = 90
	India                        Country = 91
	Pakistan                     Country = 92
	Afghanistan                  Country = 93
	SriLanka                     Country = 94
	Myanmar                      Country = 95
	Iran                         Country = 98
	SouthSudan                   Country = 211
	Morocco                      Country = 212
	WesternSahara                Country = 50212
	Algeria                      Country = 213
	Tunisia                      Country = 216
	Libya                        Country = 218
	Gambia                       Country = 220
	Senegal                      Country = 221
	Mauritania                   Country = 222
	Mali                         Country = 223
	Guinea                       Country = 224
	IvoryCoast                   Country = 225
	BurkinaFaso                  Country = 226
	Niger                        Country = 227
	Togo                         Country = 228
	Benin                        Country = 229
	Mauritius                    Country = 230
	Liberia                      Country = 231
	SierraLeone                  Country = 232
	Ghana                        Country = 233
	Nigeria                      Country = 234
	Chad                         Country = 235
	CentralAfricanRepublic       Country = 236
	Cameroon                     Country = 237
	CapeVerde                    Country = 238
	SaoTomeAndPrincipe           Country = 239
	EquatorialGuinea             Country = 240
	Gabon                        Country = 241
	RepublicOfTheCongo           Country = 242
	DemocraticRepublicOfTheCongo Country = 243
	Angola                       Country = 244
	GuineaBissau                 Country = 245
	BritishIndianOceanTerritory  Country = 246
	Seychelles                   Country = 248
	Sudan                        Country = 249
	Rwanda                       Country = 250
	Ethiopia                     Country = 251
	Somalia                      Country = 252
	Djibouti                     Country = 253
	Kenya                        Country = 254
	Tanzania                     Country = 255
	Uganda                       Country = 256
	Burundi                      Country = 257
	Mozambique                   Country = 258
	Zambia                       Country = 260
	Madagascar                   Country = 261
	Mayotte                      Country = 262
	Reunion                      Country = 262
	Zimbabwe                     Country = 263
	Namibia                      Country = 264
	Malawi                       Country = 265
	Lesotho                      Country = 266
	Botswana                     Country = 267
	Swaziland                    Country = 268
	Comoros                      Country = 269
	SaintHelena                  Country = 290
	Eritrea                      Country = 291
	Aruba                        Country = 297
	FaroeIslands                 Country = 298
	Greenland                    Country = 299
	Gibraltar                    Country = 350
	Portugal                     Country = 351
	Luxembourg                   Country = 352
	Ireland                      Country = 353
	Iceland                      Country = 354
	Albania                      Country = 355
	Malta                        Country = 356
	Cyprus                       Country = 357
	Finland                      Country = 358
	Bulgaria                     Country = 359
	Lithuania                    Country = 370
	Latvia                       Country = 371
	Estonia                      Country = 372
	Moldova                      Country = 373
	Armenia                      Country = 374
	Belarus                      Country = 375
	Andorra                      Country = 376
	Monaco                       Country = 377
	SanMarino                    Country = 378
	Vatican                      Country = 379
	Ukraine                      Country = 380
	Serbia                       Country = 381
	Montenegro                   Country = 382
	Kosovo                       Country = 383
	Croatia                      Country = 385
	Slovenia                     Country = 386
	BosniaAndHerzegovina         Country = 387
	Macedonia                    Country = 389
	CzechRepublic                Country = 420
	Slovakia                     Country = 421
	Liechtenstein                Country = 423
	FalklandIslands              Country = 500
	Belize                       Country = 501
	Guatemala                    Country = 502
	ElSalvador                   Country = 503
	Honduras                     Country = 504
	Nicaragua                    Country = 505
	CostaRica                    Country = 506
	Panama                       Country = 507
	SaintPierreAndMiquelon       Country = 508
	Haiti                        Country = 509
	SaintBarthelemy              Country = 590
	SaintMartin                  Country = 590
	Bolivia                      Country = 591
	Guyana                       Country = 592
	Ecuador                      Country = 593
	Paraguay                     Country = 595
	Suriname                     Country = 597
	Uruguay                      Country = 598
	Curacao                      Country = 599
	NetherlandsAntilles          Country = 50599
	EastTimor                    Country = 670
	// Antarctica                   Country = 672  南极洲
	Brunei             Country = 673
	Nauru              Country = 674
	PapuaNewGuinea     Country = 675
	Tonga              Country = 676
	SolomonIslands     Country = 677
	Vanuatu            Country = 678
	Fiji               Country = 679
	Palau              Country = 680
	WallisAndFutuna    Country = 681
	CookIslands        Country = 682
	Niue               Country = 683
	Samoa              Country = 685
	Kiribati           Country = 686
	NewCaledonia       Country = 687
	Tuvalu             Country = 688
	FrenchPolynesia    Country = 689
	Tokelau            Country = 690
	Micronesia         Country = 691
	MarshallIslands    Country = 692
	NorthKorea         Country = 850
	HongKong           Country = 852
	Macau              Country = 853
	Cambodia           Country = 855
	Laos               Country = 856
	Bangladesh         Country = 880
	Taiwan             Country = 886
	Maldives           Country = 960
	Lebanon            Country = 961
	Jordan             Country = 962
	Syria              Country = 963
	Iraq               Country = 964
	Kuwait             Country = 965
	SaudiArabia        Country = 966
	Yemen              Country = 967
	Oman               Country = 968
	Palestine          Country = 970
	UnitedArabEmirates Country = 971
	Israel             Country = 972
	Bahrain            Country = 973
	Qatar              Country = 974
	Bhutan             Country = 975
	Mongolia           Country = 976
	Nepal              Country = 977
	Tajikistan         Country = 992
	Turkmenistan       Country = 993
	Azerbaijan         Country = 994
	Georgia            Country = 995
	Kyrgyzstan         Country = 996
	Uzbekistan         Country = 998
)

func NewCountry(id uint16) (Country, bool) {
	c := Country(id)
	return c, true
}
func (c Country) Uint16() uint16 { return uint16(c) }
func (c Country) String() string {
	return strconv.Itoa(int(c))
}
func (c Country) Is(x uint16) bool {
	return c.Uint16() == x
}
func (c Country) In(args ...Country) bool {
	for _, a := range args {
		if a == c {
			return true
		}
	}
	return false
}

func (c Country) CallingCode() string {
	switch c {
	case Canada:
		return "001"
	case NetherlandsAntilles:
		return "00599"
	case WesternSahara:
		return "00212"
	}
	return "00" + strconv.FormatUint(uint64(c), 10)
}

// 转化成拨打模式
func ToFullTel(c Country, tel string) string {
	return c.CallingCode() + " " + tel
}
