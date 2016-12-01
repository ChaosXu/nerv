export class Form {
    name: string;
    fields: Field[];
}

export class Field {
    name: string;           //data property name
    label: string;          //display name
    control: string;        //input control
    type: string;           //data type
    validators?: {};
    display?: {};           //display setting 
    condition?: string      //condition for select table
    forms?: {
        add: string,
        edit: string,
        detail: string
    }
}
