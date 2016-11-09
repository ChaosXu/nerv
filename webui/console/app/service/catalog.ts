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

const Catalogs = [
    {
        name: '计算',
        items: [{
            name: '云主机',
            icon: 'icon__uhost',
            url: '/uhost'
        }]
    },
    {
        name: '管理',
        items: [{
            name: '账号与权限',
            icon: 'icon__company',
            url: '/company'
        }]
    }
];

@Injectable()
export class CatalogService {    
    getCatalogs(): Promise<Catalog[]> {
        return Promise.resolve(Catalogs);
    }
}