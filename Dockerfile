FROM golang:1.20 AS build
WORKDIR /build
COPY go.mod go.sum /build/
RUN go mod download
COPY . /build/
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/beeapi-linux-amd64 ./main.go

FROM alpine:latest
COPY --from=build /tmp/beeapi-linux-amd64 /app/beeapi
WORKDIR /app
CMD ["/app/beeapi"]