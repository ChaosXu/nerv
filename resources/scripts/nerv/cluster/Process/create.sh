#!/usr/bin/env bash

function create() {
    if [ -f $APP_PID ]; then
        $APP stop  || return $?
    fi
    echo "curl -L -O $PKG_FILE"
    curl -L -O $PKG_FILE
    if [ $? -ne "0" ]; then
        echo {\"error\":\"curl -L -O $PKG_FILE\"}
    fi
    tar -xf $root$PKG_LOCAL_FILE -C $root
    if [ $? -ne "0" ]; then
        echo {\"error\":\"tar -xf $root$PKG_LOCAL_FILE -C $root\"}
    fi

    chmod +x $APP
}

if [ "$pkg_url" == "" ]; then
    echo {\"error\":\"pkg_url is empty\"}
    exit 1
elif [ "$root" == ""  ]; then
    echo {\"error\":\"root is empty\"}
    exit 1
else
    APP_ROOT=$root$node_name
    APP_PID=APP_ROOT/log/app.pid
    APP=$APP_ROOT/bin/app
    PKG_FILE=$pkg_url
    PKG_LOCAL_FILE=${PKG_FILE##*/}

    create
fi

