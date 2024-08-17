FROM golang:latest

WORKDIR /eventPlanner

COPY go.mod go.sum ./
RUN go mod download


COPY ./ ./

RUN go build -o main cmd/main.go


CMD ["./main"]