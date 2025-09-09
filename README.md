# 构建镜像
docker build -t go-gin-app .

# 运行容器，映射端口
docker run --rm -p 8081:8080 -v $(pwd):/app go-gin-app
