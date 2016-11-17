import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { RestyService } from '../resty/resty.service';
import { FormBaseComponent, ModelService, Form, Field } from './forms';


@Component({
    templateUrl: 'app/lib/form/detail.html',
})
export class DetailComponent extends FormBaseComponent {

    constructor(
        modelService: ModelService,
        modalService: NgbModal,
        router: Router,
        route: ActivatedRoute,
        resty: RestyService
    ) {
        super('detail', modelService, modalService, router, route, resty);
    }
}