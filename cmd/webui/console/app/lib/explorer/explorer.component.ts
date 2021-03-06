import { Component, OnInit, ViewChild } from '@angular/core';
import { Router, ActivatedRoute, Params, UrlSegment } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { TreeComponent, TreeNode, ITreeOptions } from 'angular2-tree-component';
import { AceDirective } from '../form/ace/ace.directive';
import { FormModal } from '../form/form.modal';
import { ConfirmModal } from '../form/confirm.modal';
import { File, FileService } from './file.service';

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
        }
    ]
};

const RENAME_FORM = {
    name: "explorer_rename_form",
    fields: [
        {
            name: "name", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        }
    ]
};

@Component({
    selector: 'nerv-explorer',
    templateUrl: 'app/lib/explorer/explorer.component.html',
    styleUrls:[
        'app/lib/explorer/explorer.component.css'
    ]
})
export class ExplorerComponent implements OnInit {
    options = {
        idField: 'url',
        displayField: 'title',
        getChildren: (node: TreeNode) => {
            let file = node.data;
            return this.fileService.get(this.type,file.url);
        }
    };

    nodes = [];
    private contentLoading = false;

    @ViewChild(TreeComponent)
    private tree: TreeComponent;

    @ViewChild(AceDirective)
    private ace: AceDirective;

    selectedNode: File

    private type:string;

    get canRemove(): boolean {
        return this.selectedNode != null;
    }

    get canRename(): boolean {
        return this.selectedNode != null;
    }

    get canAddFile(): boolean {
        return this.selectedNode != null;
    }

    get canSave(): boolean {
        return this.selectedNode && this.selectedNode.type == 'file' && this.selectedNode.dirty;
    }

    get mode(): string {
        return this.selectedNode && this.selectedNode.type == 'file' ? this.getExt(this.selectedNode.name) : 'json';
    }

    constructor(
        private modalService: NgbModal,
        private fileService: FileService,
        private route: ActivatedRoute
    ) {
    }

    ngOnInit(): void {
        this.route.parent.url.forEach((segment: UrlSegment[]) => {
                this.type = segment[0].path;                
            });                             
        this.fileService.get(this.type,'')
            .then((data) => {
                this.nodes = data;
            })
            .catch((error) => this.error('加载错误', `加载目录失败\r\n${error}`));
    }

    onAddFile() {
        let newItem = {};
        const modalRef = this.modalService.open(FormModal, { backdrop: 'static' });
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
        const modalRef = this.modalService.open(FormModal, { backdrop: 'static' });
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
        const modalRef = this.modalService.open(ConfirmModal, { backdrop: 'static' });
        modalRef.componentInstance.title = '删除';
        modalRef.componentInstance.message = `删除${this.selectedNode.name}?`;
        modalRef.result.then((result) => {
            if (result == 'ok') {
                this.remove(this.selectedNode);
            }
        });
    }



    onRename() {
        let newItem = {};
        const modalRef = this.modalService.open(FormModal, { backdrop: 'static' });
        modalRef.componentInstance.title = '重命名';
        modalRef.componentInstance.form = RENAME_FORM;
        modalRef.componentInstance.data = newItem;
        modalRef.result.then((result) => {
            if (result == 'ok') {
                this.rename(newItem['name']);
            }
        });
    }

    onSave() {
        const file = this.selectedNode;
        if (file && file.type == 'file') {
            this.saveFile(file);
        }
    }

    onTextChanged(text: string) {
        if (this.contentLoading) return;
        const file = this.selectedNode;
        if (file && file.type == 'file') {
            file.dirty = true;
            file.content = this.getAceContent();
        }
    }

    onNodeSelected(event: {}) {
        const node = event['node'];
        const file = this.selectedNode = node['data'];
        if (file && file.type == 'file') {
            if (file.dirty) {
                this.contentLoading = true;
                this.ace.text = file.content;
                this.contentLoading = false;
            } else {
                this.loadContent(file);
            }
        }
    }

    private saveFile(file: File): void {
        this.fileService.update(this.type,file.url, file)
            .then((response) => {                
                file.dirty = false;
                file.content = null;
                this.tree.treeModel.update();
            })
            .catch((error) => {
                this.contentLoading = false;
                this.error('保存错误', `保存对象失败\r\n${error}`)
            });    
    }

