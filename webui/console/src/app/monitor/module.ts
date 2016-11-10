import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MonitorApp } from './monitor';

const routes: Routes = [
    { path: 'monitor', component: MonitorApp }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
    declarations: [MonitorApp]
})
export class MonitorModule { }
