import { Component } from '@angular/core';
import { CatalogItem } from './service/catalog';

@Component({
  selector: 'nerv-console',
  templateUrl: 'app/main/application.html'
})
export class Application {
  docks: CatalogItem[] = []

  toggleDock(item: CatalogItem): void {
    var index = this.docks.findIndex(e => e.name == item.name);
    if (index > -1) {
      this.docks.splice(index, 1);
    } else {
      this.docks.push(item);
    }
  }
}