import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Form } from '../lib/form/model';
import { FormConfig } from '../lib/config/form.config';

const stForm: Form = {
    name: "service_template_form",
    fields: [
        {
            name: "Name", label: "名称", control: "text", type: "string", validators: {
                'required': '不能为空'
            }
        },
        {
            name: "Version", label: "版本", control: "text", type: "string", validators: {
                'required': '不能为空'
            }
        }
    ]
};

const topoForm: Form = {
    name: "topology_form",
    fields: [
        {
            name: "Name", label: "名称", control: "text", type: "string", validators: {
                'required': '不能为空'
            }
        },
        {
            name: "Template", label: "模板", control: "text", type: "string", validators: {
                'required': '不能为空'
            }
        },
        {
            name: "Version", label: "版本", control: "text", type: "string", validators: {
                'required': '不能为空'
            }
        }
    ]
};

@NgModule({
})
export class OrchestrationModule {

    constructor(configService: FormConfig) {
        configService.put('orchestration',
            {
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
                            { label: '名称', name: 'Name' },
                            { label: '版本', name: 'Version' },
                            { label: '操作' }
                        ]
                    },
                    add: { title: '添加模板', form: stForm },
                    edit: { title: '编辑模板', form: stForm },
                    detail: { title: '查看模板', form: stForm }
                },
                Topology: {
                    list: {
                        title: '拓扑列表', columns: [
                            { label: '名称', name: 'Name' },
                            { label: '模板', name: 'Template' },
                            { label: '版本', name: 'Version' },
                            { label: '操作' }
                        ]
                    },
                    add: { title: '添加拓扑', form: topoForm },
                    edit: { title: '编辑拓扑', form: topoForm },
                    detail: { title: '查看拓扑', form: topoForm }
                },
                ResourceType: {
                    list: {
                        title: '资源类列表', columns: [
                            { label: '名称', name: 'Name' },                                                        
                            { label: '操作' }
                        ]
                    },
                    add: { title: '添加资源类', form: topoForm },
                    edit: { title: '编辑资源类', form: topoForm },
                    detail: { title: '查看资源类', form: topoForm }
                }
            });
    }
}
