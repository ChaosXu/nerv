import { Component } from '@angular/core';
import { CatalogService } from '../../service/catalog';
import { Catalog } from '../../service/catalog';

@Component({
    selector: 'nerv-startmenu',
    templateUrl: 'app/ui/startmenu/startmenu.html',
    providers: [CatalogService]
})
export class StartMenu {
    title = '产品与服务';
    display = false;
    catalogs: Catalog[];
    catalogSvc: CatalogService;

    constructor(catalogSvc: CatalogService) {
        this.catalogSvc = catalogSvc;
    }

    onToggle(): void {
        this.display = !this.display;
        if (this.display) {
            this.catalogSvc.getCatalogs().then(catalogs => this.catalogs = catalogs);
        }
    }
}