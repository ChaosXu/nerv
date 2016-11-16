import { Injectable } from '@angular/core';
import { Http, Response, Headers } from '@angular/http';
import 'rxjs/add/operator/toPromise';


@Injectable()
export class RestyService {
    private headers = new Headers({ 'Content-Type': 'application/json' });

    constructor(private http: Http) { }

    find(type: string): Promise<any> {
        const url = `api/objs/${type}`;
        return this.http.get(url, { headers: this.headers })
            .toPromise()
            .then(response => {
                var body = response.json();                
                return body;
            })
            .catch(this.error);
    }

    get(type: string, id: number): Promise<any> {
        const url = `api/objs/${type}/${id}`;
        return this.http.get(url, { headers: this.headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.error);
    }

    create(type: string, instance: any): Promise<any> {
        const url = `api/objs/${type}`;
        var body = JSON.stringify(instance);
        console.log(body)
        return this.http.post(url, body, { headers: this.headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.error);
    }

    update(type: string, id: number, instance: any): Promise<any> {
        const url = `api/objs/${type}/${id}`;
        var body = JSON.stringify(instance);
        console.log(body)
        return this.http.put(url, body, { headers: this.headers })
            .toPromise()
            .then(response => response.json())
            .catch(this.error);
    }

    remove(type: string, id: number): Promise<{ type: string, id: number }> {
        const url = `api/objs/${type}/${id}`;
        return this.http.delete(url, { headers: this.headers })
            .toPromise()
            .then(() => { return { type: type, id: id } })
            .catch(this.error);
    }

    private error(error: any): Promise<any> {
        return Promise.reject(error);
    }
}
