FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV PG_HOST=ruby.db.elephantsql.com
ENV PG_PORT=5432
ENV PG_PASS=hNSdE5i665siQjFY9KK-R-gTRU4FAG7k
ENV PG_NAME=gcyotvfr
ENV PG_USER=gcyotvfr
ENV PG_DB=gcyotvfr


RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
