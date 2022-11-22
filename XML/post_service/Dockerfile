FROM golang:latest as builder
WORKDIR /app

COPY ./post_service/go.mod ./
COPY ./post_service/go.sum ./
COPY /common ../common
COPY /monitoring ../monitoring
RUN go mod download
COPY ./post_service/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o post-service

FROM alpine:latest

WORKDIR /root/



COPY --from=builder /app/post-service .

EXPOSE 8083
CMD ["./post-service"]