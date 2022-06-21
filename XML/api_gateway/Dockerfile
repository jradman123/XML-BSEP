FROM golang:latest as builder
WORKDIR /app

COPY  ./api_gateway/go.mod ./api_gateway/go.sum ./
COPY /common ../common
RUN go mod download
COPY ./api_gateway/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api-gateway

######## Start a new stage from scratch #######
FROM alpine:latest

WORKDIR /root/


COPY --from=builder /app/api-gateway .

EXPOSE 9090
CMD ["./api-gateway"]