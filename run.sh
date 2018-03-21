#!/bin/bash

docker build -t payments-server .
docker run --rm -t -i payments-server
