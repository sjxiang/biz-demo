
# Note



```
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


# 代码生成

```shell
$ cd easy-note
...

$ kitex -module github.com/sjxiang/biz-demo  ./idl/note.proto
$ kitex -module github.com/sjxiang/biz-demo  ./idl/user.proto

```
### 2.Run Note RPC Server
```shell
cd cmd/note
sh build.sh
sh output/bootstrap.sh
```

### 3.Run User RPC Server
```shell
cd cmd/user
sh build.sh
sh output/bootstrap.sh
```

### 4.Run API Server
```shell
cd cmd/api
chmod +x run.sh
./run.sh
```

### 5.Jaeger 

visit `http://127.0.0.1:16686/` on  browser.

#### Snapshots

<img src="images/shot.png" width="2850"  alt=""/>

## Custom Error Code

Customise the response error code in the `errno` package.

```go
const (
    SuccessCode                = 0
    ServiceErrCode             = 10001
    ParamErrCode               = 10002
    UserAlreadyExistErrCode    = 10003
    AuthorizationFailedErrCode = 10004
)
```

Sample code : Replace the default error code for hertz-jwt authentication error with a custom error code.

```go
authMiddleware, _ := jwt.New(&jwt.HertzJWTMiddleware{
    Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
        c.JSON(code, map[string]interface{}{
            "code":    errno.AuthorizationFailedErrCode,
            "message": message,
        })
    },
    //Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
    //  c.JSON(code, map[string]interface{}{
    //      "code":    code,
    //      "message": message,
    //  })
    //}
})
```

## Deploy with docker

### 1.Setup Basic Dependence
```shell
docker-compose up
```

### 2.Get Default Network Gateway Ip
``docker-compose up`` will create a default bridge network for mysql,etcd and jaeger.
Get the gateway ip of this default network to reach three components.
```shell
docker inspect easy_note_default
```
![img.png](img.png)

### 3.Replace ip in Dockerfile
You can use gateway ip in ``step 2`` to replace MysqlIp , EtcdIp and JAEGER_AGENT_HOST.

* UserDockerfile:
  ```dockerfile
  FROM golang:1.17.2
  ENV GO111MODULE=on
  ENV GOPROXY="https://goproxy.io"
  ENV MysqlIp="your MysqlIp"
  ENV EtcdIp="your EtcdIp"
  ENV JAEGER_AGENT_HOST="your JAEGER_AGENT_HOST"
  ENV JAEGER_DISABLED=false
  ENV JAEGER_SAMPLER_TYPE="const"
  ENV JAEGER_SAMPLER_PARAM=1
  ENV JAEGER_REPORTER_LOG_SPANS=true
  ENV JAEGER_AGENT_PORT=6831
  WORKDIR $GOPATH/src/easy_note
  COPY . $GOPATH/src/easy_note
  WORKDIR $GOPATH/src/easy_note/cmd/user
  RUN ["sh", "build.sh"]
  EXPOSE 8889
  ENTRYPOINT ["./output/bin/demouser"]
  ```

* NoteDockerfile:
  ```dockerfile
  FROM golang:1.17.2
  ENV GO111MODULE=on
  ENV GOPROXY="https://goproxy.io"
  ENV MysqlIp="your MysqlIp"
  ENV EtcdIp="your EtcdIp"
  ENV JAEGER_AGENT_HOST="your JAEGER_AGENT_HOST"
  ENV JAEGER_DISABLED=false
  ENV JAEGER_SAMPLER_TYPE="const"
  ENV JAEGER_SAMPLER_PARAM=1
  ENV JAEGER_REPORTER_LOG_SPANS=true
  ENV JAEGER_AGENT_PORT=6831
  WORKDIR $GOPATH/src/easy_note
  COPY . $GOPATH/src/easy_note
  WORKDIR $GOPATH/src/easy_note/cmd/note
  RUN ["sh", "build.sh"]
  EXPOSE 8888
  ENTRYPOINT ["./output/bin/demonote"]
  ```

* ApiDockerfile:
  ```dockerfile
  FROM golang:1.17.2
  ENV GO111MODULE=on
  ENV GOPROXY="https://goproxy.io"
  ENV MysqlIp="your MysqlIp"
  ENV EtcdIp="your EtcdIp"
  ENV JAEGER_AGENT_HOST="your JAEGER_AGENT_HOST"
  ENV JAEGER_DISABLED=false
  ENV JAEGER_SAMPLER_TYPE="const"
  ENV JAEGER_SAMPLER_PARAM=1
  ENV JAEGER_REPORTER_LOG_SPANS=true
  ENV JAEGER_AGENT_PORT=6831
  WORKDIR $GOPATH/src/easy_note
  COPY . $GOPATH/src/easy_note
  WORKDIR $GOPATH/src/easy_note/cmd/api
  RUN go build -o main .
  EXPOSE 8080
  ENTRYPOINT ["./main"]
  ```

### 4.Build images from Dockerfile
```shell
docker build -t easy_note/user -f UserDockerfile .
docker build -t easy_note/note -f NoteDockerfile .
docker build -t easy_note/api -f ApiDockerfile .
```

### 5.Run containers
* Create bridge network for these three services.
  ```shell
  docker network create -d bridge easy_note
  ```
* Run contains in ``easy_note`` network.
  ```shell
  docker run -d --name user --network easy_note easy_note/user
  docker run -d --name note --network easy_note easy_note/note
  docker run -d -p 8080:8080 --name api --network easy_note easy_note/api
  ```