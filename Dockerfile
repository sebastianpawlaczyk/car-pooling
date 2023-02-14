FROM alpine:3.8

# This Dockerfile is optimized for go binaries, change it as much as necessary
# for your language of choice.

RUN apk add --no-cache ca-certificates

EXPOSE 9091

COPY . .
 
ENTRYPOINT [ "bin/car-pooling-challenge" ]
