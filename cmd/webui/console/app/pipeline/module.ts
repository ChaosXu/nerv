import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { Form } from '../lib/form/model';
import { FormConfig } from '../lib/config/form.config';
import { FormRegistry } from '../lib/form/form.registry';
import { AppModule } from '../lib/app/module';
import { routes } from './routes';
import { ExplorerModule } from '../lib/explorer/module';

const projectForm: Form = {
    name: "project_form",
    fields: [
        {
            name: "name", label: "编号", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "displayName", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "pipelinefile", label: "部署流水线", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "apps", label: "应用", control: "table", type: "App",
            display: {
                columns: [
                    { label: '标识', name: 'name' },
                    { label: '名称', name: 'displayName' },
                    { label: 'SCM', name: 'scm' }
                ]
            },
            forms: {
                add: 'pipeline.app.add',
                edit: 'pipeline.app.edit',
                detail: 'pipeline.app.detail'
            },
            condition: 'project_id = ?'
        }
    ]
};

const appForm: Form = {
    name: "app_form",
    fields: [
        {
            name: "name", label: "标识", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "displayName", label: "名称", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
        },
        {
            name: "scm", label: "SCM", control: "text", type: "string", validators: { 'required': { message: '不能为空' } }
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
    imports: [
        AppModule,
        ExplorerModule,
        RouterModule.forChild(routes)
    ]
})
export class PipelineModule {

    constructor(
        formRegistry: FormRegistry,
        configService: FormConfig
    ) {

        formRegistry.put('pipeline.project.add', projectForm);
        formRegistry.put('pipeline.project.edit', projectForm);
        formRegistry.put('pipeline.project.detail', projectForm);
        formRegistry.put('pipeline.app.add', appForm);
        formRegistry.put('pipeline.app.edit', appForm);
        formRegistry.put('pipeline.app.detail', appForm);

        configService.put('pipeline',
            {
                title: '部署流水线',
                menus: [
                    {
                        name: "项目管理",
                        url: "/pipeline/Project"
                    }, {
                        name: "模板管理",
                        url: "/pipeline/blueprint"
                    }
                ],
                Project: {
                    list: {
                        title: '项目列表', columns: [
                            { label: '标识', name: 'name' },
                            { label: '名称', name: 'displayName' },
                            { label: '操作' }
                        ]
                    },
                    add: { title: '添加项目', form: 'pipeline.project.add' },
                    edit: { title: '编辑项目', form: 'pipeline.project.edit' },
                    detail: { title: '查看项目', form: 'pipeline.project.detail' }
                }
            });
    }
}
