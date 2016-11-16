import { Component, Input, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators, ValidatorFn } from '@angular/forms';
import { Form } from './model';

@Component({
    //moduleId: module.id,
    selector: 'nerv-form',
    templateUrl: 'app/lib/form/form.component.html'
})
export class FormComponent implements OnInit {

    @Input() meta: Form;
    @Input() data: {};
    @Input() errorMessages: {}
    formGroup: FormGroup;

    get valid(): boolean {
        return this.formGroup ? this.formGroup.valid : false;
    }

    ngOnInit(): void {
        this.buildForm();
    }

    private buildForm(): void {
        let group: any = {};

        this.meta.fields.forEach(field => {
            group[field.name] = new FormControl('', this.getValidators(field.validators))
        });

        this.formGroup = new FormGroup(group);
    }

    private getValidators(validators?: {}): ValidatorFn[] {
        if (!validators) return null;
        let fns: ValidatorFn[] = [];
        for (let key in validators) {
            let v = this.getValidator(key)
            if (v != null) {
                fns.push(v);
            }
        }

        return fns;
    }

    private getValidator(name: string): ValidatorFn {
        switch (name) {
            case 'required':
                return Validators.required;
            default:
                return null;
        }
    }
}