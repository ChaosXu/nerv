import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormsModule } from '@angular/forms';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { TableComponent } from './table.component';
import { FormComponent } from './form.component';
import { PaginationComponent } from './pagination.component';
import { ConfirmModal } from './confirm.modal';
import { FormModal } from './form.modal';
import { InputField } from './fields/input.field';
import { TableField } from './fields/table.field';
import { SelectField } from './fields/select.field';
import { FormRegistry } from './form.registry';
import { ValidatorRegistry } from './validators/validator.registry';
import { AceDirective } from './ace/ace.directive';

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
        AceDirective,
        TableComponent,
        PaginationComponent,
        InputField,
        TableField,
        SelectField
    ],
    exports: [
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        AceDirective,
        FormComponent,
        TableComponent,
        PaginationComponent,
        InputField,
        TableField,
        SelectField
    ],
    providers: [
        FormRegistry,
        ValidatorRegistry
    ],
    entryComponents: [
        ConfirmModal,
        FormModal
    ]
})
export class FormModule { }
