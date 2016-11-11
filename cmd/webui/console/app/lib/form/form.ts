import { Component, Input, OnInit } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { Form } from './metadata';

@Component({
    selector: 'nerv-form',
    templateUrl: 'app/lib/form/form.html',
})
export class FormComponent implements OnInit {

    @Input() meta: Form;

    ngOnInit(): void {

    }
}