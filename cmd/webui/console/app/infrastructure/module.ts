import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Form } from '../lib/form/model';
import { FormConfig } from '../lib/config/form.config';

const hostForm: Form = {
    name: "host_form",
    fields: [
        {
            name: "name", label: "名称", control: "text", type: "string", validators: {
                'required': '不能为空'
            }
        },
        {
            name: "ip", label: "IP", control: "text", type: "string", validators: {
                'required': '不能为空'
            }
        }
    ]
};

@NgModule({
})
export class InfrastructureModule {

    constructor(configService: FormConfig) {
        configService.put('infrastructure',
            {
                title: '基础设施',
                menus: [
                    {
                        name: "物理主机",
                        url: "/infrastructure/Host"
                    }
                ],
                Host: {
                    list: {
                        title: '主机列表', columns: [
                            { label: '名称', name: 'name' },
                            { label: 'IP', name: 'ip' },
                            { label: '操作' }
                        ]
                    },
                    add: { title: '添加主机', form: hostForm },
                    edit: { title: '编辑主机', form: hostForm },
                    detail: { title: '查看主机', form: hostForm }
                }
            });
    }
}
