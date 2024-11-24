#!/bin/bash

go generate ./...
go run .gqlgen/bson.go

echo "Generated gql models with BSON tags"
