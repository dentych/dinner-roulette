#!/bin/bash

# Need migrate binary: https://github.com/golang-migrate/migrate/releases
migrate -path ./migrations -database postgres://postgres:password@localhost:5432/dinner-dash?sslmode=disable up
