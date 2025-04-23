FROM ubuntu:latest
LABEL authors="amerk"

ENTRYPOINT ["top", "-b"]