ARG version

FROM golang:1.19-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GOOS=linux \
    GOARCH=amd64 
    #GOFLAGS=-buildvcs=false

# Add git or disable buildvcs flag
RUN apk add --no-cache git

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags "-s -w" -o app ./cmd/server

# Compress executable is optional and increases ressource usage
#RUN apk add upx
#RUN upx --best --lzma /build/app

# Build small secure image 
FROM gcr.io/distroless/static-debian11

ARG version

LABEL project.maintainer="nico.rogalski@gmsg-david.de"
LABEL project.name="CngHelloBackend"
LABEL project.version="${version}"

COPY --from=builder /build/app /

EXPOSE 8080
USER 1000
ENTRYPOINT ["/app"]