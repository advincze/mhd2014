# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  
  config.vm.box = "mhd2014"

  config.vm.network "forwarded_port", guest: 80, host: 8080

  config.vm.synced_folder "go/src/github.com/advincze/mhd2014", "/home/vagrant/go/src/github.com/advincze/mhd2014"
  config.vm.synced_folder "public", "/home/vagrant/public"

  config.vm.provision "shell", path: "scripts/install.sh"
  
end
