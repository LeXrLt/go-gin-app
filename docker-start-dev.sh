docker stop go-gin-app
docker rm go-gin-app
docker run --name go-gin-app \
 --restart=unless-stopped -it \
 -e GIN_MODE=debug \
 -p 8081:8080 \
 -v $(pwd):/app go-gin-app:latest bash