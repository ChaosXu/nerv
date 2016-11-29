import { Routes } from '@angular/router';
import { AuthGuard } from '../login/auth.guard';
import { ListApp } from '../lib/app/list.app';

export const routes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'dashboard', loadChildren: 'app/dashboard/module#DashboardModule', canActivate: [AuthGuard] },
  { path: 'monitor', loadChildren: 'app/monitor/module#MonitorModule', canActivate: [AuthGuard] },
  { path: ':app', loadChildren: 'app/lib/app/module#AppModule', canActivate: [AuthGuard] }
];