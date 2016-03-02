# -*- mode: ruby -*-
# vi: set ft=ruby :

$bootstrap = <<SCRIPT
chown -R vagrant:vagrant /home/vagrant/go
mount bpffs /sys/fs/bpf/ -t bpf
sudo apt-get -y install socat curl
SCRIPT

$build = <<SCRIPT
~/go/src/github.com/noironetworks/cilium-net/common/build.sh
SCRIPT

$install = <<SCRIPT
sudo -E make -C /home/vagrant/go/src/github.com/noironetworks/cilium-net/ install

sudo cp /home/vagrant/go/src/github.com/noironetworks/cilium-net/contrib/cilium-docker.conf /etc/init/
sudo cp /home/vagrant/go/src/github.com/noironetworks/cilium-net/contrib/cilium-net-daemon.conf /etc/init/
sudo cp /home/vagrant/go/src/github.com/noironetworks/cilium-net/contrib/cilium-socket-proxy.conf /etc/init/
sudo service cilium-docker restart
sudo service cilium-net-daemon restart
sudo service cilium-socket-proxy restart
SCRIPT

$testsuite = <<SCRIPT
make -C ~/go/src/github.com/noironetworks/cilium-net/ tests
sudo -E make -C ~/go/src/github.com/noironetworks/cilium-net/ runtime-tests
SCRIPT

Vagrant.configure(2) do |config|
  config.vm.box = "noironetworks/net-next"
  config.vm.synced_folder ".", "/home/vagrant/go/src/github.com/noironetworks/cilium-net", disabled: false

  config.vm.provision "bootstrap", type: "shell", inline: $bootstrap
  config.vm.provision "build", type: "shell", run: "always", privileged: false, inline: $build
  config.vm.provision "install", type: "shell", run: "always", privileged: false, inline: $install
  config.vm.provision "testsuite", type: "shell", privileged: false, inline: $testsuite

  config.vm.provider :libvirt do |libvirt|
    libvirt.memory = 4096
  end

  config.vm.define "node1", primary: true  do |node1|
  end

  config.vm.define "node2", autostart: false do |node2|
  end

end
