# Using Multi Stage Build to first build the program
FROM golang:1.9 as builder
LABEL maintainer "Thomas MUNOZ <thomas.munoz30@gmail.com>"

WORKDIR /go/src/github.com/thomasmunoz13/bot

COPY . .

RUN set -x && \
    go get -d -v .

RUN GOOS=linux go build -installsuffix cgo -o bot .

# Building the Dashblue image
FROM ubuntu
LABEL maintainer "Thomas MUNOZ <thomas.munoz30@gmail.com>"

USER root

ENV TOKEN="MzY4MDUwODE2NzgwODYxNDQw.DMEpAA.Tfd-Evoc0zayXxjOoF28CxpDimU"

RUN apt-get -y update \
    && apt-get -y install ffmpeg \
    && apt-get -y install ca-certificates

WORKDIR /root/

COPY sounds/ /root/sounds/

# Copy bot executable from builder
COPY --from=builder /go/src/github.com/thomasmunoz13/bot/bot .

#CMD ["./bot"]
ENTRYPOINT ["./bot"]