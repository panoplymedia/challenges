FROM golang:1.12.7

ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir /app

COPY ./salesdata.csv /app/
COPY ./main.go /app/
COPY ./go.mod /app/
COPY ./go.sum /app/
COPY ./database /app/
COPY ./environment /app/
COPY ./auth /app/
COPY ./models /app/
COPY ./endpoints /app/

WORKDIR /app

RUN go build -a

ENTRYPOINT ["/app/api"]