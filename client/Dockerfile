FROM golang:latest

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

COPY ./ .

RUN go build -o bin main.go

CMD [ "/app/bin"]


