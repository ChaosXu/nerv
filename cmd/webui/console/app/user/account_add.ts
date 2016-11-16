import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { RestyService } from '../lib/resty/resty.service';
import { FormBaseComponent,Form, Field  } from '../lib/form/forms';
import { form } from './account_form';

@Component({
    templateUrl: 'app/user/account_add.html',
})
export class AccountAddComponent extends FormBaseComponent {

    constructor(
        router: Router,
        route: ActivatedRoute,
        resty: RestyService
    ) {
        super(router, route, resty, {
            prefix: "/user",
            type: "Account"
        },
            form
        )
    }
}