import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RestyService } from '../lib/resty/resty.service';
import { Form, Field } from '../lib/form/model';
import { form } from './account_form';

@Component({
    templateUrl: 'app/user/account_add.html',
})
export class AccountAddComponent implements OnInit {
    form = form;
    data = {};

    constructor(
        private router: Router,
        private resty: RestyService
    ) { }
    ngOnInit(): void { }

    onSave(): void {
        this.resty.create('Account', this.data)
            .then(() => {
                alert('save ok');
            })
            .catch(this.error);
    }

    onBack(): void {
        this.router.navigate(['/user/account']);
    }

    private error(error: any) {
        return alert(error.text());
    }
}