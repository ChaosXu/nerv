import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { RestyService } from '../resty/resty.service';
import { FormBaseComponent, ModelService, Form, Field } from './forms';

@Component({
    templateUrl: 'app/lib/form/edit.html',
})
export class EditComponent extends FormBaseComponent {
   
    constructor(
        modelService: ModelService,
        modalService: NgbModal,
        router: Router,
        route: ActivatedRoute,
        resty: RestyService
    ) {
        super(modelService, modalService, router, route, resty);
    }
}