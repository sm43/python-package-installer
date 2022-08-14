#!/usr/bin/env bash

source $(dirname $0)/common.sh

# $1 => venv directory
# $2 => pkg directory
# $3 => target location
main() {
   zipFile=$2.zip

   info "zipping package $3 in $1 venv"
   zip -r /tmp/$1/${zipFile} /tmp/$1/$2

   info "copying /tmp/$1/${zipFile} package to $3/${zipFile}"
   cp /tmp/$1/${zipFile} $3/${zipFile}
}

main $@