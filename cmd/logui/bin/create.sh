#!/usr/bin/env bash

APP_PID=../log/app.pid
APP=./app
PKG_FILE=../kibana-5.2.0-darwin-x86_64.tar.gz

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

function create() {
    if [ -f $APP_PID ]; then
            $APP stop  || return $?
    fi
    tar -xf $PKG_FILE -C ../
    if [ $? -ne "0" ]; then
        echo {\"error\":\"tar -xf ${PKG_FILE}\"}
    fi
}

create


