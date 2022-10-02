#!/usr/bin/env bash
for d in */ ; do
    for file in *.yml; do
        echo ""
        echo "-- $file --"
        docker compose -f $file pull
        docker compose -f $file restart
    done
done