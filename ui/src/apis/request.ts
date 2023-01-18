import type { CreateAxiosDefaults, AxiosRequestConfig, AxiosResponse, AxiosInstance, Method } from "axios";
import axios, { AxiosError } from "axios";
import qs from "query-string";
import { Message, Modal } from "@arco-design/web-vue";
import { getToken } from "@/auth";
import config from "@/config";
import type { HttpResponse, Data } from "./types";

function createClient(settings: CreateAxiosDefaults) {
    const client = axios.create(settings);
    client.interceptors.request.use(
        (config: AxiosRequestConfig) => {
            const token = getToken();
            if (token) {
                if (!config.headers) {
                    config.headers = {};
                }
                // @ts-ignore
                config.headers.Authorization = `Bearer ${token}`;
            }
            return config;
        },
        (error: any) => Promise.reject(error)
    );
    client.interceptors.response.use(
        // @ts-ignore
        (response: AxiosResponse<HttpResponse>) => {
            const res = response.data;
            if (res.code !== config.api.okCode) {
                Message.error({ content: res.message || "请求错误", duration: 5 * 1000 });
                return Promise.reject(new Error(res.message || "请求错误"));
            }
            return res;
        },
        (error: Error) => {
            if (error instanceof AxiosError) {
                if (error.response?.status === config.api.noAuthStatus) {
                    Modal.error({
                        title: "登录无效",
                        content: "您已经登出，您可以取消留在该页面，或重新登录",
                        okText: "重新登录",
                        async onOk() {
                            window.location.reload();
                        },
                    });
                    return Promise.reject(new Error("未登录或登录失效"));
                }
            }
            Message.error({
                content: error.message || "请求错误",
                duration: 5 * 1000,
            });
            return Promise.reject(error);
        }
    );
    return client;
}

class Request {
    private client: AxiosInstance;
    json: JsonRequest;

    constructor() {
        this.client = createClient({
            baseURL: config.api.baseURL || "",
            timeout: config.api.timeout || 10 * 1000,
            withCredentials: config.api.withCredentials,
            headers: {
                "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
            },
        });
        this.json = new JsonRequest(this);
    }

    request<T>(method: Method, url: string, data?: Data, headers?: object): Promise<HttpResponse<T>> {
        if (!method) {
            method = "GET";
        }
        const config: Data = {
            method,
            url,
            headers,
            paramsSerializer: {
                encode: qs.parse,
                serialize: qs.stringify,
            },
            transformRequest: [
                (d: any, headers: any) => {
                    if (headers["Content-Type"]?.indexOf("application/json") > -1) {
                        return JSON.stringify(d);
                    }
                    return qs.stringify(d);
                },
            ],
        };
        if (["PUT", "POST", "PATCH"].includes(method)) {
            config.data = data;
        } else {
            config.params = data;
        }
        return this.client.request(config);
    }

    get<T>(url: string, params?: Data): Promise<HttpResponse<T>> {
        return this.request("GET", url, params);
    }

    post<T>(url: string, params?: Data): Promise<HttpResponse<T>> {
        return this.request("POST", url, params);
    }

    put<T>(url: string, params?: Data): Promise<HttpResponse<T>> {
        return this.request("PUT", url, params);
    }

    delete<T>(url: string, params?: Data): Promise<HttpResponse<T>> {
        return this.request("DELETE", url, params);
    }
}

class JsonRequest {
    parent: Request;

    constructor(parent: Request) {
        this.parent = parent;
    }

    post<T>(url: string, params: Data): Promise<HttpResponse<T>> {
        return this.parent.request("POST", url, params, {
            "Content-Type": "application/json;charset=UTF-8",
        });
    }

    put<T>(url: string, params: Data): Promise<HttpResponse<T>> {
        return this.parent.request("PUT", url, params, {
            "Content-Type": "application/json;charset=UTF-8",
        });
    }
}

const request = new Request();

export default request;
