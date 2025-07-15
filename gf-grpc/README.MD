# gf grpc demo

## MySQL 创建设备表

拉起 `mysql` 容器

```shell
docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=123456 --restart=always bitnami/mysql:9.3.0
```

导入设备表

```sql
CREATE TABLE `device` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` bigint unsigned NOT NULL DEFAULT '0' COMMENT '软删除',
  `name` varchar(255) DEFAULT NULL COMMENT '名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='设备表';
```

## 生成相关 `entity` 和接口代码

编辑 `hack/config.yaml`

```yaml
# CLI tool, only in development environment.
# https://goframe.org/pages/viewpage.action?pageId=3673173
gfcli:
  gen:
    pbentity:
      - link:   "mysql:root:123456@tcp(127.0.0.1:3306)/test?loc=Local&parseTime=true"
        tables: "device"
        gJsonSupport: "true"
    
    dao:
      - link:   "mysql:root:123456@tcp(127.0.0.1:3306)/test?loc=Local&parseTime=true"
        tables: "device"
        gJsonSupport: "true"
```

生成代码

```shell
make pbentity
make dao
```

编写 `manifest/protobuf/device/v1/device.proto` 文件

```proto
syntax = "proto3";

package device;

option go_package = "gf-grpc/api/device/v1";

service Device{
    rpc Register(RegisterReq) returns (RegisterResp) {}
}

message RegisterReq{
    string name = 1;
}

message RegisterResp{
    int64 deviceId = 1;
}
```

生成接口代码

```shell
make pb
```

编写 `internal/logic/device` 业务代码

生成 service 代码

```shell
make service
```

编写 `internal/controller/device` 代码

注册服务 `cmd.go` 并在 `dev` 环境开启服务反射

```go
c := grpcx.Server.NewConfig()
c.Options = append(c.Options, []grpc.ServerOption{
  grpcx.Server.ChainUnary(
    grpcx.Server.UnaryValidate,
  ),
}...)
s := grpcx.Server.New(c)
device.Register(s)

isDev := g.Cfg().MustGet(ctx, "server.mode", "dev").String() == "dev"
if isDev {
  glog.Info(ctx, "gRPC server is running in development mode, enabling reflection.")
  reflection.Register(s.Server)
}

```


## 启动服务

编写配置文件 `manifest/config/config.yaml`

```yaml
server:
  mode: "dev"

grpc:
  name: "gf-grpc"
  address: ":58080"

logger:
  level : "all"
  stdout: true

database:
  logger:
    path:    "logs"
    file: "db-{Y-m-d}.log"
    level:   "all"
    stdout:  true
    rotateSize: "10M"
    rotateBackupLimit: 2
    rotateBackupExpire: "7d"
    rotateBackupCompress:  9
  default:
    - link:  "mysql:root:123456@tcp(127.0.0.1:3306)/test?loc=Local&parseTime=true"
      debug: true
```

```shell
# 
go run .
```

## 测试

```shell
# 安装 grpcurl 或者使用 insomnia 图形客户端 brew install insomnia
brew install grpcurl

# 查看所有服务
grpcurl -plaintext 127.0.0.1:58080 list

# 访问服务
grpcurl -d '{"name": "test"}' -plaintext 127.0.0.1:58080 device.Device/Register
```