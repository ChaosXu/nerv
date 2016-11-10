import { Component } from '@angular/core';
import { OnInit } from '@angular/core';
import { CatalogService } from '../service/catalog';
import { CatalogItem } from '../service/catalog';

@Component({
    selector: 'nerv-dock',
    templateUrl: 'view/dock.html',
    inputs: ['docks']    
})
export class Dock implements OnInit {
    catalogSvc: CatalogService;
    docks: CatalogItem[] = [];

    constructor(catalogSvc: CatalogService) {
        this.catalogSvc = catalogSvc
    }

    ngOnInit(): void {

    }
}