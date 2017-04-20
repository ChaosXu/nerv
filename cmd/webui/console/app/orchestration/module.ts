import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Form } from '../lib/form/model';
import { FormConfig } from '../lib/config/form.config';
import { FormRegistry } from '../lib/form/form.registry';
import { AppModule } from '../lib/app/module';
import { routes } from './routes';
import { ExplorerModule } from '../lib/explorer/module';

const topoForm: Form = {
    name: "topology_form",
    fields: [
        {
            name: "name", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "template", label: "模板", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "version", label: "版本", control: "number", type: "long", validators: { 'required': { message: '不能为空' } }
        }
    ]
};

const scriptForm: Form = {
    name: "script_form",
    fields: [
        {
            name: "name", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "type", label: "类型", control: "select", type: "string", validators: { 'required': { message: '不能为空' } },
            display: {
                options: [
                    { label: 'Shell', value: 'shell', default: true },
                    { label: 'GO', value: 'go' }
                ]
            }
        },
        {
            name: "path", label: "路径", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "file", label: "实现", control: "file", type: "file", validators: { 'required': { message: '不能为空' } }
        }
    ]
};

@NgModule({
    imports:[        
        AppModule,
        ExplorerModule,
        RouterModule.forChild(routes)        
    ]    
})
export class OrchestrationModule {

    constructor(
        formRegistry: FormRegistry,
        configService: FormConfig
    ) {
        
        formRegistry.put('orchestration.topo.add', topoForm);
        formRegistry.put('orchestration.topo.edit', topoForm);
        formRegistry.put('orchestration.topo.detail', topoForm);

        configService.put('orchestration',
            {
                title: '资源编排',
                menus: [
                    {
                        name: "拓扑管理",
                        url: "/orchestration/Topology"
                    }, {
                        name: "模板管理",
                        url: "/orchestration/templates"
                    },{
                        name: "资源模型",
                        url: "/orchestration/scripts"
                    }
                ],               
                Topology: {
                    list: {
                        title: '拓扑列表', columns: [
                            { label: '名称', name: 'name' },
                            { label: '模板', name: 'template' },
                            { label: '版本', name: 'version' },
                            { label: '操作' }
                        ]
                    },
                    add: { title: '添加拓扑', form: 'orchestration.topo.add' },
                    edit: { title: '编辑拓扑', form: 'orchestration.topo.edit' },
                    detail: { title: '查看拓扑', form: 'orchestration.topo.detail' }
                }
            });
    }
}
