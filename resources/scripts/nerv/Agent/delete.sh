#!/usr/bin/env bash

function delete() {
    echo $(pwd)
    if [ -f $APP_PID ]; then
        $APP stop  || return 1
    fi
    if [ -d $APP_ROOT ]; then
        rm -rf $APP_ROOT || return 1
    fi
    if [ -f $PKG_FILE ]; then
        rm -rf $PKG_FILE || return 1
    fi
}

if [ "$PKG_URL" == "" ]; then
    echo {\"error\":\"PKG_URL is empty\"}
elif [ "$ROOT" == ""  ]; then
    echo {\"error\":\"ROOT is empty\"}
else
    PKG=${PKG_URL##*/}
    PKG_FILE=$ROOT/$PKG
    APP_ROOT=$ROOT/${PKG%.*}
    APP_PID=APP_ROOT/log/app.pid
    APP=$APP_ROOT/bin/app
    cd $ROOT
    delete
fi

