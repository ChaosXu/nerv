import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router';
import { UserApp } from './user';
import { RoleComponent } from './role';

const routes: Routes = [
    {
        path: 'user', component: UserApp,
        children: [
            { path: '', },
            { path: 'project', component: UserApp },
            { path: 'account', component: UserApp },
            { path: 'role', component: RoleComponent }
        ]
    }
];

@NgModule({
    imports: [
        BrowserModule,
        RouterModule.forChild(routes)
    ],
    exports: [
        RouterModule,
        UserApp
    ],
    declarations: [
        UserApp,
        RoleComponent
    ],
})
export class UserModule { }
