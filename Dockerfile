# syntax=docker/dockerfile:1

FROM golang:1.18-alpine as build
MAINTAINER Jefri Herdi Triyanto, jefri.triyanto@goapotik.com

#-> Setup Environment
RUN export GO111MODULE=on &&\
    export CGO_ENABLED=0 &&\
    export GOOS=linux &&\
    export GOARCH=amd64 &&\
    export CGO=0

#-> 🌊 Install Require
RUN apk add --no-cache \
    gcc \
    musl-dev

WORKDIR /build
COPY . .

#-> 🌊 Install Golang Module
RUN go mod download

#-> ⚒️ Build
RUN go build -o ./run

#-> 🚫 Remove Project Files & Folders
RUN ls -a &&\
    rm -rf .git || true &&\
    rm -rf .vscode &&\
    rm -rf git_assets || true &&\
    rm -rf proto &&\
    rm -rf src &&\
    rm -rf tests || true &&\
    rm __debug_bin || true &&\
    rm __debug_bin.exe || true &&\
    rm .gitignore &&\
    rm application || true &&\
    rm application.exe || true &&\
    rm docker-compose.yaml &&\
    rm Dockerfile &&\
    rm go.mod &&\
    rm go.sum &&\
    rm kita || true &&\
    rm kita.exe || true &&\
    rm kita.app || true &&\
    rm *.db || true &&\
    rm main.go &&\
    rm README.md

#-> 💯 Configuration Environment (change target env)
# change this IP if your docker network difference
RUN sed -i 's/127.0.0.1/172.17.0.1/g' .env.development
RUN sed -i 's/localhost/172.17.0.1/g' .env.development

# 🚀 Finishing !!
FROM alpine:latest
WORKDIR /app
COPY --from=build /build/* .
RUN apk add --no-cache openssl curl nano
RUN chmod +x ./run
ENTRYPOINT ["/app/run"]
CMD ["run"]
