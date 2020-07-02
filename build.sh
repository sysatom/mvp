#!/usr/bin/env bash

GOOS=linux GOARCH=amd64 go build -o ./docker/server/dist/server ./server
docker build -t mvp-server:latest docker/server
