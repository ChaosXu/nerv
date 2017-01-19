# Nerv  [![Build Status](https://travis-ci.org/ChaosXu/nerv.svg?branch=master)](https://travis-ci.org/ChaosXu/nerv)

## 概述

神经元为物理机、私有云、公有云、容器及混合云环境提供PaaS服务，支持应用和服务的部署、运维。

## 从源码构建

### 准备

[安装Go](https://golang.org/doc/install)

### 构建

```shell
go get github.com/ChaosXu/nerv
cd $GOPATH/src/github.com/ChaosXu/nerv
nerv$ make all -e ENV=debug
nerv$ cd release
release$ ls
nerv            nerv.tar.gz
```

## 快速启动（单机版）

### 配置数据库
创建一个MySQL数据库 nerv
打开 release/nerv/nerv-cli/config/config.json，配置数据库连接

```shell
{
  "db": {
    "user": "root",
    "password": "root",
    "url": "/nerv?charset=utf8&parseTime=True&loc=Local"
  },
  "agent": {
    "port": "3334"
  }
}
```

### 安装与启动

```shell
cd release/nerv/nerv-cli/bin
bin$ ./nerv-cli topo create -t ../../resources/templates/nerv/env_standalone.json -o nerv-test
Create topology success. id=1
bin$ ./nerv-cli topo list
ID      Name    RunStatus       CreateAt        Template
1       nerv-test       0       XXXXX           ../../resources/templates/nerv/env_standalone_test.json
bin$ ./nerv-cli topo install -i 1
Install topology success. id=1
bin$ ./nerv-cli topo setup -i 1
Setup topology success. id=1
bin$ ./nerv-cli topo start -i 1
file: started, pid=30992
agent: started, pid=30988
server: started, pid=31038
webui: started, pid=33065
Start topology success. id=1
```

### 访问WEB控制台

### 停止

```shell
bin$ ./nerv-cli topo start -i 1
Stop topology success. id=1
```

## 工作机制

TBD

## 部署与配置

TBD

