FROM golang:1.18-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    GOFLAGS=-buildvcs=false

WORKDIR /build

COPY . .
RUN go mod download
RUN go test -v ./...
RUN go build -ldflags "-s -w" -o app ./cmd

RUN apk add upx
RUN upx --best --lzma /build/app

# Build a small image
FROM gcr.io/distroless/static AS final

LABEL maintainer="nico.rogalski@msg-david.de"
LABEL cng-hello-backend.project="CngHelloBackend"
LABEL cng-hello-backend.version="1.0.0"

COPY --from=builder /build/app /

EXPOSE 8080
USER 1000
ENTRYPOINT ["/app"]