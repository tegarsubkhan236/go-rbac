FROM ubuntu:latest
LABEL authors="fauzi"

ENTRYPOINT ["top", "-b"]