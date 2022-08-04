#! /bin/bash

export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=956cfb292fa9189332dc28d3237aa3ff63bb7c89bbec27ae
export DB_NAME=buying-frenzy

echo "building application ...."
go build -v .
./buying-frenzy etl -d data