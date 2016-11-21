import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ConfigService } from '../config/config.service';
import { RestyService } from '../resty/resty.service';
import { FormBaseComponent} from './base.form';
import { Form,Field } from '../form/model';


@Component({
    templateUrl: 'app/lib/app/detail.form.html',
})
export class DetailComponent extends FormBaseComponent {

    constructor(
        configService: ConfigService,
        modalService: NgbModal,
        router: Router,
        route: ActivatedRoute,
        resty: RestyService
    ) {
        super('detail', configService, modalService, router, route, resty);
    }
}