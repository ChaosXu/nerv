#!/usr/bin/env bash

function create() {
    if [ -f $APP_PID ]; then
        $APP stop  || return $?
    fi
    tar -xf $PKG_FILE -C $root
    if [ $? -ne "0" ]; then
        echo {\"error\":\"tar -xf ${PKG_FILE}\"}
    fi
    if [ -f $APP_ROOT/bin/create.sh ]; then
        $APP_ROOT/bin/create.sh || return $?
    fi
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
    PKG=${pkg_url##*/}
    APP_ROOT=$root${PKG%.*}
    APP_PID=APP_ROOT/log/app.pid
    APP=$APP_ROOT/bin/app
    PKG_FILE=${pkg_root}${pkg_url}

    create
fi

