#!/usr/bin/env bash

set -e

echo "=> Preparing to install..."

if [ ! -x "$(command -v git)" ]; then
    echo "error: Git doesn't seem to be installed. Aborting..."
    exit 1
fi

if [[ -z "${GOPATH}" ]]; then
    echo "error: GOPATH isn't set. Aborting..."
    exit 1
fi

install_path="${GOPATH}/src/github.com/globocom/prettylog"

if [ -d "$install_path" ]; then
    echo "=> Updating to latest version..."
    (cd ${install_path} && git pull)
else
    echo "=> Cloning repository into GOPATH..."
    mkdir -p ${install_path}
    git clone git@github.com:globocom/prettylog.git ${install_path}
fi

echo "=> Compiling and installing the application..."
(cd ${install_path} && make install)
