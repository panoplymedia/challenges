FROM golang:1.12.7

ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir /app

COPY ./main.go /app
COPY ./go.mod /app
COPY ./go.sum /app

COPY ./database /app/database
COPY ./migrations /app/migrations
COPY ./environment /app/environment
COPY ./models /app/models
COPY ./endpoints /app/endpoints

WORKDIR /app

RUN go build -i

ENTRYPOINT ["/app/sales-portal-api"]