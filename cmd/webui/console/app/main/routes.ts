import { Routes } from '@angular/router';
import { ListApp } from '../lib/app/list.app';

export const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'dashboard', loadChildren: 'app/dashboard/module#DashboardModule' },  
  { path: 'monitor', loadChildren: 'app/monitor/module#MonitorModule' },
  { path: ':app', loadChildren: 'app/lib/app/module#AppModule' }
];