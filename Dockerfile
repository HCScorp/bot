# Using Multi Stage Build to first build the program
FROM golang:1.9 as builder
LABEL maintainer "Thomas MUNOZ <thomas.munoz30@gmail.com>"

WORKDIR /go/src/github.com/thomasmunoz13/bot

COPY . .

RUN set -x && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o bot .

# Building the Dashblue image
FROM alpine:3.6
LABEL maintainer "Thomas MUNOZ <thomas.munoz30@gmail.com>"

WORKDIR /root/

#RUN apk --no-cache add ca-certificates

# Copy bot executable from builder
COPY --from=builder /go/src/github.com/thomasmunoz13/bot/bot .

CMD ["./bot"]
