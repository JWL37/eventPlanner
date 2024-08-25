FROM golang:latest

WORKDIR /contacts

COPY go.mod go.sum ./
RUN go mod download


COPY .. ./

RUN go build -o contacts cmd/main.go


CMD ["./contacts"]