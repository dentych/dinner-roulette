#!/bin/bash

docker run --rm -v $(pwd):/migrate/migrations --network backend_default dinnerdash/migrate -path /migrate/migrations -database postgres://postgres:password@postgres/dinnerdash?sslmode=disable up
