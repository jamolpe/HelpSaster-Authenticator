FROM golang:latest

LABEL maintainer="Javier Molpeceres <jamolpe@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o authorization-service ./cmd/...

EXPOSE 8080

CMD ["./authorization-service"]
