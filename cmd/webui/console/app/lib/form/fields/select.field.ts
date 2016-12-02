import { Component, Input, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Form, Field } from '../model';

@Component({
    //moduleId: module.id,
    selector: 'nerv-select-field',
    templateUrl: 'app/lib/form/fields/select.field.html',
})
export class SelectField {
    @Input('readonly') enableReadonly = false;
    @Input() get field(): Field {
        return this._field;
    }

    set field(value: Field) {
        this._field = value;
        let options = value.display['options'];
        for (let opt of options) {
            if (opt.default) {
                this.onSelect(opt);
            }
        }
    }

    @Input() get data(): any {
        return this._data;
    };

    set data(value: any) {
        const field = this._field;
        if (!value || !field) return;
        this._data = value;

        const v = value[field.name];
        if (v) {
            let options = field.display['options'];
            for (let opt of options) {
                if (opt.value == v) {
                    this.onSelect(opt);
                }
            }
        } else if (this.selected) {
            this.onSelect(this.selected);
        }
    }
    private _data: any;

    selected: any;
    private _field: Field;

    onSelect(opt: any): void {
        this.selected = opt;
        if (this._data) {
            this._data[this._field.name] = this.selected['value'];
        }
    }
}