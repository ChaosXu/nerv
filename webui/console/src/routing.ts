import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { UserApp } from './app/user/user';
import { Application } from './ui/application';
import { DashboardApp } from './app/dashboard/dashboard';

const routes: Routes = [
    { path: '', redirectTo: '/dashboard', pathMatch: 'full' },
    { path: 'dashboard', component: DashboardApp },
    { path: 'user', component: UserApp },
    //{ path: 'detail/:id', component: HeroDetailComponent },
    //{ path: 'heroes',     component: HeroesComponent }
];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class Routing { }
