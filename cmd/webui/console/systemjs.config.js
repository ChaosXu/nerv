/**
 * System configuration for Angular samples
 * Adjust as necessary for your application needs.
 */
(function (global) {
  System.config({
    paths: {
      // paths serve as alias
      'npm:': 'node_modules/'
    },
    // map tells the System loader where to look for things
    map: {
      // our app is within the app folder
      app: 'js',
      // angular bundles
      '@angular/core': 'npm:@angular/core/bundles/core.umd.js',
      '@angular/common': 'npm:@angular/common/bundles/common.umd.js',
      '@angular/compiler': 'npm:@angular/compiler/bundles/compiler.umd.js',
      '@angular/platform-browser': 'npm:@angular/platform-browser/bundles/platform-browser.umd.js',
      '@angular/platform-browser-dynamic': 'npm:@angular/platform-browser-dynamic/bundles/platform-browser-dynamic.umd.js',
      '@angular/http': 'npm:@angular/http/bundles/http.umd.js',
      '@angular/router': 'npm:@angular/router/bundles/router.umd.js',
      '@angular/forms': 'npm:@angular/forms/bundles/forms.umd.js',
      '@angular/upgrade': 'npm:@angular/upgrade/bundles/upgrade.umd.js',
      // other libraries
      'rxjs': 'npm:rxjs',
      '@ng-bootstrap/ng-bootstrap': 'npm:@ng-bootstrap/ng-bootstrap/bundles/ng-bootstrap.js',      
      'angular2-tree-component':'npm:angular2-tree-component',
      'lodash':'npm:lodash',
      'brace':'npm:brace',
      'w3c-blob':'npm:w3c-blob',
      'buffer':'npm:buffer-shims'   
    },
    // packages tells the System loader how to load when no filename and/or no extension
    packages: {
      app: {
        main: './main.js',
        defaultExtension: 'js'
      },
      rxjs: {
        defaultExtension: 'js'
      },
      'angular2-tree-component': {
        main: 'dist/angular2-tree-component.js',
        defaultExtension: 'js'
      },
      lodash: {
        main: 'lodash.js',
        defaultExtension: 'js'
      },
      brace: {
        main:'index.js',
        defaultExtension:'js'
      }, 
      'w3c-blob':{
        main:'index.js',
        defaultExtension:'js'
      },
      'buffer':{
        main:'index.js',
        defaultExtension:'js'
      }   
    }
  });
})(this);
