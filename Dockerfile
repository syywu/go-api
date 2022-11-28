FROM golang:1.19.3-alpine3.16

RUN mkdir /app

ADD . /app

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env /app

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]