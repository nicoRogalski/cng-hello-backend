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

# Build a small image
FROM gcr.io/distroless/static:latest

LABEL maintainer="nico.rogalski@msg-david.de"
LABEL cng-hello-backend.project="CngHelloBackend"
LABEL cng-hello-backend.version="1.0.0"

# RUN echo "adding run user go to system" \
#     && addgroup -S go -g 1000 \
#     && adduser -S go -u 1000 -G go

COPY --from=builder /build/app /

# USER 1000

EXPOSE 8080

CMD ["/app"]