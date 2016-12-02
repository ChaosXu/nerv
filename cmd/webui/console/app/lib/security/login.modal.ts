import { Component, Input } from '@angular/core';
import { RestyService } from '../resty/resty.service';
import { NgbModal, NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';
import { ConfirmModal } from '../form/confirm.modal';

const form = {
    name: "login_form",
    fields: [
        {
            name: "Name", label: "用户名", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "Password", label: "密码", control: "password", type: "string", validators: { 'required': { message: '不能为空' } }
        },
    ]
};

@Component({
    selector: 'nerv-modal-confirm',
    templateUrl: 'app/lib/security/login.modal.html'
})
export class LoginModal {
    title = '登录';
    form = form;
    data = {};
    doing = false;

    constructor(
        private resty: RestyService,
        private activeModal: NgbActiveModal,
        private modalService: NgbModal
    ) { }

    onLogin() {
        this.doing = true;
        this.title = '登录...';
        this.resty.create('Login', this.data)
            .then(() => {
                this.doing = false;
                this.activeModal.close(this.data);
            })
            .catch((error) => {
                this.title = '登录';
                this.doing = false;                
                this.error('登录错误', '用户名和密码错误');
            });
    }

    private error(title: string, error: any): void {
        const modalRef = this.modalService.open(ConfirmModal, { backdrop: 'static' });
        modalRef.componentInstance.title = title;
        modalRef.componentInstance.message = error;
        modalRef.componentInstance.buttons = { ok: true, cancel: false };
    }
}