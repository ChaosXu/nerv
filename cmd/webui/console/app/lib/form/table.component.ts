import { Component, Input, Output, OnInit, EventEmitter } from '@angular/core';


@Component({
    //moduleId: module.id,
    selector: 'nerv-table',
    templateUrl: 'app/lib/form/table.component.html',
})
export class TableComponent implements OnInit {
    @Input() class: string;
    @Input() columns: Array<any>;
    @Input() model: any;
    @Input() enablePaging: boolean = true;
    @Input() enableSort: boolean = true;

    @Output() show: EventEmitter<any> = new EventEmitter();
    @Output() remove: EventEmitter<any> = new EventEmitter();
    @Output() edit: EventEmitter<any> = new EventEmitter();
    @Output() sort: EventEmitter<any> = new EventEmitter();
    @Output() paging: EventEmitter<any> = new EventEmitter();
    @Output() pageSize: EventEmitter<any> = new EventEmitter();

    sortBy: { column: string, asc: boolean } = { column: '', asc: true };

    ngOnInit(): void {

    }

    onShow(item: {}) {
        this.show.emit(item);
    }

    onRemove(item: {}) {
        this.remove.emit(item);
    }

    onEdit(item: {}) {
        this.edit.emit(item);
    }

    onSort(column: string) {
        if (!this.enableSort || !column) return;
        this.sortBy = { column: column, asc: !this.sortBy.asc };
        this.sort.emit(this.sortBy);
    }

    onPaging(page: number) {
        this.paging.emit(page);
    }

    onPageSize(size: number) {
        this.pageSize.emit(size);
    }
}