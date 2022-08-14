#!/usr/bin/env bash

source $(dirname $0)/common.sh

# $1 => venv directory
# $2 => pkg directory
# $3 => package name
main() {
  info "activating venv $1"
  source /tmp/$1/bin/activate

  info "installing $3 package in $2 directory in $1 venv"
  mkdir /tmp/$1/$2
  pip install $3 -t /tmp/$1/$2

  info "deactivating venv $1"
  deactivate
}

main $@



