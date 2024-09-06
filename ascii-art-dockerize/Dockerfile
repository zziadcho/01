FROM golang:alpine

WORKDIR /app

COPY . .

RUN apk add --no-cache bash

EXPOSE 8080

RUN go build main.go

CMD [ "./main" ]