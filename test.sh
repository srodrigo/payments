#!/bin/bash

docker build -f Dockerfile-tests -t payments-server-tests .
docker run --rm -t -i payments-server-tests
