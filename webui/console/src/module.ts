import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule,Routes } from '@angular/router';
import { UserModule } from './app/user/module';
import { DashboardModule } from './app/dashboard/module';
import { MonitorModule } from './app/monitor/module';
import { Application } from './ui/application';
import { StartMenu } from './ui/startmenu';
import { Dock } from './ui/dock';
import { CatalogService } from './service/catalog';

const routes: Routes = [
    { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
];

@NgModule({
  imports: [
    BrowserModule,
    RouterModule.forRoot(routes),
    DashboardModule,
    UserModule,    
    MonitorModule
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
