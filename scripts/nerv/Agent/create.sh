#!/usr/bin/env bash

PID_FILE=agent.pid

function create() {
    echo $(pwd)
    if [ -f $APP ]; then
        $APP stop  || return 1
    fi
    if [ -f $PKG_FILE ]; then
        rm -rf $PKG_FILE
    fi
    curl -Os $PKG_URL || return 1
    tar -xf $PKG_FILE || return 1
    $APP start
}

if [ "$PKG_URL" == "" ]; then
    echo {\"error\":\"PKG_URL is empty\"}
elif [ "$ROOT" == ""  ]; then 
    echo {\"error\":\"ROOT is empty\"}
else
    PKG=${PKG_URL##*/}
    PKG_FILE=$ROOT/$PKG
    APP=$ROOT/${PKG%%.*}/bin/app
    cd $ROOT
    create    
fi
