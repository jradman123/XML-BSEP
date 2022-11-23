FROM golang:latest as builder
WORKDIR /app

COPY ./message_service/go.mod ./
COPY ./message_service/go.sum ./
COPY /common ../common
COPY /monitoring ../monitoring
RUN go mod download
COPY ./message_service/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o message_service

FROM alpine:latest

WORKDIR /root/



COPY --from=builder /app/message_service .

EXPOSE 8085
CMD ["./message_service"]