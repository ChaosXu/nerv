import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { Application } from './application';
import { StartMenu } from './startmenu';
import { Dock } from './dock';
import { CatalogService } from './service/catalog';

const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'dashboard', loadChildren: 'app/dashboard/module#DashboardModule' },
  { path: 'user', loadChildren: 'app/user/module#UserModule' },
  { path: 'monitor', loadChildren: 'app/monitor/module#MonitorModule' }
];

@NgModule({
  imports: [
    BrowserModule,
    RouterModule.forRoot(routes),
    NgbModule.forRoot()        
  ],
  declarations: [
    Application,
    StartMenu,
    Dock
  ],
  providers: [CatalogService],
  bootstrap: [Application]
})
export class AppModule { }
