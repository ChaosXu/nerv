import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { Application } from './ui/application';
import { StartMenu } from './ui/startmenu';
import { Dock } from './ui/dock';

@NgModule({
  imports: [BrowserModule],
  declarations: [
    Application,
    StartMenu,
    Dock
  ],
  bootstrap: [Application]
})
export class AppModule { }
