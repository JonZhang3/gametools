import type { CreateAxiosDefaults, AxiosRequestConfig, AxiosResponse, AxiosInstance, Method } from "axios";
import axios from "axios";
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
            if (response.status === config.api.noAuthStatus) {
                Modal.error({
                    title: "登录无效",
                    content: "您已经登出，您可以取消留在该页面，或重新登录",
                    okText: "重新登录",
                    async onOk() {
                        window.location.reload();
                    },
                });
                return Promise.reject(new Error("未登录或登录无效"));
            }
            const res = response.data;
            if (res.code !== config.api.okCode) {
                Message.error({ content: res.message || "请求错误", duration: 5 * 1000 });
                return Promise.reject(new Error(res.message || "请求错误"));
            }
            return res;
        },
        (error: any) => {
            Message.error({
                content: error.msg || "请求错误",
                duration: 5 * 1000,
            });
            return Promise.reject(error);
        }
    );
    return client;
}

class Request {
    client: AxiosInstance;
    constructor() {
        this.client = createClient({
            baseURL: config.api.baseURL || "",
            timeout: config.api.timeout || 10 * 1000,
            withCredentials: config.api.withCredentials,
            headers: {
                "Content-Type": "application/x-www-form-urlencoded;charset=UTF-8",
            },
        });
    }

    request(method: Method, url: string, data?: Data, headers?: object) {
        if (!method) {
            method = "GET";
        }
        const config: Data = {
            method,
            url,
            headers,
            paramsSerializer(params: any) {
                return qs.stringify(params);
            },
            transformRequest: [
                (d: any, headers: any) => {
                    if (headers["Content-Type"]?.indexOf("application/json")) {
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

    get(url: string, params?: Data) {
        return this.request("GET", url, params);
    }

    post(url: string, params?: Data) {
        setTimeout(function () {}, 4000);
        return this.request("POST", url, params);
    }

    put(url: string, params?: Data) {
        return this.request("PUT", url, params);
    }

    delete(url: string, params?: Data) {
        return this.request("DELETE", url, params);
    }
}

export default new Request();
