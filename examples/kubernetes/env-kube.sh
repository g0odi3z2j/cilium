#!/usr/bin/env bash

if [ -n "${IPV4}" ]; then
    ip=$(ip -4 a show cilium_host scope global | grep inet | awk '{print $2}' | sed 's/\/.*//')
    if [ -z "${ip}" ]; then
        ip=$(grep Host-IPv4 /var/run/cilium/globals/node_config.h | awk '{print $3}')
    fi
    if [ -z "$ip" ]; then
        echo "Unable to derive IPv4 address, please edit env-kube.sh manually"
        return 1
    fi
else
    ip=$(ip -6 a show cilium_host scope global | grep inet6 | awk '{print $2}' | sed 's/\/.*//')
    if [ -z "$ip" ]; then
        ip=$(grep Host-IPv6 /var/run/cilium/globals/node_config.h | awk '{print $3}')
    fi

    if [ -z "$ip" ]; then
        echo "Unable to derive IPv6 address, please edit env-kube.sh manually"
        return 1
    fi
fi

# Edit manually here and assign local Ipv6 address if needed
# ip6=dead::

dns_domain="cilium-test"
## Don't change anything bellow this line ##
## unless you know what you're doing      ##

export API_HOST_IP="${ip}"

if [ -n "${IPV4}" ]; then
    export API_HOST="${API_HOST_IP}"
    export ETCD_HOST="${ip}"
else
    export API_HOST="[${API_HOST_IP}]"
    export ETCD_HOST="[${ip}]"
fi

export DNS_SERVER_IP="${ip}"
export KUBELET_HOST="${ip}"
export DNS_DOMAIN="${dns_domain}"
export ENABLE_DNS=true
export NET_PLUGIN=cni
export API_PORT="8080"
export KUBE_OS_DISTRIBUTION="debian"
export RUNTIME_CONFIG="extensions/v1beta1,extensions/v1beta1/networkpolicies"
export kubectl="/home/vagrant/kubernetes/cluster/kubectl.sh -s ${API_HOST}:${API_PORT}"
export LOG_LEVEL=5
