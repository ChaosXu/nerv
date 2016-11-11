export class Field {
    name: string;
    label: string;
    control: string;
    type: string;
    required: boolean;
}

export class Textbox extends Field {
    control = 'textbox';
    type = 'string'
}

export class Form {
    name: string;
    fields: Field[];
}
