#!/bin/bash

# get host ip for service
LOCAL_HOST=`ip route get 1 | awk '{print $NF;exit}'`

# replace service host
sed -i "s/127.0.0.1/${LOCAL_HOST}/g" etc/prod.yaml

# start service
bin/round_robin_with_weight -c etc/prod.yaml