import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import 'rxjs/add/operator/toPromise';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';
import { FormConfig } from '../config/form.config';
import { FormRegistry } from '../form/form.registry';
import { RestyService } from '../resty/resty.service';
import { Form } from '../form/model';
import { ConfirmModal } from '../form/confirm.modal';

export abstract class FormBaseComponent implements OnInit {
    form: Form;
    title: string;
    data = {};
    private app: string;
    private type: string;
    private id: number;

    constructor(
        private mode: string,
        private formRegistry: FormRegistry,
        private configService: FormConfig,
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
        this.form = this.formRegistry.get(config.form);
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
            .then(() => this.info('更新', `更新${this.data['Name'] || this.data['name']}成功`))
            .catch((error) => this.error('错误', `更新${this.data['Name'] || this.data['name']}失败\r\n${error}`));
    }

    onCreate() {
        this.resty.create(`${this.type}`, this.data)
            .then(() => this.info('创建', `创建${this.data['Name'] || this.data['name']}成功`))
            .catch((error) => this.error('错误', `创建${this.data['Name'] || this.data['name']}失败\r\n${error}`));
    }

    onBack(url: string) {
        this.router.navigate([url], { relativeTo: this.route });
    }

    private error(title: string, error: any): void {
        const modalRef = this.modalService.open(ConfirmModal);
        modalRef.componentInstance.title = title;
        modalRef.componentInstance.message = error;
        modalRef.componentInstance.buttons = { ok: true, cancel: false };
    }

    private info(title: string, message: string): void {
        const modalRef = this.modalService.open(ConfirmModal);
        modalRef.componentInstance.title = title;
        modalRef.componentInstance.message = message;
        modalRef.componentInstance.buttons = { ok: true, cancel: false };
    }

    private load(): void {
        if (this.id) {
            this.resty.get(`${this.type}`, this.id)
                .then((data) => this.data = data)
                .catch((error) => this.error('错误', `加载${this.type}失败\r\n${error}`));
        }
    }
}

