# AaGo

## 建议的项目文件夹结构

### 服务调用

* rservice 不能调用任何其他service, model/entity
* service 可以调用 rservice, aservice, 所有 mservice，不能调用 model/dao/po
* aservice 可以调用 rservice, 所有 model/dao/po
* mservice 可以调用 rservice, aservice, 所有 model/dao/po
* sdk 临时mservice之间调用，未来方便修改成RPC调用
* cns.Service 可以调用 ss.Service/bs.Service
* cns.Model 可以调用 ss.Model/bs.Model
* ss.Service 可以调用 bs.Service
* ss.Model 可以调用 bs.Model

> service 是不能调用 DTO 的，只有 controller 可以调用
> aservice, service, 都是 service，只是调用关系不同。 与 iris, model 关系紧密的，优先用 aservice；否则尽量用 service
> rservice 仅用于远程调用或调用第三方

### 文件结构

```text

AaGo
    + aa
    + adto
    + acache
    + ae     an error
    + aenum  an enum
    + afmt   a format
    + asql   
    + cnf    保留字，备用
    + com    中间件层，处理 request/response
    + crypt  编码加密
    + dict   系统字典
    + atype  autotype converters
    + docs   documents
    + util
      + healthcheck
    + lib
    + infr   infrastructure 基础设施

Application
  
    + app
        + app_name  微服务的某个服务 
            + cache       # 缓存
            + conf       # app_name 内写死的配置
            + dao        # Data Access Object 数据访问对象 --> 数据库表的映射
                + po    Persistant Object  持久对象。service/业务层与DAO中间的数据转换
            + bo          # Business Object，app_name 内通用数据传递对象
            + dic           # 放置翻译文件；   
            + enum      APP 内enum  # 放常量、枚举 conf  和  ienum 区别是： conf 纯服务端用到；ienum 客户端也需要用到
                        # 定时任务/后台任务   job/ cron/daemon     listener 需要后缀为 Listener.go 直接放到 service里面，用Listner后缀
            + job
                + queue
                    + channel
            + module                        # 提高微服务特性，module仅用于区分客户端、服务端、CMS端
                + syncUser                        # Service/Server 对服务端的接口
                + cms                       # 内容管理系统
                + bs                        # B/S架构，Browser/Server
                    + controller
                        + dto                   # 对外开放的，API Object
                        + vo                    # View Object
                    + model
                    + pad                    # pad view
                        + dto                   # 对外开放的
                        + vo
                    + pc                     + pc view
                        + dto                    
                    + phone                  # phone view
                        + dto                   
                    + task
                    service.go
                    xxxx   
                + ss          # S/S架构，Server/Server
 
            + service   # app_name 内通用 service  
        + app_name2 ....  其他微服务应用
        + router                 # 路由
            + middleware
        + rservice                  # remote service 其他远程服务或第三方服务（如微信、支付宝）
            + rpci                  # rpc interface ，对内提供的
    + grpc
        + gboot
        + protos                    # .proto
        + pb                        # protos 生成的文件
        + gservice

    + bootstrap     # 系统启动初始化
         + register      # 注册器
         + console       # 调试控制台、自定义命令（如Go自动生成文件指令）
    + deploy          # 配置及客户端代码部署源文件
        + config      # .ini 配置文件，cert.pem 文件
        + public
            + asset
        + views      # 模版文件
    + docs        # 说明文档
    + driver      # 驱动器
    + helper      # 快捷函数
    + sdk         # 封装的各“微服务”调用接口
    + storage     # 存储文件夹，日志、临时文件
        + docs
        + logs
    + tests                     #  测试

```
## 配置文件说明
配置文件分为几类：
* deploy/[xx|test|prod].ini  一些偏服务器的配置，如文件夹位置、各driver地址和端口等
** main.go --config=./deploy/xx.ini   程序启动时，就需要加载
* deploy/rsa    放置 ras 非对称加密密钥文件
** 文件位置在 ./deploy/xx.ini 中，启动时候，自动加载进内存
* a_ini 配置文件数据库， 一般放到数据库里，或者配置文件服务，负责存储一些偏运营侧配置，如第三方APP secret等
** 不同路由，加载不同的 a_ini，加载路由开始，需要引入 `pas.LoadAIni(app, "该路由加载的微服务名称，对应配置数据库key名关系")` 来加载远程配置文件到内存
* app_name/conf   某个微服务内部写死的配置

