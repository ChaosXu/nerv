#!/usr/bin/env bash

PID_FILE=agent.pid

function create() {
    echo $(pwd)
    download
    tar -zcvf $PKG $ROOT 
}

function download() {
    wget PKG_URL
    check_pid

}

function check_pid() {
    if [ -f $PID_FILE ];then
        pid=`cat $PID_FILE`
        if [ -n $pid ]; then
            running=`ps -p $pid|grep -v "PID TTY" |wc -l`
            return $running
        fi
    fi
    return 0
}

if [ "$PKG_URL" == "" ]; then
    echo {\"error\":\"PKG_URL is empty\"}
elif [ "$PKG" == "" ]; then 
    echo {\"error\":\"PKG is empty\"}
elif [ "$ROOT" == ""  ]; then 
    echo {\"error\":\"ROOT is empty\"}
else
    cd $ROOT
    create    
fi
