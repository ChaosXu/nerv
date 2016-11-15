import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { RestyService } from '../lib/resty/resty.service';
// import { Overlay } from 'angular2-modal';
// import { Modal } from 'angular2-modal/plugins/bootstrap';


@Component({
    selector: 'nerv-app-user-account',
    templateUrl: 'app/user/account.html',
})
export class AccountComponent implements OnInit {
    accounts: any;

    constructor(
        private resty: RestyService
    ) {

    }

    ngOnInit(): void {
        this.resty.find('Account')
            .then(response => this.accounts = response.data)
            .catch(this.error);
    }

    onAdd() {

    }

    private error(error: any) {
        return alert(error.text());
    }
}