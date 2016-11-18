//app config
export class ConfigService {
    formConfig = {};

    get(name: string): any {
        return this.formConfig[name];
    }

    put(name: string, config: any): void {
        this.formConfig[name] = config;
    }
}