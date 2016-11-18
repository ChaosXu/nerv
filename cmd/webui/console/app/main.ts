import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import { MainModule } from './main/module';
const platform = platformBrowserDynamic();
platform.bootstrapModule(MainModule);
