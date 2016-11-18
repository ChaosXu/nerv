import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { ConfigService } from '../config/config.service';

export interface Menu {
    name: string;
    url: string;
}

@Component({
    selector: 'nerv-list-app',
    templateUrl: 'app/lib/app/list.app.html',
})
export class ListApp {
    menus: Menu[];
    private app: string;
    private type: string;

    constructor(
        configService: ConfigService,
        router: Router,
        route: ActivatedRoute,
    ) {
        route.params.forEach((params: Params) => {
            this.app = params['app']
            this.type = params['type']
        });
        this.menus = configService.get(this.app)['menus'];
    }
}