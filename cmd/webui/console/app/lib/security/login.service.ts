import { Injectable, EventEmitter } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { LoginModal } from './login.modal';

@Injectable()
export class LoginService {

    loginSuccess: EventEmitter<any> = new EventEmitter();

    private currentUser = '';

    constructor(private modalService: NgbModal) {

    }

    isLogin(): boolean {
        return this.currentUser != '';
    }

    login(): Promise<string> {
        
        const modalRef = this.modalService.open(LoginModal, { backdrop: 'static' });

        return modalRef.result.then((result) => {
            if (result) {
                this.currentUser = result['Name'];
                this.loginSuccess.emit(this.currentUser);
                return result;
            }
        });
    }

    logout() {

    }
}