export interface HttpResponse<T = unknown> {
    code: number;
    message: string;
    data: T;
}

export interface PageData<T = unknown> {
    total: number;
    data: Array<T>;
}

export type Data = {
    [key: string]: any;
};
