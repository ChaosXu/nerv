#!/usr/bin/env bash

function stop() {
    $APP stop  || return 1
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

    stop
fi
