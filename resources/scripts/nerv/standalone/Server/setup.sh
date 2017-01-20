#!/usr/bin/env bash

function setup() {
    $SETUP || return 1
}


function config() {
    if [ "$config_root" == ""  ]; then
        echo {\"info\":\"config_root is empty\"}
        exit 0
    elif [ "$config_url" == ""  ]; then
        echo {\"info\":\"config_url is empty\"}
        exit 0
    fi

    cp -r $config_root$config_url/ $APP_ROOT/config
    if [ $? -ne "0" ]; then
        echo {\"error\":\"cp -r $config_root$config_url/ $APP_ROOT/config\"}
        return 1
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
    APP_ROOT=$root${PKG%%.*}
    APP_PID=APP_ROOT/log/app.pid
    APP=$APP_ROOT/bin/app
    SETUP=$APP_ROOT/bin/setup.sh

    config || return 1
    setup
fi
