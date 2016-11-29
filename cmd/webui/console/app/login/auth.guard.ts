import { Injectable } from '@angular/core';
import {
  CanActivate, Router,
  ActivatedRouteSnapshot,
  RouterStateSnapshot
} from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { LoginModal } from './login.modal';

@Injectable()
export class AuthGuard implements CanActivate {

  private isLogin: boolean = false;

  constructor(
    private router: Router,
    private modalService: NgbModal
  ) { }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
    if (this.isLogin) {
      return true;
    } else {
      this.login(state.url);
    }
  }

  private login(url: string) {
    let data = {};
    const modalRef = this.modalService.open(LoginModal, { backdrop: 'static' });
    
    modalRef.result.then((result) => {
      if (result == 'ok') {
        this.isLogin = true;
        this.router.navigate([url]);
      }
    });
  }
}