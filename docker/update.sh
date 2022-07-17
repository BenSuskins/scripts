#!/usr/bin/env bash
for file in /Data/*.yml; do
    docker-compose pull -f $file
    docker-compose restart -f $file
done