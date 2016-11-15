#!/usr/bin/env bash

echo "setup all 3td dependencies..."

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

../store/bin/setup.sh
../server/bin/setup.sh
