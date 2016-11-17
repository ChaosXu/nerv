import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { RestyService } from '../lib/resty/resty.service';
import { FormsBaseComponent } from '../lib/form/forms';

@Component({
    selector: 'nerv-app-user-account',
    templateUrl: 'app/user/accounts.html',
})
export class AccountsComponent extends FormsBaseComponent {

    constructor(
        modalService: NgbModal,
        router: Router,
        resty: RestyService
    ) {
        super(modalService, router, resty, {
            prefix: "/user",
            type: 'Account'
        });
    }
}