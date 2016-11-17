import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { RestyService } from '../resty/resty.service';
import { FormBaseComponent, ModelService, Form, Field } from './forms';

@Component({
    templateUrl: 'app/lib/form/add.html',
})
export class AddComponent extends FormBaseComponent {

    constructor(
        modelService: ModelService,
        modalService: NgbModal,
        router: Router,
        route: ActivatedRoute,
        resty: RestyService
    ) {
        super('add', modelService, modalService, router, route, resty);
    }
}