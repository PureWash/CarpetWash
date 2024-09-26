FROM golang:1.23.1 AS builder

WORKDIR /carpet-app

COPY . ./
RUN go mod download

COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../myapp

FROM alpine:latest

WORKDIR /carpet-app

COPY --from=builder /carpet-app/myapp .
COPY --from=builder /carpet-app/.env .

EXPOSE 7981

CMD [ "./myapp" ]