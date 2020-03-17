FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV DATABASE_URL=:5432

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
