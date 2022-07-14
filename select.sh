#!/usr/bin/env bash

runScript() {
  case $1 in
  1)
    echo "Pulling Master"
    echo
    ./scripts/git/pull_master.sh
    ;;
  2)
    echo "Pulling Develop"
    echo
    ./scripts/git/pull_develop.sh
    ;;
  3)
    echo "Pulling All"
    echo
    ./scripts/git/pull_all.sh
    ;;
  4)
    echo "Showing Branches"
    echo
    ./scripts/git/branch.sh
    ;;
  5)
    echo "Run Docker"
    echo
    ./scripts/docker/run.sh
    ;;
  6)
    echo "Update Dev Tools"
    echo
    ./scripts/setup/update_all.sh
    ;;
  7)
    echo "QA Release"
    echo
    ./scripts/git/merge_develop.sh
    ;;
  8)
    echo "GCloud Auth Token"
    echo
    ./scripts/gcp/token.sh
    ;;
  *)
    echo "$n is an unrecognised option"
    exit 1
    ;;
  esac
}

if [ $# -eq 0 ]; then
  echo "========================================================"
  echo "1. Pull Master"
  echo "2. Pull Develop"
  echo "3. Pull on current branch"
  echo "4. See Current Branch"
  echo "5. Run Docker"
  echo "6. Update Dev Tools"
  echo "7. Trigger QA Release"
  echo "8. Print GCloud Auth Token"
  echo "========================================================"
  printf "Select Option:"

  read -s -n1 n
  echo
  runScript $n
else
  runScript $1
fi
