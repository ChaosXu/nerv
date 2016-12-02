import { Component, Input, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators, ValidatorFn } from '@angular/forms';
import { Form, Constraint } from './model';
import { ValidatorRegistry } from './validators/validator.registry';

@Component({
    //moduleId: module.id,
    selector: 'nerv-form',
    templateUrl: 'app/lib/form/form.component.html'
})
export class FormComponent implements OnInit {
    @Input() meta: Form;
    @Input('data') set setData(value: {}) {
        this.data = value;
        if (this.formGroup) {
            this.formGroup.reset(this.data);
        }
    }
    data: {};
    formGroup: FormGroup;

    constructor(
        private validatorRegistry: ValidatorRegistry
    )
    { }

    get valid(): boolean {
        return this.formGroup ? this.formGroup.valid : false;
    }


    ngOnInit(): void {
        this.buildForm();
    }

    private buildForm(): void {
        let group: any = {};

        this.meta.fields.forEach(field => {
            //if (field.control != 'table') {
                group[field.name] = new FormControl('', this.getValidators(field.validators));
            //}
        });

        this.formGroup = new FormGroup(group);
    }

    private getValidators(validators?: {}): ValidatorFn[] {
        if (!validators) return null;
        let fns: ValidatorFn[] = [];
        for (let k in validators) {
            let v = this.validatorRegistry.get(k);
            if (v) {
                fns.push(v);
            }
        }

        return fns;
    }
}