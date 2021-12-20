FROM scratch

COPY dist/secrets-init_linux-amd64 .

ENTRYPOINT ["./secrets-init_linux-amd64"]

