#! /bin/bash
DOMAIN=$1
cd ../server
docker build -t happycar/$DOMAIN -f ../deployment/$DOMAIN/Dockerfile .