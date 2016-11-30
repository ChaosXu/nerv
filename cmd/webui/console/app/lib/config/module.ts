import { NgModule } from '@angular/core';
import { FormConfig } from './form.config';
import { CatalogConfig } from './catalog.config';

@NgModule({
    providers:[
        FormConfig,
        CatalogConfig
    ]
})
export class ConfigModule {

}