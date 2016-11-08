import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { Console } from './ui/console/console';
import { StartMenu } from './ui/startmenu/startmenu';

@NgModule({
  imports: [BrowserModule],
  declarations: [
    Console,
    StartMenu
  ],
  bootstrap: [Console]
})
export class AppModule { }
