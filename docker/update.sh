#!/usr/bin/env bash
for d in */ ; do
    cd $d
    for file in *.yml; do
        echo ""
        echo "-- $file --"
        docker compose -f $file pull
        docker compose -f $file stop
        docker compose -f $file start
    done
    cd ..
done
