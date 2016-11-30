import { Component } from '@angular/core';
import { CatalogItem } from './service/catalog';
import { LoginService } from '../lib/security/login.service';

@Component({
  selector: 'nerv-console',
  templateUrl: 'app/main/application.html'
})
export class Application {
  docks: CatalogItem[] = [];
  userMenus = [
    { 'label': '退出', 'name': 'exit' }
  ];

  cmds = {
    'exit': {
      cmd: (function (self) {
        return function () {
          self.loginService.logout().then(() => window.location.href='/');
        }
      })(this)
    }
  };

  user = '';
  displayUserMenu = false;

  constructor(    
    private loginService: LoginService
  ) {
    this.loginService.loginSuccess.subscribe(function (target: Application) {
      return function (user: string) {
        target.user = user;
      }
    } (this));
  }

  toggleDock(item: CatalogItem): void {
    var index = this.docks.findIndex(e => e.name == item.name);
    if (index > -1) {
      this.docks.splice(index, 1);
    } else {
      this.docks.push(item);
    }
  }

  toggleUserMenu() {
    this.displayUserMenu = !this.displayUserMenu;
  }

  onUserMenu(name:string) {
    this.cmds[name].cmd();
  }
}