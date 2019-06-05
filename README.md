# AaGo

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

* **code >=200 && code < 300   都表示成功。**
* code == 400 应当通过字典，将英文转为其他人类易读文本反馈给用户
* code == 401 access token 异常
* code == 404 内容不存在
* code == 500 服务器错误，提示用户类似“服务器错误，请联系客服即可”

| code        | msg    |  说明  |
| --------   | -----:   | :----: |
| 200        |  OK     |   正常情况    |
| 204        |  No Content     |   返回空数组`[]`，表示列表为空的array  |

> GET 数组不存在，那么会返回 `204`；POST/PATCH/DELETE/PUT 操作成功，即使是没有内容返回，也会返回 `200`；
>> GET 请求，数据才是核心，code 并不重要；DELETE 等，code 才是判断是否成果的关键。

## 通用参数说明

目前支持用户上传数据为json或form表单数据，客户端可根据自己习惯自行选择

### 通用HEADER

* 请求BODY参数：（Content-Type:application/json  JSON 体数据 或 Content-Type: application/x-www-form-urlencoded 表单数据）

### 通用参数

```txt
Pagination:
    users/{page:int}                第N页，每页最多20条
    users/{page:int}?limit=100      第N页，每页最多100条
    users?offset=200&limit=100     从第offset（200）条数据开始，选择limit（100）条

Search: (start with `:`)
    name=Aario                                          name=Aario
    name=:%Aario%                                       name likes Aario
    name=:%Aario                                        name ends with Aario
    name=:Aario%                                        name starts Aario
    name=:Aario,Tom                                     name in [Aario, Tom]
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

## 默认值

* 日期默认值：1000-01-01 00:00:00   表示空日期
  * MySQL 8.0 datetime 区间 ['1000-01-01 00:00:00.000000', '9999-12-31 23:59:59.999999']
  * https://dev.mysql.com/doc/refman/8.0/en/datetime.html
