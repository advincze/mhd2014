#!/bin/bash

pkill -f httpster

nohup httpster -p 3333 -d /home/vagrant/public 2>&1> /dev/null &
