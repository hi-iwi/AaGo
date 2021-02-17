# code

 DES：数据加密标准，密钥偏短（56位）、生命周期短（避免被破解）。
3DES：密钥长度112位或168位，通过增加迭代次数提高安全性 。处理速度慢、密钥计算时间长、加密效率不高 。
AES：高级数据加密标准，能够有效抵御已知的针对DES算法的所有攻击 。密钥建立时间短、灵敏性好、内存需求低、安全性高 。

    在RSA加解密算法中提及到RSA加密明文会受密钥的长度限制，这就说明用RSA加密的话明文长度是有限制的，而在实际情况我们要进行加密的明文长度或许会大于密钥长度，这样一来我们就不得不舍去RSA加密了。对此，DES加密则没有此限制。

　　鉴于以上两点(个人观点)，单独的使用DES或RSA加密可能没有办法满足实际需求，所以就采用了RSA和DES加密方法相结合的方式来实现数据的加密。

　　其实现方式即：
    客户端随机生成 des密钥 
　　1、信息(明文)采用DES密钥加密。
　　2、使用RSA加密前面的DES密钥信息。
　　最终将混合信息进行传递。
　　而接收方接收到信息后：
　　1、用RSA解密DES密钥信息。
　　2、再用RSA解密获取到的密钥信息解密密文信息。
　　最终就可以得到我们要的信息(明文)。