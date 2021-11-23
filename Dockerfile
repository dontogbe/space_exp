# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
COPY . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download



RUN go build -o dns cmd/main.go

EXPOSE 8080

CMD [ "./dns" ]