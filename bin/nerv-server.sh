#!/usr/bin/env bash

WORKSPACE=$(cd $(dirname $0)/; pwd)
cd $WORKSPACE

function start() {
    ../server/bin/app start
    ../file/bin/app start
}

function stop() {
    ../server/bin/app stop
    ../file/bin/app stop
}

function restart() {
    stop
    sleep 3
    start
}

function status() {
    ../server/bin/app status
    ../file/bin/app status
}


function help() {
    echo "start|stop|restart|status"
}

if [ "$1" == "" ]; then
    help
elif [ "$1" == "stop" ];then
    stop
elif [ "$1" == "start" ];then
    start
elif [ "$1" == "restart" ];then
    restart
elif [ "$1" == "status" ];then
    status
else
    help
fi
