import { Component, ViewContainerRef, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { ConfigService } from '../config/config.service';
import { RestyService } from '../resty/resty.service';
import { ModalConfirm } from '../form/confirm.modal';

@Component({
    selector: 'nerv-app-user-account',
    templateUrl: 'app/lib/app/list.component.html',
})
export class ListComponent {

    title: string;
    data: any;
    columns: Array<any>;

    private app: string;
    private type: string;
    private sortBy: { column: string, asc: boolean };
    private offset = 0;
    private limit = 10;

    constructor(
        private configService: ConfigService,
        private modalService: NgbModal,
        private router: Router,
        private route: ActivatedRoute,
        private resty: RestyService
    ) { }

    ngOnInit() {
        this.route.parent.parent.params.forEach((params: Params) => {
            this.app = params['app']
        });
        this.route.params.forEach((params: Params) => {
            this.type = params['type']
        });

        const config = this.configService.get(this.app)[this.type]['list'];
        this.title = config.title;
        this.columns = config.columns;
        this.load();
    }

    onAdd() {
        this.router.navigate(['add'], { relativeTo: this.route });
    }

    onShow(item: {}) {
        this.router.navigate([item['ID']], { relativeTo: this.route });
    }

    onEdit(item: {}) {
        this.router.navigate([item['ID'], 'edit'], { relativeTo: this.route });
    }

    onRemove(item: {}) {
        const modalRef = this.modalService.open(ModalConfirm);
        modalRef.componentInstance.title = '删除';
        modalRef.componentInstance.message = `删除对象${item['Name']}?`;
        modalRef.result.then((result) => {
            if (result == 'ok') {
                this.remove(item);
            }
        });
    }

    onSort(column: { column: string, asc: boolean }) {
        this.sortBy = column;
        this.load();
    }

    onPaging(page: number) {
        this.offset = page;
        this.load();
    }

    onPageSize(size:number) {
        this.limit = size;
        this.load();
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
        let order;
        if (this.sortBy) {
            order = `${this.sortBy.column} ${this.sortBy.asc ? 'asc' : 'desc'}`;
        }
        this.resty.find(this.type, null, order, this.offset, this.limit)
            .then(response => this.data = response)
            .catch((error) => this.error('加载错误', `加载列表${this.type}失败\r\n${error}`));
    }

}