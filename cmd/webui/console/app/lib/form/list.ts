import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router,ActivatedRoute } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { RestyService } from '../resty/resty.service';
import { FormsBaseComponent,ModelService } from './forms';

@Component({
    selector: 'nerv-app-user-account',
    templateUrl: 'app/lib/form/list.html',
})
export class ListComponent extends FormsBaseComponent {

constructor(  
        modelService:ModelService,  
        modalService: NgbModal,
        router: Router,
        route: ActivatedRoute,
        resty: RestyService       
    ) {
        super(modelService,modalService,router,route,resty);
    }
    
}