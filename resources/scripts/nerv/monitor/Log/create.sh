#!/usr/bin/env bash

function addLog() {
    if [ "$config_url" == ""  ]; then
        echo {\"info\":\"config_url is empty\"}
        exit 1
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
    AGENT_URL=http://localhost:3334/
    APP_ROOT=$root$node_name
    APP=$APP_ROOT/bin/app
    PKG_FILE=$file_repository$pkg_url
    PKG_LOCAL_FILE=${PKG_FILE##*/}

    addLog
fi
