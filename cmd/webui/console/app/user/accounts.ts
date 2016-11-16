import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RestyService } from '../lib/resty/resty.service';
// import { Overlay } from 'angular2-modal';
// import { Modal } from 'angular2-modal/plugins/bootstrap';


@Component({
    selector: 'nerv-app-user-account',
    templateUrl: 'app/user/accounts.html',
})
export class AccountsComponent implements OnInit {
    accounts: any;

    constructor(
        private router: Router,
        private resty: RestyService
    ) { }

    ngOnInit(): void {
        this.resty.find('Account')
            .then(response => this.accounts = response.data)
            .catch(this.error);
    }

    onAdd(): void {
        this.router.navigate(['/user/account/add']);
    }

    onShow(item: {}): void {
        this.router.navigate(['/user/account', item['ID']]);
    }

    onEdit(item: {}): void {
        this.router.navigate(['/user/account', item['ID'], 'edit']);
    }

    onRemove(item: {}): void {
        
    }

    private error(error: any) {
        return alert(error.text());
    }
}