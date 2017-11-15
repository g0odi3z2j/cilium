#!/usr/bin/env bash
set -e
dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
root_dir="${dir}/../.."
dest="$1"

cp -R "${root_dir}/.git" "${dest}"
cp -R "${root_dir}/api" "${dest}"
cp -R "${root_dir}/bpf" "${dest}"
cp -R "${root_dir}/cilium" "${dest}"
cp -R "${root_dir}/cilium-health" "${dest}"
cp -R "${root_dir}/common" "${dest}"
mkdir -p "${dest}/contrib/packaging/docker"
cp "${root_dir}/contrib/packaging/docker/clang-3.8.1.key" "${dest}/contrib/packaging/docker"
cp -R "${root_dir}/contrib/scripts" "${dest}/contrib"
cp -R "${root_dir}/contrib/systemd" "${dest}/contrib"
cp -R "${root_dir}/contrib/upstart" "${dest}/contrib"
cp -R "${root_dir}/daemon" "${dest}"
cp -R "${root_dir}/envoy" "${dest}"
cp -R "${root_dir}/pkg" "${dest}"
cp -R "${root_dir}/plugins" "${dest}"
cp -R "${root_dir}/vendor" "${dest}"
cp -R "${root_dir}/Makefile" "${dest}"
cp -R "${root_dir}/Makefile.defs" "${dest}"
cp -R "${root_dir}/LICENSE" "${dest}"
cp -R "${root_dir}/AUTHORS" "${dest}"
cp -R "${root_dir}/VERSION" "${dest}"
cp -R "${root_dir}/monitor" "${dest}"
