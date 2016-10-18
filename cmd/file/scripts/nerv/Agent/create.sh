#!/usr/bin/env bash

PID_FILE=agent.pid

function create() {
    echo $(pwd)
    $APP stop
    wget $PKG_URL
    tar -xf $ROOT/$PKG 
    $APP start
}

if [ "$PKG_URL" == "" ]; then
    echo {\"error\":\"PKG_URL is empty\"}
elif [ "$ROOT" == ""  ]; then 
    echo {\"error\":\"ROOT is empty\"}
else
    PKG=${PKG_URL##*/}
    APP=$ROOT/${PKG%%.*}/bin/app
    cd $ROOT
    create    
fi
