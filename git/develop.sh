#!/usr/bin/env bash
repos=("lmd-client-api2" "lmd-schedule-manager" "lmd-tsp2" "lmd-auth2" "lmd-scheduler2" "lmd-sc" "lmd-am" "lmd-media" "lmd-web-app2")

for d in "${repos[@]}"; do
      echo ""
      echo "-- $d --"
       ( cd "$d" && git co develop )
done