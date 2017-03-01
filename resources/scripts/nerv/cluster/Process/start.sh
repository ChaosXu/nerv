#!/usr/bin/env bash

function start() {
    $APP start  || return 1
}

if [ "$root" == ""  ]; then
    echo {\"error\":\"root is empty\"}
    exit 1
else
    APP_ROOT=$root$node_name
    APP_PID=$APP_ROOT/log/app.pid
    APP=$APP_ROOT/bin/app

    start
fi
