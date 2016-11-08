import { Component } from '@angular/core';
@Component({
    selector: 'nerv-catalog',
    templateUrl: 'app/catalog/catalog.html'
})
export class Catalog {
    display = false;
    title = '产品与服务';

    onToggle(): void {
        this.display = !this.display;
    }
}