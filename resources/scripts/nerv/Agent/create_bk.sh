#!/usr/bin/env bash

function create() {
    echo $(pwd)
    if [ -f $APP_PID ]; then
        $APP stop  || return 1
    fi
    if [ -f $PKG_FILE ]; then
        rm -rf $PKG_FILE
    fi
    tar -xf $PKG_FILE || return 1
    $APP start
}

if [ "$pkg_root" == "" ]; then
    echo {\"error\":\"pkg_root is empty\"}
    exit 1
elif [ "$pkg_url" == "" ]; then
    echo {\"error\":\"pkg_url is empty\"}
    exit 1
elif [ "$root" == ""  ]; then
    echo {\"error\":\"root is empty\"}
    exit 1
else
    #PKG=${PKG_URL##*/}
    PKG_FILE=${pkg_root}${pkg_url}
    APP_ROOT=$ROOT/${PKG%.*}
    APP_PID=APP_ROOT/log/app.pid
    APP=$APP_ROOT/bin/app
    cd $ROOT
    create
fi

