import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { RestyService } from '../lib/resty/resty.service';
import { Form, Field } from '../lib/form/model';
import { form } from './account_form';

@Component({
    templateUrl: 'app/user/account_edit.html',
})
export class AccountEditComponent implements OnInit {
    form = form;
    data = {};

    constructor(
        private route: ActivatedRoute,
        private resty: RestyService
    ) { }

    ngOnInit(): void {
        this.route.params.forEach((params: Params) => {
            let id = +params['id'];
            this.resty.get('Account', id)
                .then((data) => {
                    this.data = data;
                })
                .catch(this.error);
        });
    }

    onSave(): void {
        this.resty.update('Account', this.data['ID'], this.data)
            .then(() => {
                alert('save ok');
            })
            .catch(this.error);
    }

    private error(error: any) {
        return alert(error.text());
    }
}