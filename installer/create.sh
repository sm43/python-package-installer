#!/usr/bin/env bash

source $(dirname $0)/common.sh

main() {
  info "creating venv $1"
  python3 -m venv /tmp/$1
}

main $@