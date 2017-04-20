import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Form } from '../lib/form/model';
import { FormConfig } from '../lib/config/form.config';
import { FormRegistry } from '../lib/form/form.registry';
import { AppModule } from '../lib/app/module';
import { routes } from '../lib/app/routes';

const form: Form = {
    name: "userpwd_form",
    fields: [        
        {
            name: "type", label: "类型", control: "select", type: "string",
            display: {
                options: [
                    { label: 'SSH', value: 'ssh', default: true },
                    { label: 'Git', value: 'git' }
                ]
            }
        },
        { name: "name", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } } },
        { name: "user", label: "用户名", control: "text", type: "string", validators: { 'required': { message: '不能为空' } } },
        { name: "password", label: "密码", control: "password", type: "string", validators: { 'required': { message: '不能为空' } } }
    ]
};

@NgModule({
    imports: [
        AppModule,
        RouterModule.forChild(routes)
    ]
})
export class CredentialModule {

    constructor(
        formRegistry: FormRegistry,
        configService: FormConfig
    ) {
        formRegistry.put('credential.userpwd.add', form);
        formRegistry.put('credential.userpwd.edit', form);
        formRegistry.put('credential.userpwd.detail', form);

        configService.put('credential',
            {
                title: '访问凭据',
                menus: [
                    {
                        name: "用户密码",
                        url: "/credential/Credential"
                    }
                ],
                Credential: {
                    list: {
                        title: '用户密码列表', columns: [
                            { label: '类型', name: 'type' },
                            { label: '名称', name: 'name' },
                            { label: '用户名', name: 'user' },                            
                            { label: '操作' }
                        ]
                    },
                    add: { title: '添加密码', form: 'credential.userpwd.add' },
                    edit: { title: '编辑密码', form: 'credential.userpwd.edit' },
                    detail: { title: '查看密码', form: 'credential.userpwd.detail' }
                }
            });
    }
}
