import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { TableComponent } from './table.component';
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
        TableComponent,
        ModalConfirm
    ],
    exports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        FormComponent,
        FieldComponent,
        TableComponent
    ],
    entryComponents: [
        ModalConfirm
    ]
})
export class FormModule { }
