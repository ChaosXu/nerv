import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormModule } from './form/module';
import { DbModule } from './resty/module';

@NgModule({
    imports: [
        CommonModule,
        RouterModule,
        FormModule,
        DbModule
    ],
    exports: [
        CommonModule,
        FormModule,
        RouterModule,
        DbModule
    ]
})
export class LibModule { }
