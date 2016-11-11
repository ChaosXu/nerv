import { Component, ViewContainerRef, OnInit } from '@angular/core';
// import { Overlay } from 'angular2-modal';
// import { Modal } from 'angular2-modal/plugins/bootstrap';
import { AccountService, Account } from './account.service';

@Component({
    selector: 'nerv-app-user-account',
    templateUrl: 'app/user/account.html',
})
export class AccountComponent implements OnInit {

    accounts: Account[];

    constructor(private accountSvc: AccountService) {
        
    }

    ngOnInit(): void {
        this.accountSvc.find().then(accounts => this.accounts = accounts)
    }

    onAdd() {
       
    }
}