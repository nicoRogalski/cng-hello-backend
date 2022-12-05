FROM golang:1.19 AS builder

# Used args for multi arch builds
ARG TARGETOS=linux
ARG TARGETARCH=amd64

# Set necessary environmet variables needed for our image
ENV GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH}
    #GOFLAGS=-buildvcs=false

# Add git or disable buildvcs flag in alpine builder image
#RUN apk add --no-cache git

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags "-s -w" -o app ./cmd/server

# Compress executable is optional and increases ressource usage
#RUN apk add upx
#RUN upx --best --lzma /build/app

# Build small secure image 
FROM alpine

ARG version

LABEL project.maintainer="nico.rogalski@gmsg-david.de"
LABEL project.name="CngHelloBackend"
LABEL project.version="${version}"

COPY --from=builder /build/app /

EXPOSE 8080
USER 1000

ENTRYPOINT ["/app"]