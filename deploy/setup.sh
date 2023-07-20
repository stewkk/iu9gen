#!/usr/bin/env sh

git clone --branch master --single-branch https://github.com/stewkk/iu9gen.git \
    && cp iu9gen/deploy/deploy.sh deploy.sh \
    && sudo mkdir -p /var/swag/nginx/proxy-confs \
    && sudo chown -R $USER /var/swag/nginx \
    && cp iu9gen/deploy/iu9gen.subdomain.conf /var/swag/nginx/proxy-confs/ \
    && docker compose -f iu9gen/deploy/docker-compose.yaml up -d
