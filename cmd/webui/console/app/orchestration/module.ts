import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Form } from '../lib/form/model';
import { FormConfig } from '../lib/config/form.config';
import { FormRegistry } from '../lib/form/form.registry';
import { AppModule } from '../lib/app/module';
import { routes } from './routes';
import { ExplorerModule } from '../lib/explorer/module';

// const stForm: Form = {
//     name: "service_template_form",
//     fields: [
//         {
//             name: "name", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
//         },
//         {
//             name: "version", label: "版本", control: "number", type: "long", validators: { 'required': { message: '不能为空' } }
//         }
//     ]
// };

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

// const rtForm: Form = {
//     name: "resource_type_form",
//     fields: [
//         {
//             name: "name", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
//         },
//         {
//             name: "version", label: "版本", control: "number", type: "long", validators: { 'required': { message: '不能为空' } }
//         },
//         {
//             name: "operations", label: "操作", control: "table", type: "Operation",
//             condition: 'resource_type_id=?',
//             forms: {
//                 add: 'orchestration.rt.op.add',
//                 edit: 'orchestration.rt.op.edit',
//                 detail: 'orchestration.rt.op.detail',
//             },
//             //validators: { 'table_required': { message: '不能为空' } },
//             display: {
//                 columns: [
//                     { label: '名称', name: 'name' },
//                     { label: '类型', name: 'type' },
//                     { label: '实现', name: 'implementor' },
//                 ]
//             },

//         }
//     ]
// };

// const opForm: Form = {
//     name: "op_form",
//     fields: [
//         {
//             name: "name", label: "名称", control: "select", type: "string", validators: { 'required': { message: '不能为空' } },
//             display: {
//                 options: [
//                     { label: 'Create', value: 'Create', default: true },
//                     { label: 'Delete', value: 'Delete' },
//                     { label: 'Setup', value: 'Setup' },
//                     { label: 'Start', value: 'Start' },
//                     { label: 'Stop', value: 'Stop' },
//                 ]
//             }
//         },
//         {
//             name: "type", label: "类型", control: "select", type: "string", validators: { 'required': { message: '不能为空' } },
//             display: {
//                 options: [
//                     { label: 'Shell', value: 'shell', default: true },
//                     { label: 'GO', value: 'go' }
//                 ]
//             }
//         },
//         {
//             name: "implementor", label: "实现", control: "text", type: "text", validators: { 'required': { message: '不能为空' } }
//         }
//     ]
// };

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
                        url: "/orchestration/ServiceTemplate"
                    },{
                        name: "资源模型",
                        url: "/orchestration/Script"
                    }
                ],
                ServiceTemplate: {
                    list: {
                        title: '模板列表', columns: [
                            { label: '名称', name: 'name' },
                            { label: '版本', name: 'version' },
                            { label: '操作' }
                        ]
                    },
                    add: { title: '添加模板', form: 'orchestration.st.add' },
                    edit: { title: '编辑模板', form: 'orchestration.st.edit' },
                    detail: { title: '查看模板', form: 'orchestration.st.detail' }
                },
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
                },
                // ResourceType: {
                //     list: {
                //         title: '资源类列表', columns: [
                //             { label: '名称', name: 'name' },
                //             { label: '操作' }
                //         ]
                //     },
                //     add: { title: '添加资源类', form: 'orchestration.rt.add' },
                //     edit: { title: '编辑资源类', form: 'orchestration.rt.edit' },
                //     detail: { title: '查看资源类', form: 'orchestration.rt.detail' }
                // },                
                Script: {
                    list: {
                        title: '脚本列表', columns: [
                            { label: '名称', name: 'name' },
                            { label: '操作' }
                        ]
                    },
                    add: { title: '添加脚本', form: 'orchestration.script.add' },
                    edit: { title: '编辑脚本', form: 'orchestration.script.edit' },
                    detail: { title: '查看删除', form: 'orchestration.script.detail' }
                }
            });
    }
}
