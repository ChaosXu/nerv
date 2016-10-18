#!/usr/bin/env bash
function create() {
    wget PKG_URL
}

function download() {
    echo "download"
}

if [ "$PKG_URL" == "" ]; then
    echo {\"error\":\"PKG_URL is empty\"}
else
    create    
fi
