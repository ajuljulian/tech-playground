go get github.com/labstack/echo/v4
go get github.com/go-redis/redis/v8
go get -u gorm.io/gorm


build the image:
```
$ docker build -t my-go-app .
```

run a container from the image:
```
$ docker run -p 1323:1323 -it my-go-app
```

# Docker Compose


```
$ docker-compose up -d --build
$ docker-compose down
```

-d: run in the background
--build: update the container before starting


# Client

```
$ npx create-react-app client 
```

# Docker Commands

## List containers
```
$ docker container ls
```

## Connect to container
```
 $ docker exec -it d3fbcb83abb7 /bin/sh 
```