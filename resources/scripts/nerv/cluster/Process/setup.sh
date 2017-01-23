#!/usr/bin/env bash

function setup() {
    if [ "$config_root" == ""  ]; then
        echo {\"info\":\"config_root is empty\"}
        exit 0
    elif [ "$config_url" == ""  ]; then
        echo {\"info\":\"config_url is empty\"}
        exit 0
    fi

    cd ${APP_ROOT}/config
    curl -L -O $config_root/$config_url
    if [ $? -ne "0" ]; then
        echo {\"error\":\"curl -L -O $config_root/$config_url\"}
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
    APP_PID=$APP_ROOT/log/app.pid
    APP=$APP_ROOT/bin/app
    PKG_FILE=$PKG

    setup
fi

