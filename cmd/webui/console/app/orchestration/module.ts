import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Form } from '../lib/form/model';
import { FormConfig } from '../lib/config/form.config';
import { FormRegistry } from '../lib/form/form.registry';

const stForm: Form = {
    name: "service_template_form",
    fields: [
        {
            name: "name", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "version", label: "版本", control: "number", type: "long", validators: { 'required': { message: '不能为空' } }
        }
    ]
};

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

const rtForm: Form = {
    name: "resource_type_form",
    fields: [
        {
            name: "name", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "version", label: "版本", control: "number", type: "long", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "operations", label: "操作", control: "table", type: "Operation",
            condition: 'resource_type_id=?',
            forms: {
                add: 'orchestration.rt.op.add',
                edit: 'orchestration.rt.op.edit',
                detail: 'orchestration.rt.op.detail',
            },
            //validators: { 'table_required': { message: '不能为空' } },
            display: {
                columns: [
                    { label: '名称', name: 'name' },
                    { label: '类型', name: 'type' },
                    { label: '实现', name: 'implementor' },
                ]
            },

        }
    ]
};

const opForm: Form = {
    name: "op_form",
    fields: [
        {
            name: "name", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "type", label: "类型", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "implementor", label: "实现", control: "text", type: "text", validators: { 'required': { message: '不能为空' } }
        }
    ]
};

@NgModule({
})
export class OrchestrationModule {

    constructor(
        formRegistry: FormRegistry,
        configService: FormConfig
    ) {
        formRegistry.put('orchestration.st.add', stForm);
        formRegistry.put('orchestration.st.edit', stForm);
        formRegistry.put('orchestration.st.detail', stForm);

        formRegistry.put('orchestration.topo.add', topoForm);
        formRegistry.put('orchestration.topo.edit', topoForm);
        formRegistry.put('orchestration.topo.detail', topoForm);

        formRegistry.put('orchestration.rt.add', rtForm);
        formRegistry.put('orchestration.rt.edit', rtForm);
        formRegistry.put('orchestration.rt.detail', rtForm);

        formRegistry.put('orchestration.rt.op.add', opForm);
        formRegistry.put('orchestration.rt.op.edit', opForm);
        formRegistry.put('orchestration.rt.op.detail', opForm);

        configService.put('orchestration',
            {
                title: '资源编排',
                menus: [
                    {
                        name: "模板管理",
                        url: "/orchestration/ServiceTemplate"
                    }, {
                        name: "拓扑管理",
                        url: "/orchestration/Topology"
                    }, {
                        name: "资源类",
                        url: "/orchestration/ResourceType"
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
                ResourceType: {
                    list: {
                        title: '资源类列表', columns: [
                            { label: '名称', name: 'name' },
                            { label: '操作' }
                        ]
                    },
                    add: { title: '添加资源类', form: 'orchestration.rt.add' },
                    edit: { title: '编辑资源类', form: 'orchestration.rt.edit' },
                    detail: { title: '查看资源类', form: 'orchestration.rt.detail' }
                }
            });
    }
}
