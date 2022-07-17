#!/usr/bin/env bash
for d in */ ; do
      echo ""
      echo "-- $d --"
       ( cd "$d" && git co master )
done