import { Component } from '@angular/core';
import { EventEmitter } from '@angular/core';
import { CatalogService } from './service/catalog';
import { Catalog } from './service/catalog';
import { CatalogItem } from './service/catalog';

@Component({
    selector: 'nerv-startmenu',
    templateUrl: 'app/application/startmenu.html',
    outputs: ['eventDockCatalogItem:dock']
})
export class StartMenu {
    title = '产品与服务';
    catalogs: Catalog[];
    catalogSvc: CatalogService;
    display: boolean;
    eventDockCatalogItem = new EventEmitter<CatalogItem>();

    constructor(catalogSvc: CatalogService) {
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