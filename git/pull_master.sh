#!/usr/bin/env bash
for d in */ ; do
      echo ""
      echo "-- $d --"
      cd "$d"
      hasMasterBranch=`git show-ref refs/heads/master`
      if [ -n "$hasMasterBranch" ]; then
            git checkout master && git pull
      else
            git checkout main && git pull            
      fi
      cd ..
done