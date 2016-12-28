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
        name: '资源编排',
        icon: 'icon__ucdn',
        url: '/orchestration/Topology'
    }, {
        name: '监控',
        icon: 'icon__umon',
        url: '/monitor'
    }]
}, {
    name: '基础设施',
    items: [{
        name: '主机管理',
        icon:'icon__uphost',
        url: '/infrastructure/Host'
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