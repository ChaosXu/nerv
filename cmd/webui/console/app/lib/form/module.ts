import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { TableComponent } from './table.component';
import { FormComponent } from './form.component';
import { PaginationComponent } from './pagination.component';
import { ConfirmModal } from './confirm.modal';
import { FormModal } from './form.modal';
import { InputField } from './input.field';
import { TableField } from './table.field';
import { FormRegistry } from './form.registry';

@NgModule({
    imports: [
        CommonModule,
        ReactiveFormsModule,
        NgbModule,
        FormsModule
    ],
    declarations: [
        ConfirmModal,
        FormModal,
        FormComponent,
        TableComponent,
        PaginationComponent,
        InputField,
        TableField
    ],
    exports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        FormComponent,
        TableComponent,
        PaginationComponent,
        InputField,
        TableField
    ],
    providers: [
        FormRegistry
    ],
    entryComponents: [
        ConfirmModal,
        FormModal
    ]
})
export class FormModule { }
