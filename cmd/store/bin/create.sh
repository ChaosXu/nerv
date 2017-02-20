#!/usr/bin/env bash

APP_PID=../log/app.pid
APP=./app
PKG_FILE=../elasticsearch-5.2.1.zip

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

function create() {
    if [ -f $APP_PID ]; then
        $APP stop  || return $?
    fi
    unzip $PKG_FILE -d ../
    if [ $? -ne "0" ]; then
        echo $?
        echo {\"error\":\"unzip $PKG_FILE -d ../\"}
    fi
}

create


