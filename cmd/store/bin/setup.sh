#!/usr/bin/env bash

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

curl -L -O https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-5.0.0.tar.gz

tar -xvf elasticsearch-5.0.0.tar.gz -C ..


