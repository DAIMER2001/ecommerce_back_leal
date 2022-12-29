FROM golang:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./build/ecommerce main.go

FROM alpine:3.15.6
ENV HTTP_ADDRESS="0.0.0.0:8000"
ENV ALLOW_ORIGINS=""

ENV POSTGRES_HOST="localhost"
ENV POSTGRES_PORT=5432
ENV POSTGRES_USERNAME=""
ENV POSTGRES_PASSWORD=""
ENV POSTGRES_DATABASE="postgres"
ENV POSTGRES_MAX_CONNS=5

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/build/ecommerce .
COPY app.env .
EXPOSE 8000
CMD ["./api"]
