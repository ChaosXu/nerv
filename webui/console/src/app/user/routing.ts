import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserApp } from './user';

const routes: Routes = [    
    { path: 'user', component: UserApp }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class UserRouting { }
