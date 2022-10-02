#!/usr/bin/env bash
for d in */ ; do
    cd $d
    for file in *.yml; do
        echo ""
        echo "-- $file --"
        docker compose -f $file pull
        docker compose -f $file restart
    done
    cd ..
done