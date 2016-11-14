import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormModule } from './form/module';


@NgModule({
    imports: [
        CommonModule,        
        RouterModule,        
        FormModule,
    ],
    exports: [
        CommonModule,
        FormModule,
        RouterModule
    ]
})
export class LibModule { }
