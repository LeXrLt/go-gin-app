# 使用官方 Go 镜像（基于 Debian）
FROM golang:1.21

# 设置工作目录
WORKDIR /app

# 复制项目文件
COPY . .

# 下载依赖
RUN go mod download

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["go", "run", "main.go"]
# CMD ["tail", "-f", "/dev/null"]
