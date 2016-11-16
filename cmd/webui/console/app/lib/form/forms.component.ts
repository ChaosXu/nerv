import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RestyService } from '../resty/resty.service';

export abstract class FormsComponent implements OnInit {
    list: any;
    //router: Router;
    // resty: RestyService;
    // config: {
    //     prefix: string,
    //     type: string
    // };
    constructor(
        private router: Router,
        private resty: RestyService,
        private config: {
            prefix: string,
            type: string
        }
    ) {
        //this.router = router;
        this.resty = resty;
        this.config = config;
    }

    ngOnInit(): void {
        this.resty.find(this.config.type)
            .then(response => this.list = response.data)
            .catch(this.error);
    }

    onAdd(): void {
        this.router.navigate([`${this.config.prefix}/${this.config.type.toLowerCase()}/add`]);
    }

    onShow(item: {}): void {
        this.router.navigate([`${this.config.prefix.toLowerCase()}/${this.config.type.toLowerCase()}`, item['ID']]);
    }

    onEdit(item: {}): void {
        this.router.navigate([`${this.config.prefix.toLowerCase()}/${this.config.type.toLowerCase()}`, item['ID'], 'edit']);
    }

    onRemove(item: {}): void {

    }

    private error(error: any) {
        return alert(error.text());
    }
}