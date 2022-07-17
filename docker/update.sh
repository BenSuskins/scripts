#!/usr/bin/env bash
for file in *.yml; do
    echo "Updating: $file"
    docker-compose -f $file pull
    docker-compose -f $file restart
done