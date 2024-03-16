FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init

RUN go build -o main .

CMD ["/app/main"]
