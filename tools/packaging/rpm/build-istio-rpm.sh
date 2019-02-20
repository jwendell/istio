#!/bin/bash

#set -x

# We rely on the following ENV variables:
# ISTIO_GO
# ISTIO_OUT

PKG_DIR="${ISTIO_GO}/tools/packaging"
WORK_DIR="$(mktemp -d)"

cp -r "${PKG_DIR}" "${WORK_DIR}"
mv ${WORK_DIR}/packaging/common/* ${WORK_DIR}/packaging/rpm/istio

cd "${ISTIO_GO}/.."
tar cfz "${WORK_DIR}/packaging/rpm/istio/istio.tar.gz" --exclude=.git istio

cd "${WORK_DIR}/packaging/rpm/istio"
fedpkg --name istio --release el7 local

mkdir -p "${ISTIO_OUT}/rpm"
cp -r x86_64/* "${ISTIO_OUT}/rpm"
