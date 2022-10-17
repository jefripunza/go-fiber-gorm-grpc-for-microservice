# syntax=docker/dockerfile:1

FROM golang:1.18-alpine
MAINTAINER Jefri Herdi Triyanto, jefri.triyanto@goapotik.com

#-> 🌊 Install Require
RUN apk add --no-cache gcc musl-dev nano

WORKDIR /app
COPY . .

#-> 🌊 Install Golang Module
RUN go mod download

#-> ⚒️ Build
RUN go build -o ./run

#-> 🚫 Remove Project Files & Folders
RUN ls -a &&\
    rm -rf .vscode &&\
    rm -rf proto &&\
    rm -rf src &&\
    rm __debug_bin.exe || true &&\
    rm .gitignore &&\
    rm docker-compose.yaml &&\
    rm Dockerfile &&\
    rm go.mod &&\
    rm go.sum &&\
    rm main_service.db || true &&\
    rm main.go &&\
    rm package.json || true &&\
    rm README.md

#-> 💯 Configuration Environment
# change this IP if your docker network difference
RUN sed -i 's/localhost/172.17.0.1/g' .env

# 🚀 Finish !!
CMD [ "./run" ]