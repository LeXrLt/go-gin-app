docker run --name go-gin-app \
 --rm -it \
 -p 8080:8080 \
 -v $(pwd):/app go-gin-app bash