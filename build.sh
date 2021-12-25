#!/usr/bin/env bash 
set -xe

# install packages and dependencies
go get github.com/auth0/go-jwt-middleware@v1.0.0
go get github.com/aws/aws-sdk-go@v1.41.19
go get github.com/form3tech-oss/jwt-go@v3.2.3+incompatible
go get github.com/gin-contrib/cors@v1.3.1
go get github.com/gin-gonic/gin@v1.7.2
go get github.com/golang-jwt/jwt@v3.2.2+incompatible
go get github.com/google/uuid@v1.3.0
go get github.com/mattn/go-isatty@v0.0.12
go get github.com/ugorji/go/codec@v1.2.6
go get github.com/joho/godotenv@v1.4.0 

# build command
GOOS=linux GOARCH=amd64 go build -o bin/application application.go
