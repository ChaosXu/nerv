# Nerv  [![Build Status](https://travis-ci.org/ChaosXu/nerv.svg?branch=master)](https://travis-ci.org/ChaosXu/nerv)

## 概述

神经元为物理机、私有云、公有云、容器及混合云环境提供PaaS服务，支持应用和服务的部署、运维。

![user_flow](/docs/img/use_flow.png)

## 从源码构建

### 准备

#### Go

* [安装Go](https://golang.org/doc/install)

#### WebUI

* [安装Node.js](https://nodejs.org/zh-cn)
* [安装gulp](http://gulpjs.com)
* [Angular 2 参考](https://www.angular.cn/docs/ts/latest)

### 构建

```shell
go get github.com/ChaosXu/nerv

cd $GOPATH/src/github.com/ChaosXu/nerv/cmd/webui/console
console$ npm install

cd $GOPATH/src/github.com/ChaosXu/nerv
nerv$ make build webui pkg -e ENV=debug

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

### 安装与配置

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
```

### 启动

```shell
bin$ ./nerv-cli topo start -i 1
file: started, pid=30992
agent: started, pid=30988
server: started, pid=31038
webui: started, pid=33065
Start topology success. id=1
```

### 停止

```shell
bin$ ./nerv-cli topo stop -i 1
Stop topology success. id=1
```

### 卸载

```shell
bin$ ./nerv-cli topo uninstall -i 1
Uninstall topology success. id=1
```

### 删除

```shell
bin$ ./nerv-cli topo delete -i 1
Delete topology success. id=1
```

### 访问WebUI

浏览器打开：http://localhost:3330
缺省口令：admin
![webui_login](/docs/img/webui_login.png)
![topology_list](/docs/img/topology_list.png)

## 工作机制

![concept](/docs/img/concept.png)

### Application
供人直接使用的程序。

### Service
供其它应用或服务使用的后台程序。

### Template
定义构成一个应用或服务所需的资源（服务器、安装包、配置文件等）及它们之间关系。

### Topology
使用模板创建的一个应用或服务的拓扑结构，通过Install、Setup、Start等操作部署、配置和启动应用或服务。

### Resource Model
定模板中的元素的类型，适配各种部署环境：物理机、公有云、私有云、容器以及混合云等。

### Work Cluster
部署应用或服务的集群。

### Agent
运行与集群中的每台主机上的代理程序，负责执行本机上部署的应用或服务的实例的管理和监控工作。

## 部署与配置

TBD

