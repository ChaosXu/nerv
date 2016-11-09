import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { Application } from './ui/application';
import { StartMenu } from './ui/startmenu';
import { Dock } from './ui/dock';
import { CatalogService } from './service/catalog';
import { UserApp } from './app/user/user';
import { DashboardApp } from './app/dashboard/dashboard';
import { Routing } from './routing';

@NgModule({
  imports: [
    BrowserModule,
    Routing
  ],
  declarations: [
    Application,
    StartMenu,
    Dock,
    UserApp,
    DashboardApp
  ],
  providers: [CatalogService],
  bootstrap: [Application]
})
export class AppModule { }
