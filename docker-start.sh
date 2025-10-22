docker run --name go-gin-app \
 --restart=unless-stopped -td \
 -e GIN_MODE=release \
 -e TZ=Asia/Shanghai \
 -p 8080:8080 \
 -v $(pwd):/app go-gin-app:latest