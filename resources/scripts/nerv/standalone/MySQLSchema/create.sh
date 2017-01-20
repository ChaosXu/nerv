#!/usr/bin/env bash

function create() {
    create_db_sql="create database IF NOT EXISTS ${schema} DEFAULT CHARSET utf8"
    mysql -h${host}  -P${port}  -u${user} -p${password} -e "${create_db_sql}"
    return $?
}

if [ "$host" == "" ]; then
    echo {\"error\":\"host is empty\"}
    exit 1
elif [ "$port" == "" ]; then
    echo {\"error\":\"port is empty\"}
    exit 1
elif [ "$user" == ""  ]; then
    echo {\"error\":\"user is empty\"}
    exit 1
elif [ "$password" == ""  ]; then
    echo {\"error\":\"password is empty\"}
    exit 1
elif [ "$schema" == ""  ]; then
    echo {\"error\":\"schema is empty\"}
    exit 1
else
    create
fi

