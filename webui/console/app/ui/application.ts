import { Component } from '@angular/core';
import { CatalogItem } from '../service/catalog';

@Component({
  selector: 'nerv-console',
  templateUrl: 'app/view/application.html'  
})
export class Application {
  docks: CatalogItem[] = []

  toggleDock(item: CatalogItem): void {
    var index = this.docks.findIndex(e => e.name == item.name);
    if (index > -1) {
      this.docks.splice(index);
    } else {
      this.docks.push(item);
    }
  }
}