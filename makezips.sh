#!/bin/sh
cd ./cmd/authors/
GOOS=linux GOARCH=amd64 go build -o authors
zip -j ./authors.zip authors

cd ..

cd approve
GOOS=linux GOARCH=amd64 go build -o approve
zip -j ./approve.zip approve

cd ..

cd spaces
GOOS=linux GOARCH=amd64 go build -o spaces
zip -j ./spaces.zip spaces

cd ..

cd blogs
GOOS=linux GOARCH=amd64 go build -o blogs
zip -j ./blogs.zip blogs

