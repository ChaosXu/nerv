import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router';
import { FormModule } from '../form/module';
import { RestyModule } from '../resty/module';
import { LoginModal } from './login.modal';
import { AuthGuard } from './auth.guard';
import { LoginService } from './login.service';

@NgModule({
    imports: [
        BrowserModule,
        FormModule,
        RestyModule    
    ],   
    declarations:[
        LoginModal,
    ], 
    exports: [
        RouterModule,
    ],
    providers: [
        AuthGuard,
        LoginService
    ],
    entryComponents:[
        LoginModal
    ]    
})
export class SecurityModule { }
