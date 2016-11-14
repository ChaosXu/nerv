import { Component, Input, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { FormService, Form } from './form.service';

@Component({
    //moduleId: module.id,
    selector: 'nerv-form',
    templateUrl: 'app/lib/form/form.component.html',
    providers: [
        FormService
    ]
})
export class FormComponent implements OnInit {

    @Input() meta: Form;
    @Input() data: any;
    formGroup: FormGroup;
    debugData: any;

    get valid(): boolean {
        return this.formGroup ? this.formGroup.valid : false;
    }
    
    constructor(private formService: FormService) { }

    ngOnInit(): void {
        this.formGroup = this.formService.toFormGroup(this.meta);
    }
}