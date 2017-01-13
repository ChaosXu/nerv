#!/usr/bin/env bash

function create() {
    if [ -f $APP_PID ]; then
        $APP stop  || return $?
    fi
    echo $PKG_FILE
    tar -xf $PKG_FILE -C $root
    if [ $? -ne "0" ]; then
        echo {\"error\":\"tar -xf ${PKG_FILE}\"}
    fi
    #$APP start
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
    #cd $ROOT

    APP_ROOT=$root/${PKG%%.*}
    APP_PID=APP_ROOT/log/app.pid
    APP=$APP_ROOT/bin/app
    PKG_FILE=${pkg_root}${pkg_url}

    create
fi

