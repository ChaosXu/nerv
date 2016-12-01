import { Component, Input, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Form, Field } from './model';

@Component({
    //moduleId: module.id,
    selector: 'nerv-table-field',
    templateUrl: 'app/lib/form/table.field.html',
})

export class TableField implements OnInit {
     
     ngOnInit(): void {
        //const control = this.formGroup.get(this.field.name);  
          
       // control.valueChanges.subscribe(value => this.onValueChanges(value));
    }
}