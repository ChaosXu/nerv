import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { FormConfig } from '../config/form.config';

export interface Menu {
    name: string;
    url: string;
}

@Component({
    selector: 'nerv-list-app',
    templateUrl: 'app/lib/app/list.app.html',
})
export class ListApp {
    title: string;
    menus: Menu[];

    constructor(
        configService: FormConfig,
        router: Router,
        route: ActivatedRoute,
    ) {
        route.params.subscribe((params: Params) => {
            const app = params['app'];
            const config = configService.get(app);
            this.title = config['title'];
            this.menus = config['menus'];
        });
    }
}