FROM golang:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download
RUN ls
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o generator ./script/generator/fake_user_generator.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/generator .
CMD ["./generator"]
