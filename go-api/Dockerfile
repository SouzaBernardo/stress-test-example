FROM golang:1.23 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server .

FROM debian:bookworm-slim
COPY --from=builder /app/server /server
EXPOSE 8080
CMD ["/server"]