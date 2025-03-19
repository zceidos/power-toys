# 使用官方 Go 语言镜像作为基础镜像
FROM golang:1.24 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译应用程序
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 使用轻量级的镜像来运行应用程序
FROM alpine:latest

# 设置工作目录
WORKDIR /root/

# 复制编译后的二进制文件
COPY --from=builder /app/main .

# 复制 .env 文件
COPY --from=builder /app/.env .

# 暴露应用程序的端口
EXPOSE 8080

# 运行应用程序
CMD ["./main"]