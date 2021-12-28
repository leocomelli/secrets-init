FROM alpine:latest

RUN apk --no-cache add ca-certificates \
    && update-ca-certificates

COPY dist/secrets-init_linux-amd64 .

ENTRYPOINT ["./secrets-init_linux-amd64"]

