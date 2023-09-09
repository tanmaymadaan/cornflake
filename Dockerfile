FROM ubuntu:latest
LABEL authors="tanmaymadaan"

FROM golang:latest

#CMD ["dir", "ls"]

WORKDIR /app

#CMD ["dir", "ls"]

COPY . .
#CMD ["dir", "ls"]
#COPY .env .env

RUN go build -o main

EXPOSE 8080

CMD ["./main"]