FROM golang:latest

WORKDIR /events

COPY go.mod go.sum ./
RUN go mod download


COPY .. ./

RUN go build -o events cmd/main.go


CMD ["./events"]