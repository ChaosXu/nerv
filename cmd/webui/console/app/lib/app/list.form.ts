import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ConfigService } from '../config/config.service';
import { RestyService } from '../resty/resty.service';
import { FormsBaseComponent } from './base';

@Component({
    selector: 'nerv-app-user-account',
    templateUrl: 'app/lib/app/list.form.html',
})
export class ListComponent extends FormsBaseComponent {

constructor(  
        configService: ConfigService,  
        modalService: NgbModal,
        router: Router,
        route: ActivatedRoute,
        resty: RestyService       
    ) {
        super(configService,modalService,router,route,resty);
    }
    
}