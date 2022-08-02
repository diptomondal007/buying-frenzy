#! /bin/bash

export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=test
export DB_PASSWORD=password
export DB_NAME=food

echo "docker containers starting ..."
docker-compose up postgres -d

echo "building application ...."
go build -v .
./buying-frenzy etl -d data