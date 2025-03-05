FROM golang:1.23 AS builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o web ./cmd/web/main.go
ENTRYPOINT [ "./web" ]
CMD [ "8080" ]

#Final stage
FROM alpine:latest
COPY --from=builder /app/web /web
CMD [ "./web" ]
