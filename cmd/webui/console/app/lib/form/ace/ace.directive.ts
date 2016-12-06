import { Directive, Input, Output, ElementRef, EventEmitter, OnInit } from '@angular/core';
import * as ace from 'brace';
import 'brace/mode/javascript';
import 'brace/mode/json';
import 'brace/mode/sh';
import 'brace/mode/ruby';
import 'brace/mode/python';
import 'brace/theme/monokai';

@Directive({
    selector: 'nerv-ace'
})
export class AceDirective {
    @Input() set text(value) {
        if (value === this.oldVal) return;
        this.editor.setValue(value);
        this.editor.clearSelection();
        this.editor.focus();
    }
    @Input() set mode(value) {
        this._mode = value;
        this.editor.getSession().setMode(`ace/mode/${value}`);
    }
    @Input() set options(value) {
        this.editor.setOptions(value || {});
    }

    @Input() set readOnly(value) {
        this._readOnly = value;
        this.editor.setReadOnly(value);
    }

    @Input() set theme(value) {
        this._theme = value;
        this.editor.setTheme(`ace/theme/${value}`);
    }

    @Output() textChanged = new EventEmitter();
    @Output() editorRef = new EventEmitter();
    
    private editor: ace.Editor;
    private oldVal: string;
    private newVal: string;
    private _mode: string;
    private _readOnly: boolean;
    private _theme: string;

    constructor(private elementRef: ElementRef) {
        const el = this.elementRef.nativeElement;
        el.classList.add('editor');

        this.editor = ace.edit(el);

        setTimeout(() => {
            this.editorRef.next(this.editor);
        });

        this.editor.on('change', () => {
            const newVal = this.editor.getValue();
            if (newVal === this.oldVal) return;
            if (this.oldVal) {
                this.textChanged.next(newVal);
            }
            this.oldVal = newVal;
        });
    }
}