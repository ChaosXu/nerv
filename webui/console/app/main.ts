import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';
import { AppModule } from './application/module';
const platform = platformBrowserDynamic();
platform.bootstrapModule(AppModule);
