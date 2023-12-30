#!/usr/bin/env sh

docker build -t "db" .
docker run -d -p 3306:3306 --name db db
