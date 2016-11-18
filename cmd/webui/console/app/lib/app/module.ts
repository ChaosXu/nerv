import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { RestyModule } from '../resty/module';
import { ConfigModule } from '../config/module';
import { FormModule } from '../form/module';
import { ListApp } from './list.app';
import { ListComponent } from './list.form';
import { AddComponent } from './add.form';
import { EditComponent } from './edit.form';
import { DetailComponent } from './detail.form';
import { routes } from './routes';

@NgModule({
    imports: [
        RestyModule,
        FormModule,
        RouterModule.forChild(routes)
    ],
    declarations: [
        ListApp,
        ListComponent,
        AddComponent,
        EditComponent,
        DetailComponent
    ]    
})
export class AppModule {
}
