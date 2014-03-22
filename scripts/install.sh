#!/bin/bash

cp "/vagrant/scripts/nginx.conf" "/etc/nginx/nginx.conf"

# Make actual go workspace dir structure
chown -R "vagrant" "/home/vagrant/go"


# Set GOPATH
export GOPATH="/home/vagrant/go"
echo 'export GOPATH="/home/vagrant/go"' | tee -a /etc/profile

# Adds go bin directory to path so tools
# and buils are available on the commandline
export PATH="$PATH:$GOPATH/bin"
echo 'export PATH="$PATH:$GOPATH/bin"' | sudo tee -a /etc/profile


