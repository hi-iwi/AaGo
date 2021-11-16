# AaGo

## 命名规范

###   _ 开头，表示临时变量

```go 
_pid, e1 := r.Query("pid", `^\d+$`)
pid, _ := _pid.Uint64()
```

### UnsafeXXX 开头命名 model， 表示获取的字段不完整

```go 
func (m *Model) UnsafeAddrSubdivisions(ctx context.Context, cid aenum.CountryId, pid uint64) ([]entity.AddrDivision, *ae.Error) {
}
```

## 建议的项目文件夹结构

### 服务调用

* rservice 不能调用任何其他service, model/entity
* service 可以调用 rservice, aservice, 所有 mservice，不能调用 model/entity
* aservice 可以调用 rservice, 所有 model/entity
* mservice 可以调用 rservice, aservice, 所有 model/entity
* sdk 临时mservice之间调用，未来方便修改成RPC调用

> service 是不能调用 DTO 的，只有 controller 可以调用
> aservice, service, 都是 service，只是调用关系不同。 与 iris, model 关系紧密的，优先用 aservice；否则尽量用 service
> rservice 仅用于远程调用或调用第三方

### 文件结构

```text

AaGo
    + aa
    + adto
    + ae
    + cnf    备用
    + com    通信
    + crypt  编码加密
    + dict    系统字典
    + docs
    + format
    + queue
    + util
      + healthcheck
    + lib
    + infr   infrastructure 基础设施

Application
    + app
        + app_name
            + ado    # 动态数据对象。是 entity 的局部，entity 可以引用 ado。
            + cache                         # 缓存
            + dic           # 放置翻译文件；   
            + ienum              APP 内enum  # 放常量、枚举 conf  和  ienum 区别是： conf 纯服务端用到；ienum 客户端也需要用到
                                                       # 定时任务/后台任务   job/ cron/daemon     listener 需要后缀为 Listener.go 直接放到 service里面，用Listner后缀
            + module                        # 提高微服务特性，module仅用于区分客户端、服务端、CMS端
                + syncUser                        # Service/Server 对服务端的接口
                + cms                       # 内容管理系统
                + bs                        # B/S架构，Browser/Server
                    + controller
                        + idto                   # 对外开放的
                    + model
                    + pad                    # pad view
                        + idto                   # 对外开放的
                    + pc                     + pc view
                        + idto                   # 对外开放的
                    + phone                  # phone view
                        + idto                   # 对外开放的
            + entity
            + service
        + register
        + router
            + middleware
        + lservice    不同app间，service
        + rservice                  # remote service 其他远程服务或第三方服务（如微信、支付宝）
            + rpci                  # rpc interface ，对内提供的

    + grpc
        + gboot
        + protos                    # .proto
        + pb                        # protos 生成的文件
        + gservice

    + bootstrap     # 系统启动初始化
         + console       # 调试控制台、自定义命令（如Go自动生成文件指令）
    + conf          # .go 配置文件    conf  和  ienum 区别是： conf 纯服务端用到；ienum 客户端也需要用到
    + deploy
        + config      # .ini 配置文件，cert.pem 文件
        + public
            + asset
        + views      # 模版文件


   
    + storage
        + docs
        + logs
    + tests                     #  测试

    + vendor
    + build
    + base

    + src
    + dst
    + common
  
    + sdk

    + register
    + helper
    + facility
    + biz
```

## 相关文档

* [iris wiki](https://github.com/kataras/iris/wiki)

## 错误说明

### 错误码

```json

//  DELETE /user/jack     deleter user record `jack`
{
  "code": 200,
  "msg": "OK",
  "data": null
}


// GET /users/3000      show users on page 3000
{
  "code": 204,
  "msg": "No Content",
  "data": []
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

### 通用HEADER

* 请求BODY参数：（Content-Type:application/json JSON 体数据 或 Content-Type: application/x-www-form-urlencoded 表单数据）

### 通用参数

```txt
Pagination:
    users/{page:int}                第N页，每页最多20条
    users/{page:int}?limit=100      第N页，每页最多100条
    users?offset=200&limit=100     从第offset（200）条数据开始，选择limit（100）条

Search: (start with `:`)
    name=Iwi                                          name=Iwi
    name=::Iwi:                                       name likes Iwi
    name=::Iwi                                        name ends with Iwi
    name=:Iwi:                                        name starts Iwi
    name=:Iwi,Tom                                     name in [Iwi, Tom]
    create_at=2019-06-01 00:00:00                       create_at = 2019-06-01 00:00:00
    create_at=:2019-06-01 00:00:00~2019-06-01 01:00:00  create_at >= 2019-06-01 00:00:00 && create_at < 2019-06-01 00:00:00
    create_at=:2019-06-01 00:00:00~                     create_at >= 2019-06-01 00:00:00
    create_at=:~2019-06-01 01:00:00                     create_at < 2019-06-01 00:00:00

```

### 客户端使用接口映射，不想要太多无关紧要的字段，或者为了防止日后接口减少字段而导致崩溃

可以在添加url param : _field=需要的字段名（逗号隔开）

```txt
GET http://host/user   获取用户所有属性
GET http://host/user?_field=name,age   只获取该用户name和age这两个字段
GET http://host/users  获取用户所有属性列表（数组）
GET http://host/users?_field=[name,age]  用户列表（数组）只保留name和age字段
```

### 弱类型语言，需要加 _stringify=1

服务端设计了大量uint64格式数据，超过了JS Number.MAX_VALUE，会出现数据失真的情况。故无法处理uint64的客户端，需要强制数据返回全是string类型

```txt
GET http://host/user?_stringify=1
```