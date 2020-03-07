#!/bin/sh

cp .realize.normal.yaml .realize.yaml

docker-compose stop && docker-compose up -d
