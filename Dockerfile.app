FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o /main ./app/cmd

FROM scratch  
COPY --from=builder /main /main

EXPOSE 8080
ENTRYPOINT ["/main"]
