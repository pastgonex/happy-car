#! /bin/bash
function Build {
    DOMAIN=$1
    cd ../server
    docker build -t happycar/$DOMAIN -f ../deployment/$DOMAIN/Dockerfile .
}

Build auth
Build rental
Build gateway
build blob
build car
