import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, Params, UrlSegment } from '@angular/router';

@Component({
    selector: 'nerv-code',
    templateUrl: 'app/lib/code/code.component.html',
})
export class CodeComponent implements OnInit {
    nodes = [
        {
            id: 1,
            name: 'root1',
            children: [
                { id: 2, name: 'child1' },
                { id: 3, name: 'child2' }
            ]
        },
        {
            id: 4,
            name: 'root2',
            children: [
                { id: 5, name: 'child2.1' },
                {
                    id: 6,
                    name: 'child2.2',
                    children: [
                        { id: 7, name: 'subsub' }
                    ]
                }
            ]
        }
    ];

    ngOnInit(): void {

    }

    addFile() {

    }

    removeFile(file: {}) {

    }

    renameFile(file: {}) {

    }

    saveFile(file: {}) {

    }

    onTextChanged(text:string) {

    }
}