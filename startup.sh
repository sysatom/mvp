#!/usr/bin/env bash

cd docker
docker-compose up -d
cd ../
mysql -h 127.0.0.1 -P 3318 -u root -p123456 mvp < schema.sql
go run mock/faker.go
