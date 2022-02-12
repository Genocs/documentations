#!/bin/bash

docker-compose -f ./docker-compose.yml up -d 

sleep 5

docker exec mongo_db1 /scripts/rs-init.sh