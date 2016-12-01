import { Component, Input, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Form, Field } from './model';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { FormConfig } from '../config/form.config';
import { RestyService } from '../resty/resty.service';
import { ConfirmModal } from '../form/confirm.modal';
import { FormModal } from '../form/form.modal';

@Component({
    //moduleId: module.id,
    selector: 'nerv-table-field',
    templateUrl: 'app/lib/form/table.field.html',
})
export class TableField implements OnInit {

    @Input() field: Field;
    @Input() data: {};

    constructor(
        private configService: FormConfig,
        private modalService: NgbModal,
        private resty: RestyService
    ) { }

    ngOnInit(): void {
    }

    onAdd() {
        //this.configService.get(this.field.)
        
    }

    onShow(item: {}) {
        alert('onShow ' + JSON.stringify(item));
    }

    onEdit(item: {}) {
        alert('onShow ' + JSON.stringify(item));
    }

    onRemove(item: {}) {
        alert('onShow ' + JSON.stringify(item));
        // const modalRef = this.modalService.open(ConfirmModal);
        // modalRef.componentInstance.title = '删除';
        // modalRef.componentInstance.message = `删除${item['Name'] || item['name']}?`;
        // modalRef.result.then((result) => {
        //     if (result == 'ok') {
        //         this.remove(item);
        //     }
        // });
    }

    onSort(column: { column: string, asc: boolean }) {
        // this.sortBy = column;
        // this.load();
    }

    onPaging(page: number) {
        // this.offset = page;
        // this.load();
    }

    onPageSize(size: number) {
        // this.limit = size;
        // this.load();
    }

    private remove(item: {}): void {
        // this.resty.remove(this.type, item['ID'])
        //     .then(() => {
        //         this.load();
        //     })
        //     .catch((error) => this.error('删除错误', `创建对象${item['Name']}失败\r\n${error}`));
    }

    private error(title: string, error: any): void {
        const modalRef = this.modalService.open(ConfirmModal);
        modalRef.componentInstance.title = title;
        modalRef.componentInstance.message = error.toString();
        modalRef.componentInstance.buttons = { ok: true, cancel: false };
    }

    private load(): void {
        // let order;
        // if (this.sortBy) {
        //     order = `${this.sortBy.column} ${this.sortBy.asc ? 'asc' : 'desc'}`;
        // }
        // this.resty.find(this.type, null, order, this.offset, this.limit)
        //     .then(response => this.data = response)
        //     .catch((error) => this.error('加载错误', `加载列表${this.type}失败\r\n${error}`));
    }
}