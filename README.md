# pkey_svr
golang 实现的获取 私钥的http服务器

## 运行
`./pkey_svr --port 8080 -ip 127.0.0.1
> HTTP Server监听`127.0.0.1:8080`

## HTTP接口
`/get_pkey`

* 请求方法 `POST`   
* 请求Body及返回值均为`json`

### 请求值
```json
{
    "tk": "08a6a8f98f2c6312089219c9b1eeda81",
    "sig": "hostname"
}
```
|  参数名 |     类型    |     含义     |        备注       |
|:-------:|:-----------:|:------------:|:-----------------:|
|   tk  | string     |    请求token                   |                     |
|   sig     | string     |    签名标记                   | ||

### 返回值
```json
{
  "code": 0,
  "message": "uccess",
  "data": "msgpack 加密串"
}
```

|  参数名 |     类型    |     含义     |        备注       |
|:-------:|:-----------:|:------------:|:-----------------:|
|   code  |     int     |    状态码    | 0: 成功 非0: 失败 |
| message |    string   | 状态描述信息 |                   |
|   data  | object, null |   msgpack 加密串   |                   ||


## 项目部署

```
$ export GOPATH=/data/daemon/golang/
$ cd /data/daemon/golang/src/pkey_svr
$ go get github.com/vmihailenco/msgpack
$ go build && rm -f /data/daemon/release/bin/pkey_svr && mv pkey_svr /data/daemon/release/bin
```

- Supervisor配置

```
[program:pkey_svr]
# 进程要执行的命令
command=/data/daemon/release/bin/pkey_svr --port 8080 -ip 127.0.0.1
directory=/data/daemon/release/bin
user=www

# 自动重启
autorestart=true
# 日志路径
stdout_logfile=/data/logs/supervisor/web_scraping/pkey_svr.log
loglevel=error
```