#!/usr/bin/env bash
repos=("lmd-scheduler2" "lmd-client-api2" )

if [ $# -eq 0 ]; then
  echo "Enter Release name:"
  read -n10 n
  echo
  for d in "${repos[@]}"; do
      echo ""
      echo "-- $d --"
       ( cd "$d" && git checkout develop && gh pr create --title "Release $n" --body "Merge Develop into Master for Release $n." --base master --head develop && gh pr create --title "Merge Master $1" --body "Merge Master back Into Devleop for Release $1." --base develop --head master)
  done
else
      for d in "${repos[@]}"; do
            echo ""
            echo "-- $d --"
            ( cd "$d" && git checkout develop && gh pr create --title "Release $1" --body "Merge Develop into Master for Release $1." --base master --head develop && gh pr create --title "Merge Master $1" --body "Merge Master back Into Devleop for Release $1." --base develop --head master)
      done
fi

# Command Example: ./merge_develop.sh "2022-R3-01"