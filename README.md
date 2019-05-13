# AaGo

## 客户端使用接口映射，不想要太多无关紧要的字段

可以在添加url param : _field=需要的字段名（逗号隔开）

```txt
GET http://luexu.com/user   获取用户所有属性
GET http://luexu.com/user?_field=name,age   只获取该用户name和age这两个字段
GET http://luexu.com/users  获取用户所有属性列表（数组）
GET http://luexu.com/users?_field=[name,age]  用户列表（数组）只保留name和age字段
```

## 弱类型语言，需要加 _weak=1

服务端设计了大量uint64格式数据，超过了JS Number.MAX_VALUE，会出现数据失真的情况。故无法处理uint64的客户端，需要强制数据返回全是string类型

```txt
GET http://luexu.com/user?_weak=1
```