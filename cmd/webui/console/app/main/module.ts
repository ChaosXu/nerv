import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { ConfigModule } from '../lib/config/module';
import { Application } from './application';
import { StartMenu } from './startmenu';
import { Dock } from './dock';
import { CatalogConfig } from '../lib/config/catalog.config';
import { rootRoutes } from './routes';
import { SecurityModule } from '../lib/security/module';

@NgModule({
  imports: [
    BrowserModule,
    RouterModule.forRoot(rootRoutes),
    NgbModule.forRoot(),
    ConfigModule,
    SecurityModule
  ],
  declarations: [
    Application,
    StartMenu,
    Dock
  ],
  providers: [
    CatalogConfig
  ],
  bootstrap: [Application]
})
export class MainModule { }
