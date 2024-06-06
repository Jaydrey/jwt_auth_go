#!/bin/bash

go get github.com/githubnemo/CompileDaemon
go get github.com/joho/godotenv
go get -u github.com/golang-jwt/jwt/v5
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/gin-gonic/gin
go get -u gorm.io/driver/postgres
go get -u gorm.io/gorm

go install github.com/githubnemo/CompileDaemon

gpg --gen-random 30