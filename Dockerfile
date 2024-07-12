FROM golang:latest

WORKDIR /RE_WEBSITE

COPY go.mod go.sum ./
RUN go mod download


COPY ./ ./

RUN go build -o main cmd/main.go


CMD ["./main"]