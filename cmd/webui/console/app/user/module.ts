import { NgModule } from '@angular/core';
import { LibModule } from '../lib/module';
import { RouterModule, Routes } from '@angular/router';
import { UserApp } from './app';
import { RoleComponent } from './role';
import { AccountComponent } from './account';
import { AccountAddComponent } from './account_add';

const routes: Routes = [
    {
        path: '', component: UserApp,
        children: [
            { path: '', redirectTo: "account", pathMatch: "prefix" },
            {
                path: 'account',
                children: [
                    { path: '', component: AccountComponent },
                    { path: 'add', component: AccountAddComponent }
                ]
            },
            { path: 'role', component: RoleComponent }
        ]
    }
];

@NgModule({
    imports: [
        LibModule,
        RouterModule.forChild(routes)
    ],
    exports: [
        RouterModule,
        UserApp
    ],
    declarations: [
        UserApp,
        AccountComponent,
        AccountAddComponent,
        RoleComponent
    ]
})
export class UserModule { }
