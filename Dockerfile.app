FROM golang:1.24-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o /main ./app/cmd

FROM alpine:latest
COPY --from=builder /main /bin/main
EXPOSE 8080
CMD ["/bin/main"]
