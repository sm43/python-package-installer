#!/usr/bin/env bash

info() {
  echo "INFO: $@"
}

main() {
  info "creating venv"
  python3 -m venv /tmp/venv-1

  info "activating venv"
  source /tmp/venv-1/bin/activate

  info "installing a package inside a directory"
  mkdir /tmp/venv-1/target
  pip install install -t /tmp/venv-1/target

  info "deactivating venv"
  deactivate

  info "zipping the package"
  zip -r /tmp/venv-1/temp.zip /tmp/venv-1/target

  info "copying the package to the target location"
  cp /tmp/venv-1/temp.zip temp.zip

  info "removing venv"
  rm -rf /tmp/venv-1
}

main



