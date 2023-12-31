
# easy-note


```shell
$ tree
├─cmd
│  ├─api
│  │  ├─handlers
│  │  ├─middleware
│  │  └─rpc
│  ├─note
│  │  ├─dal      # 数据库操作 
│  │  │  └─db
│  │  ├─pack     # 数据包装
│  │  ├─rpc      # RPC 调用逻辑
│  │  └─service  # 主要业务逻辑
│  └─user
│      ├─dal
│      │  └─db
│      ├─pack
│      └─service
├─gen
│  └─pb
├─idl 
└─pkg
   ├─consts  # 常量
   ├─errno   # 自定义错误
   ├─jwt     # jwt 扩展
   └─script  # SQL 配置

```

## 架构图

```text
                                    http
                           ┌────────────────────────┐
 ┌─────────────────────────┤                        ├───────────────────────────────┐
 │                         │          api           │                               │
 │      ┌──────────────────►                        │◄──────────────────────┐       │
 │      │                  └───────────▲────────────┘                       │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                           resolve                                 │       │
 │      │                              │                                    │       │
req    resp                            │                                   resp    req
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                              │                                    │       │
 │      │                   ┌──────────▼─────────┐                          │       │
 │      │                   │                    │                          │       │
 │      │       ┌───────────►       Etcd         ◄─────────────────┐        │       │
 │      │       │           │                    │                 │        │       │
 │      │       │           └────────────────────┘                 │        │       │
 │      │       │                                                  │        │       │
 │      │     register                                           register   │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
 │      │       │                                                  │        │       │
┌▼──────┴───────┴───┐                                           ┌──┴────────┴───────▼─┐
│                   │───────────────── req ────────────────────►│                     │
│       note        │                                           │         user        │
│                   │◄──────────────── resp ────────────────────│                     │
└───────────────────┘                                           └─────────────────────┘
      protobuf                                                           protobuf
```


## 代码生成

```shell
$ cd easy-note

# 结合 pb 文件，考虑
$ protoc --go_out=. --go-grpc_out=. ./idl/easy_note.proto
$ protoc --go_out=. --go-grpc_out=. ./idl/user.proto

```




## 测试


### Register

```shell
curl --location --request POST '127.0.0.1:8080/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"kinggo",
    "password":"123456"
}'
```

#### response
```javascript
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}
// failed
{
    "code": 10003,
    "message": "User already exists",
    "data": null
}
```

### Login

#### will return jwt token
```shell
curl --location --request POST '127.0.0.1:8080/v1/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"kinggo",
    "password":"123456"
}'
```

#### response
```javascript
// successful
{
    "code": 0,
    "expire": "2022-01-19T01:56:46+08:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDI1Mjg2MDYsImlkIjoxLCJvcmlnX2lhdCI6MTY0MjUyNTAwNn0.k7Ah9G4Enap9YiDP_rKr5HSzF-fc3cIxwMZAGeOySqU"
}
// failed
{
    "code": 10004,
    "message": "Authorization failed",
    "data": null
}
```

### Create Note
```shell
curl --location --request POST '127.0.0.1:8080/v1/note' \
--header 'Authorization: Bearer $token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"test title",
    "content":"test content"
}'
```

#### response
```javascript
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}
// failed
{
    "code": 10002,
    "message": "Wrong Parameter has been given",
    "data": null
}
```

### Query Note
```shell
curl --location --request GET '127.0.0.1:8080/v1/note/query?offset=0&limit=20&search_keyword=test' \
--header 'Authorization: Bearer $token'
```

#### response
```javascript
// successul
{
    "code": 0,
    "message": "Success",
    "data": {
        "notes": [
            {
                "note_id": 1,
                "user_id": 1,
                "user_name": "kinggo",
                "user_avatar": "test",
                "title": "test title",
                "content": "test content",
                "create_time": 1642525063
            }
        ],
        "total": 1
    }
}
// failed
{
    "code":10002,
    "message":"Wrong Parameter has been given",
    "data":null
}
```

### Update Note
```shell
curl --location --request PUT '127.0.0.1:8080/v1/note/$note_id' \
--header 'Authorization: Bearer $token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title":"test",
    "content":"test"
}'
```

#### response
```javascript
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}
// failed
{
    "code":10001,
    "message":"strconv.ParseInt: parsing \"$note_id\": invalid syntax",
    "data":null
}
```

### Delete Note
```shell
curl --location --request DELETE '127.0.0.1:8080/v1/note/$note_id' \
--header 'Authorization: Bearer $token'
```

#### response
```javascript
// successful
{
    "code": 0,
    "message": "Success",
    "data": null
}
// failed
{
    "code":10001,
    "message":"strconv.ParseInt: parsing \"$note_id\": invalid syntax",
    "data":null
}
```


## 未完待续

```text
1. pb 生成统一 package name
2. 中间件 ...
```