### ./deply/xx.ini 配置demo
```ini
env = dev                         ; 构建环境：dev test prep prod
timezone_id = Asia/Shanghai
time_format = 2006-02-01 15:04:05
mock = 1  ; 关闭mock，必须要把 route 里面的 mock 中间件删掉
rsa_root = ./deploy/config/rsa


[app]
log_file = ./storage/logs/app.log
crashlog_file = ./storage/logs/crash.log
config_root = /Users/iwi/proj/dockerfile/a-xixi/deploy/config
views_root = /Users/iwi/proj/dockerfile/a-xixi/deploy/view
asset_root = /Users/iwi/proj/dockerfile/a-xixi/deploy/asset

[a_oss]
src_root = ./storage/oss

[a_mall]
weixinpay_perm_root= /Users/iwi/proj/dockerfile/a-xixi/deploy/config/weixinpay

[biz_xixi]
port = 80

[biz_huodonger]
port = 8080       ; http 端口

[biz_luexu]
port = 8081

[mysql]
host = mysqldocker:3306
user = root
password = 
tls = false
timeout = 5s,5s,5s
pool_max_idle_conns = 0
pool_max_open_conns = 0
pool_conn_max_life_time = 0
pool_conn_max_idle_time = 0

; 基于微服务的架构，基础服务，以及各业务微服务数据库都必须要独立
[mysql_svc_gds]
host = xxx.xx.xx.xx:3306        ; 覆盖掉 [mysql] 上面 host 配置
schema = svc_gds              

[mysql_svc_log]
schema = svc_log

[mysql_svc_oss]
schema = svc_oss

[mysql_svc_pas]
host = xxx.xx.xx.xx:3306       ; 覆盖掉 [mysql] 上面 host 配置
schema = svc_pas

[mysql_svc_mall]
schema = svc_mall

[mysql_biz_xixi]
schema = biz_xixi

[mysql_biz_huodonger]
schema = biz_huodonger

[mysql_biz_luexu]
schema = biz_luexu


[redis]
host = redisdocker:6379
auth = 
tls = false
db = 0
timeout = 3s,3s,3s
pool_max_idle = 0
pool_max_active = 0
pool_idle_timeout = 0
pool_wait = false
pool_conn_life_time = 0

[redis_svc_pas]
host = xxx.xx.xx:6379     ; 覆盖掉 [redis] 上面 host 配置
db = 0

[redis_svc_gds]
db = 1

[redis_a_oss]
db = 2

[redis_svc_mall]
host = xxx.xx.xx:6379     ; 覆盖掉 [redis] 上面 host 配置
db = 0

[redis_mq]
db = 9
; redis read timeout (redis.DialReadTimeout()) 需要比 redigo pubsub  心跳时间长
; pubsub 心跳 时间是1分钟，这里就 70s。数据相对比较大，写时间相对延长
timeout = 3s,70s,10s


[redis_biz_xixi]
db = 11

[redis_biz_huodonger]
db = 12

[redis_biz_luexu]
db = 13


```
## 相关文档

