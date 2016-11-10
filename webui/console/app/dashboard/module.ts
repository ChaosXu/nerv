import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
import { DashboardApp } from './dashboard';

const routes: Routes = [
    { path: '', component: DashboardApp }
];

@NgModule({
    imports: [
        CommonModule,
        RouterModule.forChild(routes)
    ],
    exports: [RouterModule],
    declarations: [DashboardApp]
})
export class DashboardModule { }
