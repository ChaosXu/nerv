import { Injectable } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

export class Field {
    name: string;
    label: string;
    control: string;
    type: string;
    required: boolean;

}

export class Textbox extends Field {
    control = 'textbox';
    type = 'string'
}

export class Form {
    name: string;
    fields: Field[];
}


@Injectable()
export class FormService {
    constructor() { }

    toFormGroup(form: Form) {
        let group: any = {};

        form.fields.forEach(field => {
            group[field.name] = field.required ? new FormControl('', Validators.required)
                : new FormControl('');
        });

        return new FormGroup(group);
    }
}