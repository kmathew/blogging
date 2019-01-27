#!/bin/sh

mkdir ./zips

cd ./cmd/authors/
GOOS=linux GOARCH=amd64 go build -o authors
zip -j ./authors.zip authors
mv ./authors.zip ../../zips/authors.zip

cd ..

cd approve
GOOS=linux GOARCH=amd64 go build -o approve
zip -j ./approve.zip approve
mv ./approve.zip ../../zips/approve.zip

cd ..

cd spaces
GOOS=linux GOARCH=amd64 go build -o spaces
zip -j ./spaces.zip spaces
mv ./spaces.zip ../../zips/spaces.zip

cd ..

cd blogs
GOOS=linux GOARCH=amd64 go build -o blogs
zip -j ./blogs.zip blogs
mv ./blogs.zip ../../zips/blogs.zip

