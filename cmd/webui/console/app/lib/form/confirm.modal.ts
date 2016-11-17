import { Component, Input } from '@angular/core';

import { NgbModal, NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'nerv-modal-confirm',
  templateUrl: 'app/lib/form/confirm.modal.html'
})
export class ModalConfirm {
  @Input() title;
  @Input() message;
  @Input() buttons: {
    ok: boolean,
    cancel: boolean
  } = {
    ok: true,
    cancel: true
  };

  constructor(public activeModal: NgbActiveModal) { }
}