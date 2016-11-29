import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule, Routes } from '@angular/router';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { ConfigModule } from '../lib/config/module';
import { Application } from './application';
import { StartMenu } from './startmenu';
import { Dock } from './dock';
import { CatalogService } from './service/catalog';
import { routes } from './routes';
import { UserModule } from '../user/module';
import { LoginModule } from '../login/module';

@NgModule({
  imports: [
    BrowserModule,
    RouterModule.forRoot(routes),
    NgbModule.forRoot(),
    ConfigModule,
    UserModule,
    LoginModule
  ],
  declarations: [
    Application,
    StartMenu,
    Dock
  ],
  providers: [
    CatalogService    
  ],
  bootstrap: [Application]
})
export class MainModule { }
