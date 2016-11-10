import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardApp } from './dashboard';

const routes: Routes = [    
    { path: 'dashboard', component: DashboardApp }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
    declarations: [DashboardApp]
})
export class DashboardModule { }
