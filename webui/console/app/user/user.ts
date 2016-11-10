import { Component } from '@angular/core';
import { OnInit } from '@angular/core';

interface Menu {
    name:string;
    url:string;
}

@Component({
    selector: 'nerv-app-user',
    templateUrl: 'app/user/user.html', 
})
export class UserApp {
    menus: Menu[];

    constructor(){
        this.menus=[{
            name:"项目管理",
            url:"/user/project"
        },{
            name:"人员管理",
            url:"/user/account"
        },{
            name:"权限管理",
            url:"/user/role"
        }];
    }
}