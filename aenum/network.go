package aenum

type Net uint8

const (
	UnknownNet Net = 0
	NetWIFI    Net = 1
	Tel2G      Net = 2
	Tel3G      Net = 3
	Tel4G      Net = 4
	Tel5G      Net = 5
	Tel6G      Net = 6
)

type Network uint16

/*
GPRS : 2G(2.5) General Packet Radia Name 114kbps
EDGE : 2G(2.75G) Enhanced Data Rate for GSM Evolution 384kbps
UMTS : 3G WCDMA 联通3G Universal Mobile Telecommunication System 完整的3G移动通信技术标准
CDMA : 2G 电信 Code Division Multiple Access 码分多址
EVDO_0 : 3G (EVDO 全程 CDMA2000 1xEV-DO) Evolution - Data Only (Data Optimized) 153.6kps - 2.4mbps 属于3G
EVDO_A : 3G 1.8mbps - 3.1mbps 属于3G过渡，3.5G
1xRTT : 2G CDMA2000 1xRTT (RTT - 无线电传输技术) 144kbps 2G的过渡,
HSDPA : 3.5G 高速下行分组接入 3.5G WCDMA High Speed Downlink Packet Access 14.4mbps
HSUPA : 3.5G High Speed Uplink Packet Access 高速上行链路分组接入 1.4 - 5.8 mbps
HSPA : 3G (分HSDPA,HSUPA) High Speed Packet Access
IDEN : 2G Integrated Dispatch Enhanced Networks 集成数字增强型网络 （属于2G，来自维基百科）
EVDO_B : 3G EV-DO Rev.B 14.7Mbps 下行 3.5G
LTE : 4G Long Term Evolution FDD-LTE 和 TDD-LTE , 3G过渡，升级版 LTE Advanced 才是4G
EHRPD : 3G CDMA2000向LTE 4G的中间产物 Evolved High Rate Packet Data HRPD的升级
HSPAP : 3G HSPAP 比 HSDPA 快些
*/
const (
	UnknownNetwork Network = 0
	WIFI           Network = 1

	// 2G
	TelOther2G Network = 200
	TelEDGE    Network = 201
	TelGPRS    Network = 202
	TelCDMA1x  Network = 203
	Tel1xRTT   Network = 204
	TelIDEN    Network = 205

	TelOther3G   Network = 300
	TelHSDPA     Network = 301
	TelHSUPA     Network = 302
	TelWCDMA     Network = 303
	TelCdmaEvdo0 Network = 304
	TelCdmaEvdoA Network = 305
	TelCdmaEvdoB Network = 306
	TelEHRPD     Network = 307
	TelHSPA      Network = 308
	TelHSPAP     Network = 309
	TelUMTS      Network = 310

	TelOther4G Network = 400
	TelLTE     Network = 401
)

func ToNet(network Network) Net {
	switch network {
	case WIFI:
		return NetWIFI
	case TelOther2G, TelEDGE, TelGPRS, TelCDMA1x, Tel1xRTT, TelIDEN:
		return Tel2G
	case TelOther3G, TelHSDPA, TelHSUPA, TelWCDMA, TelCdmaEvdo0, TelCdmaEvdoA, TelCdmaEvdoB, TelEHRPD, TelHSPA, TelHSPAP, TelUMTS:
		return Tel3G
	case TelOther4G, TelLTE:
		return Tel4G
	}
	return UnknownNet
}
