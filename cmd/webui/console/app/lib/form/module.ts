import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { TableComponent } from './table.component';
import { FormComponent } from './form.component';
import { FieldComponent } from './field.component';
import { ConfirmModal } from './confirm.modal';
import { FormModal } from './form.modal';
import { PaginationComponent } from './pagination.component';


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
        ConfirmModal,
        FormModal,
        PaginationComponent
    ],
    exports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        FormComponent,
        FieldComponent,
        TableComponent,
        PaginationComponent,
    ],
    entryComponents: [
        ConfirmModal,
        FormModal
    ]
})
export class FormModule { }
