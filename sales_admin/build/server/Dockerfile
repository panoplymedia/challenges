FROM golang:alpine as builder
COPY . /app
WORKDIR /app
RUN GO111MODULE=on go build -o /server ./cmd/server

FROM alpine
COPY --from=builder /server /server
ENTRYPOINT ["/server"]



