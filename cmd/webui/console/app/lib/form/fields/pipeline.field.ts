import { Component, Input, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Form, Field } from '../model';

@Component({
    //moduleId: module.id,
    selector: 'nerv-pipeline-field',
    templateUrl: 'app/lib/form/fields/pipeline.field.html',
})
export class PipelineField implements OnInit {
    @Input('readonly') enableReadonly: boolean = false;
    
    @Input() field: Field;
    @Input() data: {};
    error: string = '';

    get value(): any {
        return this.data[this.field.name];
    }

    get required(): boolean {
        return this.field.validators ? this.field.validators['required'] : false;
    }

    ngOnInit(): void {
        
    }    
}