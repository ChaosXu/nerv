import { Injectable } from '@angular/core';

export interface Catalog {
    name: string;
    items: CatalogItem[];
}

export interface CatalogItem {
    name: string;
    icon: string;
    url: string;
}

const Catalogs = [
    {
        name: '计算',
        items: [{
            name: '云主机',
            icon: 'icon__uhost',
            url: '/uhost'
        }]
    }
];

@Injectable()
export class CatalogService {
    getCatalogs(): Promise<Catalog[]> {
        return Promise.resolve(Catalogs);
    }
}