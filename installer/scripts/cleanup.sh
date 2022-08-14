#!/usr/bin/env bash

source $(dirname $0)/common.sh

main() {
  info "removing venv $1"
  rm -rf /tmp/$1
}

main $@