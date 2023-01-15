import type { AxiosRequestConfig, AxiosResponse } from "axios";
import axios from "axios";
import { Message, Modal } from "@arco-design/web-vue";
import { getToken } from "@/auth";
import config from "@/config";
import type { HttpResponse } from "./types";

if (config.api.baseURL) {
    axios.defaults.baseURL = config.api.baseURL;
    axios.defaults.timeout = config.api.timeout;
}

axios.interceptors.request.use(
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
    (error) => {
        return Promise.reject(error);
    }
);

axios.interceptors.response.use(
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
    (error) => {
        Message.error({
            content: error.msg || "请求错误",
            duration: 5 * 1000,
        });
        return Promise.reject(error);
    }
);
