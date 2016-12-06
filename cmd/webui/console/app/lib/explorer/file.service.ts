export class File {
    name: string;
    type: string;
    title: string;
    content?: string;
    dirty?: boolean = false;
    extension?: string = '';
    children?: Array<File>
}

export class FileService {

}