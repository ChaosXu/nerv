import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';
import { RestyService } from '../resty/resty.service';
import { ModalConfirm } from './confirm.modal';

export abstract class FormsBaseComponent implements OnInit {
    list: any;

    constructor(
        private modalService: NgbModal,
        private router: Router,
        private resty: RestyService,
        private config: {
            prefix: string,
            type: string
        }
    ) {
    }

    ngOnInit(): void {
        this.load();
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
        const modalRef = this.modalService.open(ModalConfirm);
        modalRef.componentInstance.title = '删除';
        modalRef.componentInstance.message = `删除对象${item['Name']}?`;
        modalRef.result.then((result) => {
            if (result == 'ok') {
                this.remove(item);
            }
        });
    }

    private remove(item: {}): void {
        this.resty.remove(this.config.type, item['ID'])
            .then(() => {
                this.load();
            })
            .catch((error) => this.error('删除错误', `创建对象${item['Name']}失败\r\n${error}`));
    }

    private error(title: string, error: any): void {
        const modalRef = this.modalService.open(ModalConfirm);
        modalRef.componentInstance.title = title;
        modalRef.componentInstance.message = error.toString();
        modalRef.componentInstance.buttons = { ok: true, cancel: false };
    }

    private load(): void {
        this.resty.find(this.config.type)
            .then(response => this.list = response.data)
            .catch((error) => this.error('加载错误', `加载列表${this.config.type}失败\r\n${error}`));
    }
}



import { ActivatedRoute, Params } from '@angular/router';

export abstract class FormBaseComponent implements OnInit {
    data = {};

    constructor(
        private modalService: NgbModal,
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
                .then((data) => this.data = data)
                .catch((error) => this.error('加载错误', `加载对象${this.data['Name']}失败\r\n${error}`));
        });
    }

    onUpdate(): void {
        this.resty.update(`${this.config.type}`, this.data)
            .then(() => this.info('更新成功', `对象${this.data['Name']}已更新`))
            .catch((error) => this.error('更新错误', `更新对象${this.data['Name']}失败\r\n${error}`));
    }

    onCreate(): void {
        this.resty.create(`${this.config.type}`, this.data)
            .then(() => this.info('创建成功', `对象${this.data['Name']}已创建`))
            .catch((error) => this.error('创建错误', `创建对象${this.data['Name']}失败\r\n${error}`));
    }

    onBack(): void {
        this.router.navigate([`/${this.config.prefix}/${this.config.type.toLowerCase()}`]);
    }

    private error(title: string, error: any): void {
        const modalRef = this.modalService.open(ModalConfirm);
        modalRef.componentInstance.title = title;
        modalRef.componentInstance.message = error;
        modalRef.componentInstance.buttons = { ok: true, cancel: false };
    }

    private info(title: string, message: string): void {
        const modalRef = this.modalService.open(ModalConfirm);
        modalRef.componentInstance.title = title;
        modalRef.componentInstance.message = message;
        modalRef.componentInstance.buttons = { ok: true, cancel: false };
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
