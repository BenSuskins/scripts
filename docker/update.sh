#!/usr/bin/env bash
for file in *.yml; do
    docker-compose pull -f $file
    docker-compose restart -f $file
done