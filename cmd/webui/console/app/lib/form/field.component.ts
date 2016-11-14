import { Component, Input, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Form, Field } from './form.service';

@Component({
    //moduleId: module.id,
    selector: 'nerv-form-field',
    templateUrl: 'app/lib/form/field.component.html',
})
export class FieldComponent implements OnInit {
    @Input() formGroup: FormGroup;
    @Input() field: Field;
    @Input() data: any;  //form data
    @Input() set value(value: any) {
        if(value) {
            this.data[this.field.name] = value;
        }
    }

    get value(): any {
        return this.data[this.field.name];
    }

    ngOnInit(): void {

    }
}