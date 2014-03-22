#!/bin/bash

pkill -f httpster

httpster -p 3333 -d /home/vagrant/public 2>&1> /dev/null &
