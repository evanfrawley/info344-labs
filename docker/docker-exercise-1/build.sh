#!/usr/bin/env bash
GOOS=linux go build
docker build -t docker-exercise-1 .
docker run -d -p 4000:4000 --name ex1 docker-exercise-1
go clean
