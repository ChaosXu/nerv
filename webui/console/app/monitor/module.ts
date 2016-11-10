import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CommonModule } from '@angular/common';
import { MonitorApp } from './monitor';

const routes: Routes = [
    { path: '', component: MonitorApp }
];

@NgModule({
    imports: [
        CommonModule,
        RouterModule.forChild(routes)
    ],
    exports: [RouterModule],
    declarations: [MonitorApp]
})
export class MonitorModule { }
