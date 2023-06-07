ARG version

FROM golang:1.20-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GOOS=linux \
    GOARCH=amd64 \
    GOFLAGS=-buildvcs=false

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags "-s -w" -o cng-hello-backend ./cmd/cng-hello-backend

# RUN apk add upx
# RUN upx --best --lzma /build/cng-hello-backend

# Build a small image
FROM gcr.io/distroless/base AS final

ARG version

LABEL maintainer="nico.rogalski@msg-david.de"
LABEL cng-hello-backend.project="CngHelloBackend"
LABEL cng-hello-backend.version="${version}"

COPY --from=builder /build/cng-hello-backend /

EXPOSE 8080

ENTRYPOINT ["/cng-hello-backend"]