#!/usr/bin/env bash

startDocker() {
  case $1 in
  1)
    echo "Launching MoDe:Link Locally using refactored Java services"
    echo
    docker-compose -f docker-compose.yml up
    ;;
  2)
    echo "Launching MoDe:Link Locally using Debug Scheduler"
    echo
    docker-compose -f docker-compose-debug.yml up
    ;;
  *)
    echo "$n is an unrecognised option"
    exit 1
    ;;
  esac
}

if [ $# -eq 0 ]; then
  echo "========================================================"
  echo "What version of MoDe:Link would you like to run locally?"
  echo "1. Java"
  echo "2. Scheduler2 Debug"
  echo "========================================================"
  printf "Select Option:"

  read -s -n1 n
  echo
  startDocker $n
else
  startDocker $1
fi