* [iris wiki](https://github.com/kataras/iris/wiki)

## 错误说明

### 错误码

```json

//  DELETE /users/jack     deleter user record `jack`
{
  "code": 200,
  "msg": "OK",
  "data": null
}


// GET /users/jack      
{
  "code": 204,
  "msg": "No Content",
  "data": {}
}

{
  "code": 404,
  "msg": "Not Found",
  "data": null
}
```

* **code >=200 && code < 300 都表示成功。**
* code == 400 应当通过字典，将英文转为其他人类易读文本反馈给用户
* code == 401 access token 异常
* code == 404 Not Found 内容不存在
* code == 410 Gone 资源已经被删除
* code == 500 服务器错误，提示用户类似“服务器错误，请联系客服即可”

| code        | msg    |  说明  |
| --------   | -----:   | :----: |
| 200        |  OK     |   正常情况    |
| 204        |  No Content     |   返回空数组`[]`，表示列表为空的array  |

> GET 数组不存在，那么会返回 `204`；POST/PATCH/DELETE/PUT 操作成功，即使是没有内容返回，也会返回 `200`；
>> GET 请求，数据才是核心，code 并不重要；DELETE 等，code 才是判断是否成果的关键。

## 通用参数说明

目前支持用户上传数据为json或form表单数据，客户端可根据自己习惯自行选择；习惯使用json数据，需要带上 `Content-Type:application/json` 头；习惯form表单，客户端需要带上头 `Content-Type: application/x-www-form-urlencoded`

> 上传数据是不区分数据类型的，如 "uid": 10086 或 "uid": "10086" 都可以

* JS 可使用封装的 aa.js 函数，下列函数会自动带上_stringify=1，以及自动更新及带上客户端 access token，以beartoken方式传递
```javascript
Aa.Ajax({
    async: bool,
    method: string, // POST|GET|PUT|DELETE|...
    contentType: "application/json",
    dataType: "json",
    url: string,
    data: rdata,
    iSuccess: resolve,
    iAuth: ()=>{},  // 处理401未登录错误
    iError: reject  // 处理其他错误
})
Aa.JsonAjax({
    async: bool,
    method: string, // POST|GET|PUT|DELETE|...
    url: string,
    data: rdata,
    iSuccess: resolve,
    iAuth: ()=>{}, 
    iError: reject
})
Aa.AjaxResp((rdata, resolve, reject, async)=>{}, rdata) // 返回 null|resp
```

### 通用HEADER

* 请求BODY参数：（Content-Type:application/json JSON 体数据 或 Content-Type: application/x-www-form-urlencoded 表单数据）

### 通用参数

```txt
Pagination:
    users/page/{page:int}                第N页，每页最多20条
    users/page/{page:int}?limit=100      第N页，每页最多100条
    users?page=10             第N页
    users?offset=200&limit=100     从第offset（200）条数据开始，选择limit（100）条

Search: (start with `:`)
    name=Iwi                                          name=Iwi
    name=::Iwi:                                       name likes Iwi
    name=::Iwi                                        name ends with Iwi
    name=:Iwi:                                        name starts Iwi
    name=:Iwi,Tom                                     name in [Iwi, Tom]
    created_at=2019-06-01 00:00:00                       created_at = 2019-06-01 00:00:00
    created_at=:2019-06-01 00:00:00~2019-06-01 01:00:00  created_at >= 2019-06-01 00:00:00 && create_at < 2019-06-01 00:00:00
    created_at=:2019-06-01 00:00:00~                     created_at >= 2019-06-01 00:00:00
    created_at=:~2019-06-01 01:00:00                     created_at < 2019-06-01 00:00:00

```

### 客户端使用接口映射，不想要太多无关紧要的字段，或者为了防止日后接口减少字段而导致崩溃

可以在添加url param : _field=需要的字段名（逗号隔开）

```txt
GET http://host/users/xxx   获取用户所有属性
GET http://host/users/xxx?_field=name,age   只获取该用户name和age这两个字段
GET http://host/users  获取用户所有属性列表（数组）
GET http://host/users?_field=[name,age]  用户列表（数组）只保留name和age字段
```

### 弱类型语言，需要加 _stringify=1

服务端设计了大量uint64格式数据，超过了JS Number.MAX_VALUE，会出现数据失真的情况。故无法处理uint64的客户端，需要强制数据返回全是string类型

```txt
GET http://host/user?_stringify=1
```

## 命名规范
### mservice 里面的 mq_ 开头表示消息队列处理
```golang

// 处理收到支付成功信号
func (s *Service) ListenChannels() {
	ctx := aa.ContextWithTraceID(nil, "ch")
	ticker := time.NewTicker(time.Hour)
	s.app.Log.Debug(ctx, "Q--->[a_mall]<---Q")
	for {
		select {
		case <-ticker.C:
			s.app.Log.Debug(ctx, "Q--->[a_pas]<---Q")
		case sms := <-vericodeSmsChannel:
			s.goSendVericodeSms(ctx, sms)
		case log := <-vericodeSmsSendingLogChannel:
			ilog.New(s.app).PublishAVericodeSmsLog(ctx, log)
		case log := <-vericodeSmsVerificationChannel:
			ilog.New(s.app).PublishASmsVerificationLog(ctx, log)
		case u := <-cacheSimpleUserChannel:
			s.goCacheSimpleUser(ctx, u)
		case uid := <-refreshSimpleUserCacheChannel:
			s.goRefreshSimpleUserCache(ctx, uid)
		case msg := <-Qos0NotificationChannel:
			a_warning.New(s.app).GoSendL1WarningMsg(ctx, msg)
		}
	}
}
```
### 如果是跨微服务的微量消息队列，需要在 sdk 里面封装
可以根据情况使用 redis queue/ rabbitMQ / kafka 等
```golang 

// 这里禁止 a_mall 服务调用
// 这里会阻塞，需要用 go 协程
func (s *Service) Listen() {
	//ctx, cancel := context.WithCancel(context.Background())
	ctx := aa.ContextWithTraceID(nil, "mq")

	err := iorm.ListenRedisChannels(ctx, s.Redis, func() error {
		s.app.Log.Debug(ctx, "Q--->[abmallsub]<---Q")
		return nil
	}, func(channel string, msg []byte) error {
		// 必须要一直返回 nil，否则会终止
		switch channel {
		case ienum.RedisMqBmallPaySuccessChannel:
			var b do.BatchBill
			if err := json.Unmarshal(msg, &b); err != nil {
				s.app.Log.Error(ctx, err.Error())
				return nil
			}
			switch b.SvcId {
			case ienum.BizLuexu:
				// @TODO 增加自动尝试
				e := luexums.New(s.app).HandlePaidBillNotify(ctx, b)
				s.app.Try(ctx, e)
			}
		}
		return nil
	}, ienum.RedisMqBmallPaySuccessChannel)

	if err != nil {
		fmt.Println(err.Error())
		s.app.Log.Error(ctx, err.Error())
		panic(err.Error())
	}
}
```

###   _ 开头，表示临时变量

```go 
_pid, e1 := r.Query("pid", `^\d+$`)
pid, _ := _pid.Uint64()
```


### Controller Demo
```golang
func (c *Controller) PostFastBills(ictx iris.Context) {
	defer ictx.Next()
	r, resp, ctx := com.ReqResp(ictx)
	_sku, e0 := r.Body("sku_id", `^[1-9]\d*$`)                          // r.Body 表示获取 body 数据，区别于 r.Query
	_qty, e1 := r.Body("qty", `^\d+$`)
	data, e2 := r.Body("data", false)
	_missionId, e3 := r.Body("mission_id", `^\d+$`, false)
	_promoId, e4 := r.Body("promo_id", `^\d+$`, false)

	if e = resp.CatchErrors(e0, e1, e2, e3, e4); err != nil {       // 捕获 400 错误，并准确提示哪个参数错误
		c.app.Log.Info(ctx, err.Error())
		return
	}
	uid := iwibroker.SessionUid(ictx)                            // 通用获取客户端传递的bear token，解析后的uid
	fromUid := iwibroker.SessionFromUid(ictx)
	skuId := _sku.DefaultUint64(0)
	var cartItem do.CartItem
	cartItem.Checked = true
	cartItem.SkuId = skuId
	cartItem.Qty = _qty.DefaultUint16(1)
	cartItem.PromoId = _promoId.DefaultUint64(0)
	cartItem.Data = data.String()

	_, certs, skuVipType, e := mall.New(c.app).SimpleSpuBySkuId(ctx, skuId) // 可以使用的VIP类型，并不代表用户真实具有，即使具有也可能过期了
	if len(certs) != 0 {
		if e = c.s.CheckMyLocalCerts(ctx, uid, certs); e != nil {
			resp.WriteE(e)
			return
		}
	}

	// vip type 都是业务层的，服务层不存在VIP概念（只存在VIP价格）。所以放到业务层判断
	isVip := c.s.IsVip(ctx, uid, skuVipType)
	batch, e := mall.New(c.app).FastConfirmBill(ctx, conf.Biz, uid, isVip, cartItem, _missionId.DefaultUint16(0), fromUid, nil, 0, 0)
	if e != nil {
		resp.WriteE(e)
		return
	}
	xhost := c.app.Config.GetString(conf.Biz.Sid() + ".xhost.phone")
	pageUrl := xhost + "/payment/batch/" + strconv.FormatUint(batch, 10)
	prepay := dto.PrepayBill{
		Batch:   batch,
		PageUrl: pageUrl,
	}
	resp.Write(prepay)
}

```

## dao Demo
```golang
// 商品必须要全部记录进 Redis
type Sku struct {
	Id          uint64              `name:"id"`
	SpuId       uint64              `name:"spu_id"`
	Spec        atype.NullStringMap `name:"spec"`         // {颜色:白色, 内存:256G}  // 规格名称，如  512G 白色 => sku name = spu name + sku spec name
	GrossWeight uint32              `name:"gross_weight"` // 毛重 g
	Imgs        atype.NullImgSrcs   `name:"imgs"`
	Video       atype.NullVideoSrc  `name:"video"`
	Price       uint                `name:"price"`
	VipPrice    uint                `name:"vip_price"` // 会员价
	Ean13       uint64              `name:"ean13"`     // 13 位全球贸易项目代码，如果Ean13 小于一定阈值，表示采用某种计价方案
	Stock       uint32              `name:"stock"`     // 库存数量
	Status      ienum.SkuStatus     `name:"status"`
	CreatedAt   atype.Datetime      `name:"created_at"`
	UpdatedAt   atype.Datetime      `name:"updated_at"`
}

// 根据spu_id 拆表
func (t Sku) Table() string {
	return "sku_table_name_" + strconv.Format(t.SpuId,10)
}
// 这里在 asql 里面会优先选用 Indexes 里面的索引字段。如果省事的话，可不写这个函数，即不进行智能优先选用索引字段，自己手动处理筛选顺序。
func (t Sku)Indexes()[]string{              
    return []string{
        "id", "spu_id"
    }
}
```