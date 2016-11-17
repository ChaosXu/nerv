import { NgModule } from '@angular/core';
import { LibModule } from '../lib/module';
import { RouterModule, Routes } from '@angular/router';
import { UserApp } from './app';
import { RoleComponent } from './role';
import { ListComponent } from '../lib/form/list'
import { AddComponent } from '../lib/form/add'
import { DetailComponent } from '../lib/form/detail'
import { EditComponent } from '../lib/form/edit'

const routes: Routes = [
    {
        path: '', component: UserApp,
        children: [
            {
                path: ''
            },
            {
                path: ':type',
                children: [
                    { path: '', component: ListComponent },
                    { path: 'add', component: AddComponent },
                    { path: ':id', component: DetailComponent },
                    { path: ':id/edit', component: EditComponent }
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
        ListComponent,
        DetailComponent,
        AddComponent,
        EditComponent,
        RoleComponent
    ]
})
export class UserModule { }
