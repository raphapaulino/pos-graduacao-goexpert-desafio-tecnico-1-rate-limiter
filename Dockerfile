FROM golang:1.22.3 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ratelimiter .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/ratelimiter .
COPY .env .
EXPOSE 8080
# ENTRYPOINT ["./ratelimiter"]
CMD ["./ratelimiter"]