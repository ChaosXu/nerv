import { Injectable } from '@angular/core';

export interface Account {
    name: string;
    nick: string;
    mail: string;
    phone: string;

}

var accounts = [
    {
        name: "xj",
        nick: "chaos",
        mail: "chaosxu@foxmail.com",
        phone: "12345678901"
    },
    {
        name: "xj",
        nick: "chaos",
        mail: "chaosxu@foxmail.com",
        phone: "12345678901"
    },
    {
        name: "xj",
        nick: "chaos",
        mail: "chaosxu@foxmail.com",
        phone: "12345678901"
    },
    {
        name: "xj",
        nick: "chaos",
        mail: "chaosxu@foxmail.com",
        phone: "12345678901"
    },
    {
        name: "xj",
        nick: "chaos",
        mail: "chaosxu@foxmail.com",
        phone: "12345678901"
    },
    {
        name: "xj",
        nick: "chaos",
        mail: "chaosxu@foxmail.com",
        phone: "12345678901"
    },
    {
        name: "xj",
        nick: "chaos",
        mail: "chaosxu@foxmail.com",
        phone: "12345678901"
    }
];

@Injectable()
export class AccountService {
    find(): Promise<Account[]> {
        return Promise.resolve(accounts);
    }
}