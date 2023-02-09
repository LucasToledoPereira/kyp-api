FROM golang:1.19-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN apk add --no-cache git gcc g++ libc-dev librdkafka-dev pkgconf

RUN go build -o server . 

CMD [ "/app/server" ]
