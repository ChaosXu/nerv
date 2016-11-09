import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { Application } from './ui/application';
import { StartMenu } from './ui/startmenu';
import { Dock } from './ui/dock';
import { CatalogService } from './service/catalog';

@NgModule({
  imports: [
    BrowserModule//ßß,
    // RouterModule.forRoot(
    //   {path:'company',component:}
    // )
  ],
  declarations: [
    Application,
    StartMenu,
    Dock
  ],
  providers: [CatalogService],
  bootstrap: [Application]
})
export class AppModule { }
