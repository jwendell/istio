#!/bin/bash

#set -x

# We rely on the following ENV variables:
# ISTIO_ENVOY_VERSION
# ISTIO_OUT
# ISTIO_GO

PKG_DIR="${ISTIO_GO}/tools/packaging"
WORK_DIR="$(mktemp -d)"
RPM_DIR="${WORK_DIR}/packaging/rpm/proxy"

cp -r "${PKG_DIR}" "${WORK_DIR}"
mv ${WORK_DIR}/packaging/common/* ${RPM_DIR}

source /opt/rh/devtoolset-6/enable

cd /builder
git clone  https://github.com/istio/proxy.git istio-proxy
cd istio-proxy
git checkout ${ISTIO_ENVOY_VERSION}

cat "${RPM_DIR}/bazelrc" >> .bazelrc
BUILD_SCM_REVISION=$(date +%s)
sed -i "1i echo BUILD_SCM_REVISION ${BUILD_SCM_REVISION}\necho BUILD_SCM_STATUS Clean\nexit 0" tools/bazel_get_workspace_status

cd ..
tar cfz "${RPM_DIR}/istio-proxy.tar.gz" --exclude=.git istio-proxy
cd "${RPM_DIR}"

fedpkg --release el7 local

mkdir -p "${ISTIO_OUT}/rpm"
cp -r x86_64/* "${ISTIO_OUT}/rpm"
