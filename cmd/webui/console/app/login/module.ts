import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router';
import { FormModule } from '../lib/form/module';
import { RestyModule } from '../lib/resty/module';
import { LoginModal } from './login.modal';
import { AuthGuard } from './auth.guard';

@NgModule({
    imports: [
        BrowserModule,
        FormModule,
        RestyModule    
    ],   
    declarations:[
        LoginModal
    ], 
    exports: [
        RouterModule,
    ],
    providers: [
        AuthGuard
    ],
    entryComponents:[
        LoginModal
    ]    
})
export class LoginModule { }
