FROM golang:1.18-alpine

RUN mkdir "/app"
WORKDIR "/app"

COPY gin_oj.cn /app/app
ENTRYPOINT ["./app"]