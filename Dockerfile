# 使用官方 Go 镜像作为构建环境
FROM golang:1.23-alpine AS builder

ENV GOPROXY https://goproxy.cn,direct
# 设置工作目录
WORKDIR /app
# 复制源代码
COPY . .
# 下载依赖
RUN go mod download
# 编译 Go 应用
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/server server.go

FROM alpine:latest

# 安装必要的工具（如 ca-certificates）
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制可执行文件
COPY --from=builder /app/bin/server .

# 复制配置文件
COPY config.yaml ./config.yaml

# 暴露端口
EXPOSE 8081

# 启动应用
CMD ["./server"]