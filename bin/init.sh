#!/bin/bash
#
# Copyright 2017,2018 Istio Authors. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# Init script downloads or updates envoy. Called from Makefile, which sets
# the needed environment variables.

ROOT=$(cd $(dirname $0)/..; pwd)
ISTIO_GO=$ROOT

set -o errexit
set -o nounset
set -o pipefail

# Set GOPATH to match the expected layout
GO_TOP=$(cd $(dirname $0)/../../../..; pwd)
OUT=${GO_TOP}/out

export GOPATH=${GOPATH:-$GO_TOP}
# Normally set by Makefile
export ISTIO_BIN=${ISTIO_BIN:-${GO_TOP}/bin}

$ROOT/bin/verify_go_version.sh

# Ensure expected GOPATH setup
if [ ${ROOT} != "${GO_TOP:-$HOME/go}/src/istio.io/istio" ]; then
       echo "Istio not found in GOPATH/src/istio.io/"
       exit 1
fi

# Original circleci - replaced with the version in the dockerfile, as we deprecate bazel
#ISTIO_PROXY_BUCKET=$(sed 's/ = /=/' <<< $( awk '/ISTIO_PROXY_BUCKET =/' WORKSPACE))
#PROXYVERSION=$(sed 's/[^"]*"\([^"]*\)".*/\1/' <<<  $ISTIO_PROXY_BUCKET)
PROXYVERSION=$(grep envoy-debug pilot/docker/Dockerfile.proxy_debug  |cut -d: -f2)
PROXY=debug-$PROXYVERSION

if [ ! -f $ISTIO_BIN/envoy-$PROXYVERSION ] ; then
    mkdir -p $OUT
    pushd $OUT
    # New version of envoy downloaded. Save it to cache, and clean any old version.

    DOWNLOAD_COMMAND=""
    if command -v curl > /dev/null; then
       if curl --version | grep Protocols  | grep https; then
	   DOWNLOAD_COMMAND='curl -Lo -'
       else
           echo curl does not support https, will try wget for downloading files.
       fi
    else
       echo curl is not installed, will try wget for downloading files.
    fi

    if [ -z "${DOWNLOAD_COMMAND}" ]; then
        if command -v wget > /dev/null; then
	    DOWNLOAD_COMMAND='wget -qO -'
        else
            echo wget is not installed.
        fi
    fi

    if [ -z "${DOWNLOAD_COMMAND}" ]; then
        echo Error: curl is not installed or does not support https, wget is not installed. \
             Cannot download envoy. Please install wget or add support of https to curl.
        exit 1
    fi

    ${DOWNLOAD_COMMAND} https://storage.googleapis.com/istio-build/proxy/envoy-$PROXY.tar.gz | tar xz
    rm -f ${ISTIO_BIN}/envoy-*

    mkdir -p $ISTIO_BIN
    cp usr/local/bin/envoy $ISTIO_BIN/envoy-$PROXYVERSION
    popd

    # Make sure the envoy binary exists.
    ln -sf $ISTIO_BIN/envoy-$PROXYVERSION $ISTIO_BIN/envoy

    # Deprecated, may still be used in some tests
    ln -sf $ISTIO_BIN/envoy ${ROOT}/pilot/proxy/envoy
fi
