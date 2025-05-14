## goframe v2 nacos

### 拉起 nacos 服务

> 使用 rnacos

```shell
docker run -d \
    --restart always \
    --name nacos \
    -p 8848:8848 \
    -p 9848:9848 \
    -p 10848:10848 \
    qingpan/rnacos:stable-alpine

```

默认账号密码 `admin/admin`

在 `public` 命名空间 `user` 组下写入 `config.yaml` 

```yaml
{
    "server": {
        "address": "127.0.0.1:8080"
    }
}
```


### 代码

[main.go](main.go)

```shell
# 运行
go run main.go
```