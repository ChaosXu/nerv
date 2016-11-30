import { Component } from '@angular/core';
import { CatalogItem } from '../lib/config/catalog.config';

@Component({
    selector: 'nerv-dock',
    templateUrl: 'app/main/dock.html',
    inputs: ['docks']
})
export class Dock {    
    docks: Array<CatalogItem> = [];
}