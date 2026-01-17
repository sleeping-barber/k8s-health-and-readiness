FROM golang:1.21-alpine as builder
EXPOSE 8080

ENV CGO_ENABLED=0

WORKDIR /
COPY . .
RUN go build -o app main.go

FROM scratch
COPY --from=builder /app .
COPY --from=builder /templates /templates

ENTRYPOINT ["./app"]