# First stage: build the executable.
FROM golang:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/api main.go

FROM alpine:3.15.6
ENV HTTP_ADDRESS="0.0.0.0:8000"
ENV ALLOW_ORIGINS=""

ENV POSTGRES_HOST="dpg-cemc3dsgqg4ekmaufbr0-a"
ENV POSTGRES_PORT=5432
ENV POSTGRES_DATABASE="ecommerce_back"
ENV POSTGRES_USERNAME="daimer"
ENV POSTGRES_PASSWORD="Lw9hQ5lVQhQYMGjWQZcaCgsUjc3wntiU"

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/build/api .
COPY app.env .
EXPOSE 8000
CMD ["./api"]
