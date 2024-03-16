FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go get -u github.com/swaggo/swag/cmd/swag

RUN swag init

RUN go build -o main .

CMD ["/app/main"]
