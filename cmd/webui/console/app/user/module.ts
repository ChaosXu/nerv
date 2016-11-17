import { NgModule } from '@angular/core';
import { LibModule } from '../lib/module';
import { RouterModule, Routes } from '@angular/router';
import { UserApp } from './app';
import { ListComponent } from '../lib/form/list';
import { AddComponent } from '../lib/form/add';
import { DetailComponent } from '../lib/form/detail';
import { EditComponent } from '../lib/form/edit';
import { ModelService, Form } from '../lib/form/forms';

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
        ]
    }
];

const form: Form = {
    name: "user_form",
    fields: [
        {
            name: "Name", label: "用户名", control: "text", type: "string", validators: {
                'required': '不能为空'
            }
        },
        {
            name: "Nick", label: "昵称", control: "text", type: "string", validators: {
                'required': '不能为空'
            }
        },
        { name: "Mail", label: "邮件", control: "email", type: "string" },
        { name: "Phone", label: "电话", control: "text", type: "long" }
    ]
};

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
        EditComponent
    ]
})
export class UserModule {

    constructor(modelService: ModelService) {
        modelService.put('Account', {
            list: { title: '使用系统的用户' },
            add: { title: '添加', form: form },
            edit: { title: '添加', form: form },
            detail: { title: '添加', form: form }
        });
        // modelService.put('Role', {
        //     list: { title: '用户角色' },
        //     add: { title: '添加', form: form },
        //     edit: { title: '添加', form: form },
        //     detail: { title: '添加', form: form }
        // });
    }
}
