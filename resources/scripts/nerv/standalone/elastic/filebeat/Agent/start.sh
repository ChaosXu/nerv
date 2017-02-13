#!/usr/bin/env bash

function start() {
    sudo $APP start -e -c ../config/filebeat.yml -d "publish" || return 1
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
    APP=$APP_ROOT/bin/filebeat

    #echo $(pwd)
    #echo $APP
    start
fi
