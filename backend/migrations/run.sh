#!/bin/bash

USERNAME=$1
PASSWORD=$2
HOST=$3

docker run --rm -v $(pwd):/migrate/migrations dinnerdash/migrate -path /migrate/migrations -database postgres://$USERNAME:$PASSWORD@$HOST/dinnerdash?sslmode=disable up
