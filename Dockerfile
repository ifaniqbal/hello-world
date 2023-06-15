FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o myapp

FROM debian
WORKDIR /app
COPY --from=builder /app/myapp .
EXPOSE 80
CMD ["./myapp"]