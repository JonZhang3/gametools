export interface HttpResponse<T = unknown> {
    code: number;
    message: string;
    data: T;
}
