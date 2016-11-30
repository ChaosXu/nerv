import { Component } from '@angular/core';
import { EventEmitter } from '@angular/core';
import { Catalog,CatalogItem, CatalogConfig } from '../lib/config/catalog.config';

@Component({
    selector: 'nerv-startmenu',
    templateUrl: 'app/main/startmenu.html',
    outputs: ['eventDockCatalogItem:dock']
})
export class StartMenu {
    title = '产品与服务';
    catalogs: Catalog[];
    catalogSvc: CatalogConfig;
    display: boolean;
    eventDockCatalogItem = new EventEmitter<CatalogItem>();

    constructor(catalogSvc: CatalogConfig) {
        this.catalogSvc = catalogSvc;
    }

    showMenu(): void {
        this.display = !this.display;
        this.catalogSvc.getCatalogs().then(catalogs => this.catalogs = catalogs);
    }

    onDockCatalogItem(item: CatalogItem): void {
        item.dock = !item.dock;
        this.eventDockCatalogItem.emit(item);
    }
}