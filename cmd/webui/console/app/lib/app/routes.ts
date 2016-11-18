import { Routes } from '@angular/router';
import { ListApp } from './list.app';
import { ListComponent } from './list.form';
import { AddComponent } from './add.form';
import { EditComponent } from './edit.form';
import { DetailComponent } from './detail.form';

export const routes: Routes = [
    {
        path: '', component: ListApp,
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
        ]
    }
];