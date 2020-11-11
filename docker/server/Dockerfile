# 构建：使用golang:1.13版本
FROM golang:alpine as build

LABEL maintainer="stone<shdzstone@gmail.com>"

# 容器环境变量添加
#CGO_ENABLED 指定cgo编译器是否启用
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct 
#    GOROOT=/usr/local/go \
#    GOPATH=/Users/stone/go

# 设置工作区
WORKDIR /go/release

# 把全部文件复制到/go/release目录
copy . .

# 编译：把/go/release/main.go编译成可执行的二进制文件，命名为irainserver
# -tags netgo指定使用go编译器，而不是cgo
#RUN go build -tags netgo -o irainserver main.go
RUN go build -o irainserver main.go

# 运行：使用alpine作为基础镜像
FROM alpine:latest

WORKDIR /root

# 在build阶段复制可执行的go二进制文件app
COPY --from=build /go/release/irainserver /root
COPY --from=build /go/release/ginServer/conf /root/ginServer/conf

# 容器向外提供服务的暴露端口
EXPOSE 8080

# 启动服务
CMD ["./irainserver"]

#ENTRYPOINT ["./irainserver"]