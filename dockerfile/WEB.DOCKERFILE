FROM golang:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o web .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/web .
EXPOSE 8080
CMD ["./web"]
