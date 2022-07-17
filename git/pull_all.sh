#!/usr/bin/env bash
for d in */ ; do
      echo ""
      echo "-- $d --"
      echo -n "Branch: "
       ( cd "$d" && git branch --show-current && git pull)
done