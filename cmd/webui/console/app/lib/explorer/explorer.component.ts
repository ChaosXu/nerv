import { Component, OnInit, ViewChild } from '@angular/core';
import { Router, ActivatedRoute, Params, UrlSegment } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { FormModal } from '../form/form.modal';
import { File } from './file.service';
import { TreeComponent } from 'angular2-tree-component';
import { AceDirective } from '../form/ace/ace.directive';

const ADD_DIR_FORM = {
    name: "explorer_add_form",
    fields: [
        {
            name: "name", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "dirType", label: "类型", control: "select", type: "string",
            display: {
                options: [
                    { label: '子文件夹', value: 'child', default: true },
                    { label: '根文件夹', value: 'root' }
                ]
            }
        }
    ]
};

const ADD_FILE_FORM = {
    name: "explorer_add_form",
    fields: [
        {
            name: "name", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "fileType", label: "类型", control: "select", type: "string",
            display: {
                options: [
                    { label: 'JSON', value: 'json', default: true },
                    { label: 'Shell', value: 'sh' },
                    { label: 'Ruby', value: 'ruby' },
                    { label: 'Python', value: 'python' }
                ]
            }
        }
    ]
};

@Component({
    selector: 'nerv-explorer',
    templateUrl: 'app/lib/explorer/explorer.component.html',
})
export class ExplorerComponent implements OnInit {
    options: { idField: 'name', displayField: 'title' };
    selectedNode: File;
    nodes = [];

    @ViewChild(TreeComponent)
    private tree: TreeComponent;

    @ViewChild(AceDirective)
    private ace:AceDirective;

    get canRemove(): boolean {
        return this.selectedNode != null;
    }

    get canRename(): boolean {
        return this.selectedNode != null;
    }

    get canAddFile(): boolean {
        return this.selectedNode && this.selectedNode.type == 'dir';
    }

    get canAddDir(): boolean {
        return !this.selectedNode || this.selectedNode.type == 'dir';
    }

    get canSave(): boolean {
        return this.selectedNode && this.selectedNode.type == 'file';
    }

    get mode():string {
        return this.selectedNode ? this.selectedNode.extension : 'json';
    }

    constructor(private modalService: NgbModal) { }
    ngOnInit(): void {

    }

    onAddFile() {

        let newItem = {};
        const modalRef = this.modalService.open(FormModal);
        modalRef.componentInstance.title = '添加文件';
        modalRef.componentInstance.form = ADD_FILE_FORM;
        modalRef.componentInstance.data = newItem;
        modalRef.result.then((result) => {
            if (result == 'ok') {
                this.addFile(this.selectedNode, newItem);
            }
        });
    }

    onAddDir() {
        let newItem = {};
        const modalRef = this.modalService.open(FormModal);
        modalRef.componentInstance.title = '添加文件夹';
        modalRef.componentInstance.form = ADD_DIR_FORM;
        modalRef.componentInstance.data = newItem;
        modalRef.result.then((result) => {
            if (result == 'ok') {
                this.addDir(this.selectedNode, newItem);
            }
        });
    }

    removeFile(file: {}) {

    }

    renameFile(file: {}) {

    }

    saveFile(file: {}) {

    }

    onTextChanged(text: string) {
        if (this.selectedNode && this.selectedNode.type=='file'){
            this.selectedNode.content = text;
        }        
    }

    onActive(event: {}) {
        this.selectedNode = event['node']['data'];
        const selectedNode =this.selectedNode;
        if (selectedNode && selectedNode.type=='file'){
            this.ace.text = selectedNode.content || '';
        }         
    }

    private addFile(parent: File, node: {}): void {
        if (!this.canAddFile) return;
        if (this.selectedNode) {
            const name = node['name'];
            let children = this.selectedNode.children;
            if (!children) {
                this.selectedNode.children = children = new Array<File>();
            }

            children.push({ name: name, type: 'file', title: `${name}.${node['fileType']}`, extension: node['fileType'] });
            this.tree.treeModel.update();
        }
    }

    private addDir(parent: File, node: {}): void {
        if (!this.canAddDir) return;
        const name = node['name'];
        if (this.selectedNode) {
            if (node['dirType'] == 'child') {
                let children = this.selectedNode.children;
                if (!children) {
                    this.selectedNode.children = children = new Array<File>();
                }
                children.push({ name: name, type: 'dir', title: name });
            } else {
                this.nodes.push({ name: name, type: 'dir', title: name });
            }
        } else {
            this.nodes.push({ name: name, type: 'dir', title: name });
        }
        this.tree.treeModel.update();
    }
}