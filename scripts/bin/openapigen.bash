#!/usr/bin/env bash
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    -i $1 \ #https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/2_0/petstore.yaml \
    -g go \
    -o /local/out/go