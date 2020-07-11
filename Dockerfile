FROM golang:alpine as builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /build 
COPY go.mod /build/
COPY go.sum /build/
WORKDIR /build
RUN go mod download
ADD . /build/
RUN go build -o test .

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata && \
    cp /usr/share/zoneinfo/America/Argentina/Buenos_Aires /etc/localtime && \
    apk del tzdata && rm -rf /var/cache/apk/* && date

WORKDIR /app
COPY --from=builder /build/test /app/

CMD ["./test"]