#!/usr/bin/env bash

source "./helpers.bash"

NETPERF_IMAGE="tgraf/netperf"

function cleanup {
	cilium policy delete --all 2> /dev/null || true
	docker rm -f server1 server2 2> /dev/null || true
	docker network rm $TEST_NET > /dev/null 2>&1
}

cleanup

trap cleanup EXIT

create_cilium_docker_network

set -x

sudo service cilium restart
wait_for_cilium_status

docker run -dt --net=$TEST_NET --name server1 -l id.server1 $NETPERF_IMAGE

docker run -dt --net=$TEST_NET --name server2 -l id.server2 $NETPERF_IMAGE

wait_for_cilium_ep_gen

before_restart=$(cilium endpoint list)
before_restart_md5=$(echo "${before_restart}" | md5sum)

sudo service cilium restart

wait_for_cilium_status

wait_for_cilium_ep_gen

after_restart=$(cilium endpoint list)
after_restart_md5=$(echo "${after_restart}" | md5sum)

if [[ "${before_restart_md5}" != "${after_restart_md5}" ]]; then
    echo "Wanted:"
    echo "${before_restart}"
    echo "Received:"
    echo "${after_restart}"
    abort "Restore functionality didn't work!"
fi

set +x
