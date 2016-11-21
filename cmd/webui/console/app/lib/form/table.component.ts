import { Component, Input, Output, OnInit, EventEmitter } from '@angular/core';


@Component({
    //moduleId: module.id,
    selector: 'nerv-table',
    templateUrl: 'app/lib/form/table.component.html',
})
export class TableComponent implements OnInit {
    @Input() class: string;
    @Input() columns:Array<any>;
    @Input() data: Array<any>;

    @Output() show: EventEmitter<any> = new EventEmitter();
    @Output() remove: EventEmitter<any> = new EventEmitter();
    @Output() edit: EventEmitter<any> = new EventEmitter();

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
}