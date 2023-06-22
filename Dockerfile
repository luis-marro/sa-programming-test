FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./interface/rest

WORKDIR /build

RUN cp /app/main .

EXPOSE 8080

CMD ["/build/main"]