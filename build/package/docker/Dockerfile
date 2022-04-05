FROM golang:1.18-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

WORKDIR /build
RUN apk add upx

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test -v -buildvcs=false ./...

RUN go build -buildvcs=false -ldflags "-s -w" -o app ./cmd
RUN upx -9 /build/app

RUN addgroup -S app \
    && adduser -S -u 1000 -g app app

# Build a small image
FROM scratch

LABEL maintainer="nico.rogalski@msg-david.de"
LABEL cng-hello-backend.project="CngHelloBackend"
LABEL cng-hello-backend.version="1.0.0"

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /build/app /

EXPOSE 8080
USER 1000
ENTRYPOINT ["/app"]