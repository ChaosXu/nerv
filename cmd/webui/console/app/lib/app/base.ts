import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import 'rxjs/add/operator/toPromise';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';
import { ConfigService } from '../config/config.service';
import { RestyService } from '../resty/resty.service';
import { ModalConfirm } from '../form/confirm.modal';
import { Form } from '../form/model';

export abstract class FormsBaseComponent implements OnInit {
    title: string;
    list: any;
    private app: string;
    private type: string;

    constructor(
        private configService: ConfigService,
        private modalService: NgbModal,
        private router: Router,
        private route: ActivatedRoute,
        private resty: RestyService
    ) { }

    ngOnInit(): void {
        this.route.parent.parent.params.forEach((params: Params) => {
            this.app = params['app']
        });
        this.route.params.forEach((params: Params) => {
            this.type = params['type']
        });

        this.title = this.configService.get(this.app)[this.type]['list'].title
        this.load();
    }

    onAdd(): void {
        this.router.navigate(['add'], { relativeTo: this.route });
    }

    onShow(item: {}): void {
        this.router.navigate([item['ID']], { relativeTo: this.route });
    }

    onEdit(item: {}): void {
        this.router.navigate([item['ID'], 'edit'], { relativeTo: this.route });
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
        this.resty.remove(this.type, item['ID'])
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
        this.resty.find(this.type)
            .then(response => this.list = response.data)
            .catch((error) => this.error('加载错误', `加载列表${this.type}失败\r\n${error}`));
    }
}





export abstract class FormBaseComponent implements OnInit {
    form: Form;
    title: string;
    data = {};
    private app: string;
    private type: string;
    private id: number;

    constructor(
        private mode: string,
        private configService: ConfigService,
        private modalService: NgbModal,
        private router: Router,
        private route: ActivatedRoute,
        private resty: RestyService
    ) { }

    ngOnInit(): void {
        this.route.parent.parent.parent.params.forEach((params: Params) => {
            this.app = params['app']
        });
        this.route.params.forEach((params: Params) => {            
            this.type = params['type']
            this.id = +params['id'];
        });
        const config = this.configService.get(this.app)[this.type][this.mode];
        this.form = config.form;
        this.title = config.title;
        if (this.id) {
            this.load();
        }
    }

    onUpdate(): void {
        this.resty.update(`${this.type}`, this.data)
            .then(() => this.info('更新成功', `对象${this.data['Name']}已更新`))
            .catch((error) => this.error('更新错误', `更新对象${this.data['Name']}失败\r\n${error}`));
    }

    onCreate(): void {
        this.resty.create(`${this.type}`, this.data)
            .then(() => this.info('创建成功', `对象${this.data['Name']}已创建`))
            .catch((error) => this.error('创建错误', `创建对象${this.data['Name']}失败\r\n${error}`));
    }

    onBack(url: string): void {
        this.router.navigate([url], { relativeTo: this.route });
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

    private load(): void {
        if (this.id) {
            this.resty.get(`${this.type}`, this.id)
                .then((data) => this.data = data)
                .catch((error) => this.error('加载错误', `加载对象${this.data['Name']}失败\r\n${error}`));
        }
    }
}

