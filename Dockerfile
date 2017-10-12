# Using Multi Stage Build to first build the program
FROM golang:1.9 as builder
LABEL maintainer "Thomas MUNOZ <thomas.munoz30@gmail.com>"

WORKDIR /go/src/github.com/thomasmunoz13/bot

COPY . .

RUN set -x && \
    go get -d -v .

RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o bot .

# Building the Dashblue image
FROM alpine
LABEL maintainer "Thomas MUNOZ <thomas.munoz30@gmail.com>"

USER root

RUN apk --no-cache add ca-certificates

RUN apk add --update ffmpeg \
    && rm -rf /var/cache/apk/*

WORKDIR /root/

COPY sounds/ /root/sounds/

# Copy bot executable from builder
COPY --from=builder /go/src/github.com/thomasmunoz13/bot/bot .

#CMD ["./bot"]
ENTRYPOINT ["./bot"]