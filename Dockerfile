# 使用官方 Go 镜像（基于 Debian）
FROM golang:1.24

# 设置工作目录
WORKDIR /app

# 复制项目文件
COPY . .

# 下载依赖
RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
RUN go mod tidy
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["go", "run", "main.go"]
# CMD ["tail", "-f", "/dev/null"]
