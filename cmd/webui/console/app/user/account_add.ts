import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { AccountService, Account } from './account.service';
import { Form, Field } from '../lib/form/form.service';

const form: Form = {
    name: "form_user_add",
    fields: [
        { name: "name", label: "用户名", control: "textbox", type: "string", required: true },
        { name: "nick", label: "昵称", control: "textbox", type: "string", required: true },
        { name: "mail", label: "邮件", control: "textbox", type: "string", required: true },
        { name: "phone", label: "电话", control: "textbox", type: "string", required: true }
    ]
};


@Component({
    templateUrl: 'app/user/account_add.html',
})
export class AccountAddComponent implements OnInit {
    form = form;
    data = {};
    accounts: Account[];

    constructor(private accountSvc: AccountService) {

    }

    ngOnInit(): void { }

    onSave(): void {
        //this.accountSvc.create(this.data).then()
    }
}