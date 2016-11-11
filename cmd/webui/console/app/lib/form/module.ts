import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormComponent } from './form';
import { FieldComponent } from './field';

@NgModule({
    imports: [
        CommonModule
    ],
    declarations: [
        FormComponent,
        FieldComponent,
    ],
    exports: [
        FormComponent,
        FieldComponent
    ]
})
export class FormModule { }
