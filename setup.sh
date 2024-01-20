#!/bin/bash
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
echo "export PATH=$PATH:$(go env GOPATH)/bin"

cd backend
openssl genrsa -out server.key 2048
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

cd ../
openssl genrsa -out client.key 2048
openssl req -new -x509 -sha256 -key client.key -out client.crt -days 3650