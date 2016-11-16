import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { RestyService } from '../resty/resty.service';

export abstract class FormsBaseComponent implements OnInit {
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



import { ActivatedRoute, Params } from '@angular/router';

export abstract class FormBaseComponent implements OnInit {
    data = {};

    constructor(
        private router: Router,
        private route: ActivatedRoute,
        private resty: RestyService,
        private config: {
            prefix: string,
            type: string,
        },
        public form: Form
    ) { }

    ngOnInit(): void {
        this.route.params.forEach((params: Params) => {
            let id = +params['id'];
            this.resty.get(`${this.config.type}`, id)
                .then((data) => {
                    this.data = data;
                })
                .catch(this.error);
        });
    }

    onUpdate(): void {
        this.resty.update(`${this.config.type}`, this.data)
            .then(() => {
                alert('save ok');
            })
            .catch(this.error);
    }

    onCreate(): void {
        this.resty.create(`${this.config.type}`, this.data)
            .then(() => {
                alert('save ok');
            })
            .catch(this.error);
    }

    onBack(): void {
        this.router.navigate([`/${this.config.prefix}/${this.config.type.toLowerCase()}`]);
    }

    private error(error: any) {
        return alert(error.text());
    }
}

import { Injectable } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

export class Form {
    name: string;
    fields: Field[];
}

export class Field {
    name: string;
    label: string;
    control: string;
    type: string;
    validators?: {};
}

export class Textbox extends Field {
    control = 'textbox';
    type = 'string'
}
