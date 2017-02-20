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

    tar -xf $PKG_LOCAL_FILE -C $root
    if [ $? -ne "0" ]; then
        echo {\"error\":\"tar -xf $PKG_LOCAL_FILE -C $root\"}
    fi
    chmod +x $APP

    local os=$(uname)
    local lib=filebeat-5.2.0-$os-x86_64.tar.gz
    curl -L -O $pkg_repository/$lib
    if [ $? -ne "0" ]; then
        echo {\"error\":\"curl -L -O $pkg_root/$lib\"}
    fi
    tar -xf $lib -C $APP_ROOT
    if [ $? -ne "0" ]; then
        echo {\"error\":\"tar -xf $APP_ROOT/$lib -C $APP_ROOT\"}
        return 1
    fi
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
    APP_PID=APP_ROOT/log/app.pid
    APP=$APP_ROOT/bin/app
    PKG_FILE=$file_repository$pkg_url
    PKG_LOCAL_FILE=${PKG_FILE##*/}

    create
fi

