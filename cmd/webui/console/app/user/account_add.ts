import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { RestyService } from '../lib/resty/resty.service';
import { Form, Field } from '../lib/form/form.service';

const form: Form = {
    name: "form_user_add",
    fields: [
        { name: "name", label: "用户名", control: "text", type: "string", required: true },
        { name: "nick", label: "昵称", control: "text", type: "string", required: true },
        { name: "mail", label: "邮件", control: "email", type: "string", required: true },
        { name: "phone", label: "电话", control: "text", type: "long", required: true }
    ]
};


@Component({
    templateUrl: 'app/user/account_add.html',
})
export class AccountAddComponent implements OnInit {
    form = form;
    data = {};

    constructor(private resty: RestyService) { }
    ngOnInit(): void { }

    onSave(): void {
        this.resty.create('Account', this.data)
            .then(() => {
                alert('save ok');
            })
            .catch(this.error);
    }

    private error(error: any) {
        return alert(error.text());
    }
}