FROM golang:latest
MAINTAINER  pibing

WORKDIR /app
COPY . .

ENV GO111MODULE=on \
    GOPROXY="https://goproxy.io"
RUN go build -o http-server main.go

CMD ["./http-server"]
EXPOSE 11111

