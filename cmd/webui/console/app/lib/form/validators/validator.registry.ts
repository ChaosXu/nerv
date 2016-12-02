import { ValidatorFn, Validators } from '@angular/forms';


export class ValidatorRegistry {
    private _validators = {};

    constructor() {
        this.put('required', Validators.required);
    }

    put(name: string, validator: ValidatorFn): void {
        this._validators[name] = validator;
    }

    get(name: string): ValidatorFn {
        return this._validators[name];
    }
}