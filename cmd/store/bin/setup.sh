#!/usr/bin/env bash

ES=elasticsearch-5.0.0
ES_PKG=${ES}.tar.gz

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

if ![ -f ${ES_PKG} ]; then
    curl -L -O https://artifacts.elastic.co/downloads/elasticsearch/${ES_PKG}
fi

tar -xvf ${ES_PKG} -C ..


