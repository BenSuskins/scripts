#!/usr/bin/env bash

runScript() {
  case $1 in
  1)
    echo "Pulling All"
    echo
    $HOME/scripts/git/pull_all.sh
    ;;
  2)
    echo "Updating Docker Compose"
    echo
    $HOME/scripts/docker/update.sh
    ;;
  3)
    echo "Updating System"
    echo
    $HOME/scripts/setup/update.sh
    ;;
  4)
    echo "Cleanup"
    echo
    $HOME/scripts/setup/cleanup.sh
    ;;
  *)
    echo "$n is an unrecognised option"
    exit 1
    ;;
  esac
}

if [ $# -eq 0 ]; then
  echo "========================================================"
  echo "1. Pull on current branch"
  echo "2. Update Docker Compose"
  echo "3. Update System"
  echo "4. Cleanup"
  echo "========================================================"
  printf "Select Option:"

  read -s -n1 n
  echo
  runScript $n
else
  runScript $1
fi
