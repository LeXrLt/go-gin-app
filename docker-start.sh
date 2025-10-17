docker run --name go-gin-app \
 -it \
 -p 8080:8080 \
 -v $(pwd):/app go-gin-app bash