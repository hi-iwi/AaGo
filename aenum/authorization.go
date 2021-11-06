package aenum

type Authorization string

const (
	// 适合server 2 server 简易接口验证   字段可以自定义（如 key， secret 自定义命名即可）
	ApiKey Authorization = "API Key"

	// Bearer Token（ JWT 令牌）
	//定义：为了验证使用者的身份，需要客户端向服务器端提供一个可靠的验证信息，称为Token，这个token通常由Json数据格式组成，通过hash散列算法生成一个字符串，所以称为Json Web Token
	BearerToken        Authorization = "Bearer Token"
	BasicAuth          Authorization = "Basic Auth"
	DigestAuth         Authorization = "Digest Auth"
	//OAuth1             Authorization = "OAuth 1.0"

	// 需要用户手动点授权，之后才获取用户信息
	OAuth2             Authorization = "OAuth 2.0"
	HawkAuthentication Authorization = "Hawk Authentication"
	AwsSignature       Authorization = "Aws Signature"
)
