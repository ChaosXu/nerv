import { Component } from '@angular/core';
import { EventEmitter } from '@angular/core';
import { CatalogService } from '../service/catalog';
import { Catalog } from '../service/catalog';
import { CatalogItem } from '../service/catalog';

@Component({
    selector: 'nerv-startmenu',
    templateUrl: 'app/view/startmenu.html',
    outputs: ['eventDockCatalogItem:dock'],
    providers: [CatalogService]
})
export class StartMenu {
    title = '产品与服务';
    display: boolean;
    catalogs: Catalog[];
    catalogSvc: CatalogService;
    eventDockCatalogItem = new EventEmitter<CatalogItem>();

    constructor(catalogSvc: CatalogService) {
        this.catalogSvc = catalogSvc;
    }

    onToggle(): void {
        this.display = !this.display;
        if (this.display) {
            this.catalogSvc.getCatalogs().then(catalogs => this.catalogs = catalogs);
        }
    }

    onDockCatalogItem(item: CatalogItem): void {
        item.dock = !item.dock;
        this.eventDockCatalogItem.emit(item);
    }
}