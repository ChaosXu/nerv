import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params, UrlSegment } from '@angular/router';
import { FormConfig } from '../config/form.config';

export interface Menu {
    name: string;
    url: string;
}

@Component({
    selector: 'nerv-list-app',
    templateUrl: 'app/lib/app/list.app.html',
})
export class ListApp implements OnInit {
    title: string;
    menus: Menu[];

    constructor(
        private configService: FormConfig,
        router: Router,
        private route: ActivatedRoute,
    ) {

    }

    ngOnInit(): void {
        let app;
        this.route.params.subscribe((params: Params) => {
            app = params['app'];
            if (!app) return;
            const config = this.configService.get(app);
            this.title = config['title'];
            this.menus = config['menus'];
        });

        if (!app) {
            this.route.parent.url.forEach((segment: UrlSegment[]) => {
                app = segment[0].path;
                const config = this.configService.get(app);
                this.title = config['title'];
                this.menus = config['menus'];
            });
        }
    }
}