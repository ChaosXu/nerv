import { Injectable } from '@angular/core';

export interface Catalog {
    name: string;
    items: CatalogItem[];
}

export interface CatalogItem {
    name: string;
    icon: string;
    url: string;
    dock: boolean;
}

const Catalogs = [{
    name: '管理',
    items: [{
        name: '服务编排',
        icon: 'icon__ucdn',
        url: '/orchestration/Topology'
    }, {
        name: '部署流水线',
        icon: 'icon__umon',
        url: '/pipeline/Project'
    }, {
        name: '监控',
        icon: 'icon__umon',
        url: '/monitor'
    }]
}, {
    name: '中间件',
    items: [{
        name: '分布式数据',
        icon: 'icon__uphost',
        url: '/infrastructure/Host'
    }, {
        name: '消息队列',
        icon: 'icon__uphost',
        url: '/infrastructure/Host'
    }, {
        name: '数据传输',
        icon: 'icon__uphost',
        url: '/infrastructure/Host'
    }, {
        name: 'MySQL HA',
        icon: 'icon__uphost',
        url: '/infrastructure/Host'
    }, {
        name: '分布式服务框架',
        icon: 'icon__uphost',
        url: '/infrastructure/Host'
    }, {
        name: 'Nginx',
        icon: 'icon__uphost',
        url: '/infrastructure/Host'
    }]
}, {
    name: '数据库',
    items: [{
        name: 'MySQL',
        icon: 'icon__company',
        url: '/user/Account'
    }, {
        name: 'Redis',
        icon: 'icon__company',
        url: '/user/Account'
    }]
}, {
    name: '存储',
    items: [{
        name: '对象存储',
        icon: 'icon__company',
        url: '/user/Account'
    }]
}, {
    name: '系统',
    items: [{
        name: '用户',
        icon: 'icon__company',
        url: '/user/Account'
    }]
}];

@Injectable()
export class CatalogConfig {
    getCatalogs(): Promise<Catalog[]> {
        return Promise.resolve(Catalogs);
    }
}