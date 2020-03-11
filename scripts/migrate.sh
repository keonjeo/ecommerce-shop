#!/usr/bin/env bash

docker run -it -v /home/danko/apps/ecommerce/server/migrations:/migrations --network server_default migrate/migrate -path=/migrations/ -database "postgres://postgres:postgres@database:5432/ecommerce?sslmode=disable" $@
