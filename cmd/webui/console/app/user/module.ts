import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Form } from '../lib/form/model';
import { ConfigService } from '../lib/config/config.service';

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
})
export class UserModule {

    constructor(configService: ConfigService) {
        configService.put('user',
            {
                menus: [
                    {
                        name: "人员管理",
                        url: "/user/Account"
                    }, {
                        name: "权限管理",
                        url: "/user/Role"
                    }
                ],
                Account: {
                    list: {
                        title: '人员列表', columns: [
                            { label: '用户名' },
                            { label: '昵称' },
                            { label: '邮件' },
                            { label: '电话' },
                            { label: '操作' }
                        ]
                    },
                    add: { title: '添加人员', form: form },
                    edit: { title: '编辑人员', form: form },
                    detail: { title: '查看人员', form: form }
                }
            });
    }
}
