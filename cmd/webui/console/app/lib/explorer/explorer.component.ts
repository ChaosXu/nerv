import { Component, OnInit, ViewChild } from '@angular/core';
import { Router, ActivatedRoute, Params, UrlSegment } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { FormModal } from '../form/form.modal';
import { ConfirmModal } from '../form/confirm.modal';
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

    nodes = [];

    @ViewChild(TreeComponent)
    private tree: TreeComponent;

    @ViewChild(AceDirective)
    private ace: AceDirective;

    get selectedNode(): File {
        const nodes = this.tree.treeModel.activeNodes;
        if (nodes && nodes.length > 0)
            return nodes[0].data;
        else
            return null;

    }
    get canRemove(): boolean {
        return this.selectedNode != null;
    }

    get canRename(): boolean {
        return this.selectedNode != null;
    }

    get canAddFile(): boolean {
        return this.selectedNode != null;
    }

    // get canAddDir(): boolean {
    //     return !this.selectedNode || this.selectedNode.type == 'dir';
    // }

    get canSave(): boolean {
        return this.selectedNode && this.selectedNode.type == 'file';
    }

    get mode(): string {
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

    onRemove() {
        if (!this.canRemove) return;
        const modalRef = this.modalService.open(ConfirmModal);
        modalRef.componentInstance.title = '删除';
        modalRef.componentInstance.message = `删除${this.selectedNode.title}?`;
        modalRef.result.then((result) => {
            if (result == 'ok') {
                this.remove(this.selectedNode);
            }
        });
    }



    renameFile(file: {}) {

    }

    saveFile(file: {}) {

    }

    onTextChanged(text: string) {
        if (this.selectedNode && this.selectedNode.type == 'file') {
            this.selectedNode.content = text;
        }
    }

    onActive(event: {}) {
        const selectedNode = this.selectedNode;
        if (selectedNode && selectedNode.type == 'file') {
            this.ace.text = selectedNode.content || '';
        }
    }

    private remove(file: File): void {
        const activeNode = this.tree.treeModel.activeNodes[0];
        const index = activeNode.parent.data.children.indexOf(activeNode.data);
        activeNode.parent.data.children.splice(index, 1);
        this.tree.treeModel.update();
        this.ace.text = '';
    }
    
    private addFile(parent: File, node: {}): void {
        const selectedNode = this.selectedNode;
        if (selectedNode) {
            if (selectedNode.type == 'dir') {
                const name = node['name'];
                let children = selectedNode.children;
                if (!children) {
                    this.selectedNode.children = children = new Array<File>();
                }
                children.push({ name: name, type: 'file', title: `${name}.${node['fileType']}`, extension: node['fileType'] });
                this.tree.treeModel.update();
        //        this.tree.treeModel.activeNodes[0].children[children.length-1].focus();
            } else {
                let children = this.tree.treeModel.activeNodes[0].parent.data.children;
                children.push({ name: name, type: 'file', title: `${name}.${node['fileType']}`, extension: node['fileType'] });
                this.tree.treeModel.update();
      //          this.tree.treeModel.activeNodes[0].parent.children[children.length-1].focus();
            }

            
        }
    }

    private addDir(parent: File, node: {}): void {
        const name = node['name'];
        const selectedNode = this.selectedNode;
        if (selectedNode) {
            if (node['dirType'] == 'child') {
                let children = selectedNode.children;
                if (!children) {
                    selectedNode.children = children = new Array<File>();
                }
                children.push({ name: name, type: 'dir', title: name });
                this.tree.treeModel.update();
    //            this.tree.treeModel.activeNodes[0].children[children.length-1].focus();
            } else {
                this.nodes.push({ name: name, type: 'dir', title: name });
                this.tree.treeModel.update();
  //              this.tree.treeModel.nodes[this.tree.treeModel.nodes.length-1].focus();
            }
        } else {
            this.nodes.push({ name: name, type: 'dir', title: name });
            this.tree.treeModel.update();
//this.tree.treeModel.nodes[this.tree.treeModel.nodes.length-1].focus();
        }
        
    }
}