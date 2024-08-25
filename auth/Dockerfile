FROM golang:latest

WORKDIR /auth

COPY go.mod go.sum ./
RUN go mod download


COPY .. ./

RUN go build -o auth cmd/main.go


CMD ["./auth"]