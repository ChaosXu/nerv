import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { Console } from './console/console';
import { Catalog } from './catalog/catalog';

@NgModule({
  imports: [BrowserModule],
  declarations: [
    Console,
    Catalog
  ],
  bootstrap: [Console]
})
export class AppModule { }
