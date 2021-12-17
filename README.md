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

* Expose a `User` service using the [Echo](https://echo.labstack.com) module

* Persisting users to a Postgres database using the [Gorm](https://gorm.io) ORM module

* Connecting to a redis cluster using the [Go-Redis](https://github.com/go-redis/redis) Redis client

* Reading and wring to a Kafka topic using the [Kafka-Go](https://github.com/segmentio/kafka-go) module 

### Third-party modules
```
go get github.com/labstack/echo/v4
go get github.com/go-redis/redis/v8
go get -u gorm.io/gorm
go get github.com/segmentio/kafka-go 
```

## Client

The client is a very basic React app that displays the list of users and provides CRUD functionality for users.

```
$ npx create-react-app client 
```

## Postgres

Postgres database

## Redis

Create a redis cluster to be used by other containers

## Nginx

Nginx for web serving/routing etc.

## Zookeeper

Used with Kafka

## Kafka

Create a kafka broker to be used by other containers

# Next Steps

1. Add Flink
1. Add CI/CD using Travis
1. Deploy to AWS Elastic Beanstalk
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