import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { Application } from './ui/application';
import { StartMenu } from './ui/startmenu';
import { Dock } from './ui/dock';
import { CatalogService } from './service/catalog';
import { UserApp } from './app/user/user';
import { DashboardApp } from './app/dashboard/dashboard';
import { MonitorApp } from './app/monitor/monitor';
import { Routing } from './routing';
import { UserRouting } from './app/user/routing';
import { DashboardRouting } from './app/dashboard/routing';
import { MonitorRouting } from './app/monitor/routing';


@NgModule({
  imports: [
    BrowserModule,
    Routing,
    DashboardRouting,
    UserRouting,    
    MonitorRouting
  ],
  declarations: [
    Application,
    StartMenu,
    Dock,
    UserApp,
    DashboardApp,
    MonitorApp
  ],
  providers: [CatalogService],
  bootstrap: [Application]
})
export class AppModule { }
