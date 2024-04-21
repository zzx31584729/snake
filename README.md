# Snake



### 安装golang相关环境
```shell
# 1. download golang installer 
# https://golang.google.cn/dl/go1.22.2.windows-amd64.msi

# 2. set go proxy
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

# 3. install swagger
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/swag  
```

### 编译项目
```shell
# 1. 生成文档
swag init -g cmd/snake/main.go
```

### 项目启动后，查看api文档
```shell
# swag api url
http://localhost:8080/swagger/index.html
```