#!/usr/bin/env bash

function create() {
    if [ ! -d $root ]; then
        mkdir -p $root
    fi
}

if [ "$root" == ""  ]; then
    $root=nerv
fi

create


