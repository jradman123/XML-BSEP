FROM golang:latest as builder
WORKDIR /app

COPY ./connection_service/go.mod ./
COPY ./connection_service/go.sum ./
COPY /common ../common
RUN go mod download
COPY ./connection_service/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o connection-service

FROM alpine:latest

WORKDIR /root/



COPY --from=builder /app/connection-service .

EXPOSE 8084
CMD ["./connection-service"]