    private getExt(name: string): string {
        let index = name.lastIndexOf('.');
        if (index > 0) {
            return name.substr(index + 1);
        } else {
            return 'json';
        }
    }

    private getAceContent(): string {
        return this.ace.text;
    }

    private loadContent(file: File): void {
        if (file.dirty) {
            const content = file.content;
            this.ace.text = content || '';
        } else {
            this.contentLoading = true;
            this.fileService.get(this.type,file.url)
                .then((content) => {
                    this.ace.text = content;
                    this.contentLoading = false;
                })
                .catch((error) => {
                    this.contentLoading = false;
                    this.error('加载错误', `加载文件失败\r\n${error}`)
                });
        }
    }

    private rename(name: string): void {
        const node = this.selectedNode;
        let newUrl = node.url;
        let index = newUrl.lastIndexOf(node.name);
        newUrl = newUrl.substring(0,index) + name;
        let updateFile = { url: newUrl, name: name, type: node.type }
        this.fileService.update(this.type,node.url, updateFile)
            .then((response) => {                
                let children = node.children;
                this.updateChildrenUrl(children,newUrl);
                node.name = name;
                node.url = newUrl;
                this.tree.treeModel.update();
            })
            .catch((error) => {
                this.contentLoading = false;
                this.error('重命名错误', `重命名对象失败\r\n${error}`)
            });
    }

    private updateChildrenUrl(children:File[],parentUrl:string) {
        if (!children) return;
        for(let c of children) {
            c.url = parentUrl+"/"+c.name;
            this.updateChildrenUrl(c.children,c.url);            
        }
    }

    private remove(file: File): void {
        this.fileService.remove(this.type,file.url)
            .then((remvoed) => {
                const activeNode = this.tree.treeModel.activeNodes[0];
                const index = activeNode.parent.data.children.indexOf(activeNode.data);
                activeNode.parent.data.children.splice(index, 1);
                this.tree.treeModel.update();
                this.selectedNode = null;
                this.ace.text = '';
            })
            .catch((error) => {
                this.contentLoading = false;
                this.error('删除错误', `删除对象失败\r\n${error}`)
            });
    }

    private addFile(parent: File, node: {}): void {
        const selectedNode = this.selectedNode;
        const name = node['name'];
        let file;
        let children;
        if (selectedNode) {
            if (selectedNode.type == 'dir') {
                children = selectedNode.children;
                if (!children) {
                    this.selectedNode.children = children = new Array<File>();
                }
                file = { url: selectedNode.url+'/'+name, name: name, type: 'file', extension: this.getExt(name) };                
            } else {
                const parent = this.tree.treeModel.activeNodes[0].parent.data;
                children = parent.children;                
                file = { url: parent.url+'/'+name, name: name, type: 'file', extension: this.getExt(name) };                                
            }
        }

        this.fileService.create(this.type,file.url)
            .then((response) => {                                
                children.push(file);
                this.tree.treeModel.update();
            })
            .catch((error) => {
                this.contentLoading = false;
                this.error('保存错误', `保存对象失败\r\n${error}`)
            });
    }


    private addDir(parent: File, node: {}): void {
        const name = node['name'];
        const selectedNode = this.selectedNode;
        let nodes;
        let parentNode;
        if (selectedNode) {
            if (selectedNode['type'] == 'file') {
                parentNode = this.tree.treeModel.activeNodes[0].parent;
                nodes = parentNode.data.children;
            } else if (node['dirType'] == 'child') {
                parentNode = this.tree.treeModel.activeNodes[0];
                nodes = selectedNode.children;
                if (!nodes) {
                    selectedNode.children = nodes = new Array<File>();
                }
            } else {
                nodes = this.nodes;
            }
        } else {
            nodes = this.nodes;
        }
        let path = parentNode ? parentNode.data['url'] : '/';
        if (path == ('/')) {
            path = `${path}${name}`;
        } else {
            path = `${path}/${name}`;
        }
        this.fileService.create(this.type,path)
            .then((dir) => {
                nodes.push(dir);
                this.tree.treeModel.update();
            })
            .catch((error) => this.error('添加目录错误', `添加目录失败\r\n${error}`));


    }

    private error(title: string, error: any): void {
        const modalRef = this.modalService.open(ConfirmModal, { backdrop: 'static' });
        modalRef.componentInstance.title = title;
        modalRef.componentInstance.message = error.toString();
        modalRef.componentInstance.buttons = { ok: true, cancel: false };
    }
}