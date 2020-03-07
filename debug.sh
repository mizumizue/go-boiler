#!/bin/sh

cp .realize.debug.yaml .realize.yaml

docker-compose stop && docker-compose up -d
