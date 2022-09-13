FROM golang:latest as builder
LABEL author="ductnn"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o loo .


FROM alpine:latest
LABEL author="ductnn"

RUN apk --no-cache add ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /root/
COPY --from=builder /app/loo .

CMD ["./loo"]
