#!/usr/bin/env bash

function create() {
    if [ -f $APP_PID ]; then
        $APP stop  || return $?
    fi
    curl -L -O $PKG_FILE
    mkdir $APP_ROOT
    tar -xf $PKG_LOCAL_FILE -C $APP_ROOT --strip-components 1
    if [ $? -ne "0" ]; then
        echo {\"error\":\"tar -xf $PKG_LOCAL_FILE -C $root\"}
    fi

    chmod +x $APP
}

if [ "$file_repository" == "" ]; then
    echo {\"error\":\"file_repository is empty\"}
    exit 1
elif [ "$pkg_url" == "" ]; then
    echo {\"error\":\"pkg_url is empty\"}
    exit 1
elif [ "$root" == ""  ]; then
    echo {\"error\":\"root is empty\"}
    exit 1
else
    APP_ROOT=$root$node_name
    APP_PID=$APP_ROOT/log/app.pid
    APP=$APP_ROOT/bin/app
    PKG_FILE=$file_repository$pkg_url
    PKG_LOCAL_FILE=${PKG_FILE##*/}

    create
fi

