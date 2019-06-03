# AaGo

## 错误说明

### 错误码

```json

//  DELETE /user/jack     删除用户名为jack的用户
{
    "code": 200,
    "msg": "OK",
    "data": null
}


// GET /users/3000      获取第300页用户列表
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

```txt
分页：
    users/{page:int}     每页20，第N页
    users?offset=200&limit=100     从第offset（200）条数据开始，选择limit（100）条

搜索：
    name=Aario                                                          找到 name=Aario
    name=%Aario%                                                        name 包含Aario
    name=%Aario                                                         name 以Aario结尾
    name=Aario%                                                         name 以Aario开头
    name=Aario,Tom                                                      name=Aario || name=Tom
    create_at=2019-06-01 00:00:00                                       create_at = 2019-06-01 00:00:00
    create_at=2019-06-01 00:00:00~2019-06-01 01:00:00                   create_at >= 2019-06-01 00:00:00 && create_at < 2019-06-01 00:00:00
    create_at=2019-06-01 00:00:00~                                      create_at >= 2019-06-01 00:00:00
    create_at=~2019-06-01 01:00:00                                      create_at < 2019-06-01 00:00:00

```

### 客户端使用接口映射，不想要太多无关紧要的字段

可以在添加url param : _field=需要的字段名（逗号隔开）

```txt
GET http://host/user   获取用户所有属性
GET http://host/user?_field=name,age   只获取该用户name和age这两个字段
GET http://host/users  获取用户所有属性列表（数组）
GET http://host/users?_field=[name,age]  用户列表（数组）只保留name和age字段
```

### 弱类型语言，需要加 _weak=1

服务端设计了大量uint64格式数据，超过了JS Number.MAX_VALUE，会出现数据失真的情况。故无法处理uint64的客户端，需要强制数据返回全是string类型

```txt
GET http://host/user?_weak=1
```