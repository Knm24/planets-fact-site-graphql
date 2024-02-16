FROM debian:latest

# Add Maintainer Info
LABEL maintainer="Talgat Saribayev <talgat.s@protonmail.com>"

WORKDIR /app

RUN apt-get update \
    && apt-get install -y curl

RUN curl -L https://github.com/hasura/graphql-engine/raw/stable/cli/get.sh | /bin/bash
