import { Routes } from '@angular/router';
import { AuthGuard } from '../lib/security/auth.guard';
import { routes } from '../lib/app/routes';

export const rootRoutes: Routes = [
  { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
  { path: 'dashboard', loadChildren: 'app/dashboard/module#DashboardModule', canActivate: [AuthGuard] },
  { path: 'monitor', loadChildren: 'app/monitor/module#MonitorModule', canActivate: [AuthGuard] },
  { path: 'orchestration', loadChildren: 'app/orchestration/module#OrchestrationModule', canActivate: [AuthGuard] },
  { path: 'pipeline', loadChildren: 'app/pipeline/module#PipelineModule', canActivate: [AuthGuard] },
  { path: 'credential', loadChildren: 'app/credential/module#CredentialModule', canActivate: [AuthGuard] },
  { path: 'user', loadChildren: 'app/user/module#UserModule', canActivate: [AuthGuard] },
  { path: 'infrastructure', loadChildren: 'app/infrastructure/module#InfrastructureModule', canActivate: [AuthGuard] } 
];