import { Component, Input } from '@angular/core';

import { NgbModal, NgbActiveModal } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'nerv-modal-confirm',
  templateUrl: 'app/lib/form/form.modal.html'
})
export class FormModal {
  @Input('readonly') enableReadonly = false;
  @Input() title;
  @Input() form;
  @Input() data;
  @Input() buttons: {
    ok: boolean,
    cancel: boolean
  } = {
    ok: true,
    cancel: true
  };

  constructor(public activeModal: NgbActiveModal) { }

  onOk() {
    this.activeModal.close('ok');
  }

  onCancel() {
    this.activeModal.close('cancel');
  }
}