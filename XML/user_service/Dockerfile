FROM golang:latest as builder
WORKDIR /app

COPY ./user_service/go.mod ./
COPY ./user_service/go.sum ./
COPY /common ../common
COPY /monitoring ../monitoring
RUN go mod download
COPY ./user_service/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-service

FROM alpine:latest

WORKDIR /root/


COPY --from=builder /app/user-service .

EXPOSE 8082
CMD ["./user-service"]