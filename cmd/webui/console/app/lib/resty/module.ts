import { NgModule } from '@angular/core';
import { HttpModule } from '@angular/http';
import { RestyService } from './resty.service';


@NgModule({
    imports: [
        HttpModule
    ],
    providers: [RestyService]
})
export class RestyModule { }
