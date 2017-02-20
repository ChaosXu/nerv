#!/usr/bin/env bash

function setup() {
    if [ "$config_url" == ""  ]; then
        echo {\"info\":\"config_url is empty\"}
        exit 0
    fi

    if [ ! -d $APP_ROOT/config ]; then
        mkdir $APP_ROOT/config
    fi
    cd $APP_ROOT/config
    curl -L -O $file_repository$config_url
    if [ $? -ne "0" ]; then
        echo {\"error\":\"curl -L -O $file_repository$config_url $APP_ROOT\"}
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

    setup
fi
