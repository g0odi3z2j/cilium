#!/bin/bash

# Copyright 2017-2021 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

set -o errexit
set -o pipefail
set -o nounset

root_dir="$(git rev-parse --show-toplevel)"
hubble_version_script="${root_dir}/images/cilium/hubble-version.sh"
hubble_cli_example="${root_dir}/examples/hubble/hubble-cli.yaml"

cd "${root_dir}"

if [ "$#" -ne 1 ] ; then
  echo "$0 supports exactly 1 argument - <hubble_version>"
  exit 1
fi

# remove the leading v, if any, allowing both vX.Y.Z and X.Y.Z as version
# parameter.
hubble_version="$(printf "%s" "$1" | sed 's/^v\(.*\)/\1/')"

# this is a simple array that assumes the order of the loop;
# it's not an associative array because this script needs to
# work on any version of bash, and (most notably) macOS ships
# an old version that doesn't support associative arrays
hubble_sha256=()

for arch in amd64 arm64 ; do
  tmpout="$(mktemp)"
  curl --fail --show-error --silent --location \
    "https://github.com/cilium/hubble/releases/download/v${hubble_version}/hubble-linux-${arch}.tar.gz.sha256sum" \
    --output "${tmpout}"
  read -ra sha256 < "${tmpout}"
  rm -f "${tmpout}"
  hubble_sha256+=("${sha256[0]}")
done

cat > "$hubble_version_script" << EOF
# Code generated by images/scripts/update-hubble-version.sh; DO NOT EDIT.
hubble_version="v${hubble_version}"
declare -A hubble_sha256
hubble_sha256[amd64]="${hubble_sha256[0]}"
hubble_sha256[arm64]="${hubble_sha256[1]}"
EOF

sed -e "s#\\(image: .*/hubble:\\).*#\1v${hubble_version}#" "$hubble_cli_example" > "${hubble_cli_example}.tmp" &&
    mv "${hubble_cli_example}.tmp" "$hubble_cli_example"
