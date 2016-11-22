import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import 'rxjs/add/operator/toPromise';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';
import { ConfigService } from '../config/config.service';
import { RestyService } from '../resty/resty.service';
import { Form } from '../form/model';
import { ModalConfirm } from '../form/confirm.modal';

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

    onAdd() {
        this.router.navigate(['add'], { relativeTo: this.route });
    }

    onUpdate() {
        this.resty.update(`${this.type}`, this.data)
            .then(() => this.info('更新成功', `对象${this.data['Name']}已更新`))
            .catch((error) => this.error('更新错误', `更新对象${this.data['Name']}失败\r\n${error}`));
    }

    onCreate() {
        this.resty.create(`${this.type}`, this.data)
            .then(() => this.info('创建成功', `对象${this.data['Name']}已创建`))
            .catch((error) => this.error('创建错误', `创建对象${this.data['Name']}失败\r\n${error}`));
    }

    onBack(url: string) {
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

