import { Form } from './model';

export class FormRegistry {
    forms = Form;

    put(name: string, form: Form): void {
        this.forms[name] = form;
    }

    get(name: string): Form {
        return this.forms[name];
    }
}