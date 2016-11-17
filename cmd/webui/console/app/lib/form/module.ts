import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { FormComponent } from './form.component';
import { FieldComponent } from './field.component';
import { ModalConfirm } from './confirm.modal';

@NgModule({
    imports: [
        CommonModule,
        ReactiveFormsModule,
        NgbModule,
        FormsModule
    ],
    declarations: [
        FormComponent,
        FieldComponent,
        ModalConfirm
    ],
    exports: [
        FormComponent,
        FieldComponent
    ],
    entryComponents: [
        ModalConfirm
    ]
})
export class FormModule { }
