import { Component, Input, Output, EventEmitter } from '@angular/core';


@Component({
    //moduleId: module.id,
    selector: 'nerv-pagination',
    templateUrl: 'app/lib/form/pagination.component.html',
})
export class PaginationComponent {
    @Input() page: number = 0;
    @Input() set pageCount(value: number) {
        this._pageCount = value;
        this.show();
    }
    get pageCount(): number {
        return this._pageCount;
    }

    @Output() paging: EventEmitter<any> = new EventEmitter();
    @Output() pageSize: EventEmitter<any> = new EventEmitter();

    pages: Array<{ label: string, index: number }> = [];
    pageSizes = [10, 20, 50];
    limit: number = 10;
    private _pageCount: number;
    

    onPaging(page: number) {
        this.page = page;
        this.paging.emit(page);
    }

    onPageSize(size: number) {
        this.limit = size;
        this.pageSize.emit(size);
    }

    private show() {
        this.pages = [];
        for (let i = 0; i < this._pageCount; i++) {
            this.pages.push({ label: (i + 1).toString(), index: i });
        }
    }
}