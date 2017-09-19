#!/usr/bin/env bash

set -e

echo -e "\e[33m=> Preparing to install...\e[0m"

if [ ! -x "$(command -v git)" ]; then
    echo -e "\e[31merror: Git doesn't seem to be installed. Aborting..."
    exit 1
fi

if [[ -z "${GOPATH}" ]]; then
    echo -e "\e[31merror: GOPATH isn't set. Aborting..."
    exit 1
fi

install_path="${GOPATH}/src/github.com/globocom/prettylog2"

if [ -d "$install_path" ]; then
    echo -e "\e[33m=> Updating to latest version...\e[0m"
    (cd ${install_path} && git pull)
else
    echo -e "\e[33m=> Cloning repository into GOPATH...\e[0m"
    mkdir -p ${install_path}
    git clone git@github.com:globocom/prettylog.git ${install_path}
fi

echo -e "\e[33m=> Compiling and installing the application...\e[0m"
(cd ${install_path} && make install)