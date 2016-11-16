import { Injectable } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

export class Form {
    name: string;
    fields: Field[];
}

export class Field {
    name: string;
    label: string;
    control: string;
    type: string;
    validators?: {};
}

export class Textbox extends Field {
    control = 'textbox';
    type = 'string'
}


