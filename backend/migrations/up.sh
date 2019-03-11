#!/bin/bash

migrate -source file://. -database postgres://postgres:password@localhost:5432/dinner-dash?sslmode=disable up
