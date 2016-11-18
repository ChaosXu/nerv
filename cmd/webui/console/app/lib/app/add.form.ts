import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ConfigService } from '../config/config.service';
import { RestyService } from '../resty/resty.service';
import { FormBaseComponent} from './base';
import { Form,Field } from '../form/model';

@Component({
    templateUrl: 'app/lib/app/add.form.html',
})
export class AddComponent extends FormBaseComponent {

    constructor(
        configService: ConfigService,
        modalService: NgbModal,
        router: Router,
        route: ActivatedRoute,
        resty: RestyService
    ) {
        super('add', configService, modalService, router, route, resty);
    }
}