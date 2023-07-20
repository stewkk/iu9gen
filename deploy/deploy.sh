#!/usr/bin/env sh

git -C iu9gen/ pull \
    && docker compose -f iu9gen/deploy/docker-compose.yaml pull iu9gen \
    && docker compose -f iu9gen/deploy/docker-compose.yaml up --no-deps -d iu9gen
