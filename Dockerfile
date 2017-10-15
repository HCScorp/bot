# Using Multi Stage Build to first build the program
FROM golang:1.9.1 as builder
LABEL maintainer = "Thomas MUNOZ <thomas.munoz30@gmail.com>"

WORKDIR /go/src/github.com/HCScorp/bot

COPY . .

RUN set -x && \
    go get -d -v .

RUN GOOS=linux go build -installsuffix cgo -o bot .

# Building the Bot image
FROM ubuntu
LABEL maintainer = "Thomas MUNOZ <thomas.munoz30@gmail.com>"

USER root

ENV TOKEN="my-token"

RUN apt-get -y update \
    && apt-get -y install ffmpeg \
    && apt-get -y install ca-certificates

WORKDIR /root/

COPY sounds/ /root/sounds/

# Copy bot executable from builder
COPY --from=builder /go/src/github.com/HCScorp/bot/bot .

ENTRYPOINT ["./bot"]
