FROM golang:latest

RUN mkdir -p /app
WORKDIR /app

ADD . /app

# ENV GO111MODULE=on

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

# RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main.go /app/cmd/api/main.go
# RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /dist/cron /app/cmd/cron/main.go
RUN go build -o main .
# RUN apk add --update gcc make build-base

EXPOSE 8080

CMD ["tail -f"]