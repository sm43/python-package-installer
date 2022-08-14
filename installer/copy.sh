#!/usr/bin/env bash

source $(dirname $0)/common.sh

# $1 => venv directory
# $2 => target directory
# $3 => zip name
main() {
   info "zipping package $3 in $1 venv"
   zip -r /tmp/$1/$3 /tmp/$1/$2

  info "copying $3 package to the target location"
  cp /tmp/$1/$3 $3
}

main $@