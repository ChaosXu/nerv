import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RestyService } from '../lib/resty/resty.service';
import { FormsBaseComponent } from '../lib/form/forms';
// import { Modal } from 'angular2-modal/plugins/bootstrap';


@Component({
    selector: 'nerv-app-user-account',
    templateUrl: 'app/user/accounts.html',
})
export class AccountsComponent extends FormsBaseComponent{
    accounts: any;

    constructor(
        router: Router,
        resty: RestyService
    ) {
        super(router, resty, {
            prefix:"/user",
            type:'Account'
        });
    }
}