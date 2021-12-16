# What It Is

This is a playground, of sort, to help me play around with different technologies.

The idea is that I will slowly keep adding different components and connect them together.

For now, I want to be able to deploy everything with docker-compose, but will eventually add the ability to deploy to Kubernetes as well.

# Deployment

### Deploy with docker-compose

```
$ docker-compose up -d --build
```

`-d`: start in non-interactive mode, e.g. get your shell prompt back :)

`--build`: build or rebuild the services

### Stop the containers
```
$ docker-compose down
```

# Services (Containers)

### Server

### What it does

I'm using the [Echo](https://echo.labstack.com) module to expose a `User` service.

I'm using the [Gorm](https://gorm.io) ORM module to persist users.

I'm using the [Go-Redis](https://github.com/go-redis/redis) Redis client to connect to a redis cluster.

### Third-party modules
```
go get github.com/labstack/echo/v4
go get github.com/go-redis/redis/v8
go get -u gorm.io/gorm
```

## Client

The client is a very basic React app that displays the list of users and provides CRUD functionality for users.

```
$ npx create-react-app client 
```

## Postgres

Postgres database

## Redis

Redis cluster

## Nginx

Nginx for web serving/routing etc.

# Next Steps

1. Add Redis for cache
1. Add Kafka
1. Add Flink
1. Deploy using Kubernetes

# Docker Commands

### List containers
```
$ docker container ls
```

### Connect to container interactively
```
 $ docker exec -it d3fbcb83abb7 /bin/sh 
```