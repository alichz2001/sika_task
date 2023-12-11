FROM golang:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migrator ./script/migrator/migrator.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/migrator .
COPY --from=builder /app/db.sql ./db.sql
CMD ["./migrator"]
