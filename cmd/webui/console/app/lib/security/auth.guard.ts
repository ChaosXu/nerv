import { Injectable } from '@angular/core';
import {
  CanActivate, Router,
  ActivatedRouteSnapshot,
  RouterStateSnapshot
} from '@angular/router';
import { LoginService } from './login.service';

@Injectable()
export class AuthGuard implements CanActivate {

  private isLogin: boolean = false;

  constructor(
    private router: Router,
    private loginService:LoginService
  ) { }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
    if (this.loginService.isLogin()) {
      return true;
    } else {
      this.login(state.url);
    }
  }

  private login(url: string) {
    
    this.loginService.login().then((result) => {
      if (result == 'ok') {
        this.isLogin = true;
        this.router.navigate([url]);
      }
    });
  }
}