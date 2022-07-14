#!/usr/bin/env bash
repos=("lmd-client-api2" "lmd-tsp2" "lmd-auth2" "lmd-scheduler2" "lmd-sc" "lmd-am" "lmd-media" "lmd-web-app2" "lmd-ios-app" "lmd-jenkins-files" "lmd-schedule-manager")

for d in "${repos[@]}"; do
      echo ""
      echo "-- $d --"
      echo -n "Branch: "
       ( cd "$d" && git branch --show-current && git pull)
done