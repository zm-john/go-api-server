## Server

### 运行服务
#### 配置
* 复制 `config.ini.example` 文件，并命名为 `config.ini`
* 打开 `config.ini` 并进行配置
#### 执行数据库迁移
```go
go run main.go migrate
```
#### 开启服务
```go
go run main.go serve
```

### 生成日志
* 安装 `go-swagger`
> go install github.com/go-swagger/go-swagger/cmd/swagger
> swagger=$GOPATH/bin/swagger
* 生成 RESTApi 文档
> swagger generate spec -o ./swagger.yml
* 查看文档
> swagger serve swagger.yml
