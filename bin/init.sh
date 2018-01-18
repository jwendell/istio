#!/bin/bash

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
    curl -Lo - https://storage.googleapis.com/istio-build/proxy/envoy-$PROXY.tar.gz | tar xz
    rm -f ${ISTIO_BIN}/envoy-*

    mkdir -p $ISTIO_BIN
    cp usr/local/bin/envoy $ISTIO_BIN/envoy-$PROXYVERSION
    popd

    # Make sure the envoy binary exists.
    ln -sf $ISTIO_BIN/envoy-$PROXYVERSION $ISTIO_BIN/envoy

    # Deprecated, may still be used in some tests
    ln -sf $ISTIO_BIN/envoy ${ROOT}/pilot/proxy/envoy
fi
