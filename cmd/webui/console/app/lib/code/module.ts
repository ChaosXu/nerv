import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TreeModule } from 'angular2-tree-component';
import { RestyModule } from '../resty/module';
import { ConfigModule } from '../config/module';
import { FormModule } from '../form/module';
import { CodeComponent } from './code.component';

@NgModule({
    imports: [
        TreeModule,
        RestyModule,
        FormModule,
    ],
    declarations: [
        CodeComponent
    ],
    exports: [
        CodeComponent
    ]
})
export class CodeModule {
}
