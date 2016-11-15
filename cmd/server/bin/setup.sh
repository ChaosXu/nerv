#!/usr/bin/env bash

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

./server -s true


