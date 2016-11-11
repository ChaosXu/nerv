import { Component, Input, OnInit } from '@angular/core';
import { Form, Field } from './metadata';

@Component({
    selector: 'nerv-form-field',
    templateUrl: 'app/lib/form/field.html',
})
export class FieldComponent implements OnInit {
    @Input() form: Form;
    @Input() field: Field;
    @Input() data: any;
    
    ngOnInit(): void {

    }
}