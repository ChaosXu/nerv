import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule,FormsModule } from '@angular/forms';
import { FormComponent } from './form.component';
import { FieldComponent } from './field.component';

@NgModule({
    imports: [
        CommonModule,
        ReactiveFormsModule,
        FormsModule    
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
