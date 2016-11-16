import {Form} from '../lib/form/model';

export const form: Form = {
    name: "form_user_add",
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