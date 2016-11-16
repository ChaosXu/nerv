import { NgModule } from '@angular/core';
import { LibModule } from '../lib/module';
import { RouterModule, Routes } from '@angular/router';
import { UserApp } from './app';
import { RoleComponent } from './role';
import { AccountsComponent } from './accounts';
import { AccountComponent } from './account';
import { AccountAddComponent } from './account_add';
import { AccountEditComponent } from './account_edit';

const routes: Routes = [
    {
        path: '', component: UserApp,
        children: [
            { path: ''}, 
            {
                path: 'account',
                children: [
                    { path: '', component: AccountsComponent },
                    { path: 'add', component: AccountAddComponent },
                    {
                        path: ':id',
                        children: [
                            { path: '', component: AccountComponent },
                            { path: 'edit', component: AccountEditComponent }
                        ]
                    }
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
        AccountsComponent,
        AccountComponent,
        AccountAddComponent,
        AccountEditComponent,
        RoleComponent
    ]
})
export class UserModule { }
