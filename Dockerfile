FROM golang:1.16-alpine as builder

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY templates/* templates/
EXPOSE 8080
CMD ["./main"]
