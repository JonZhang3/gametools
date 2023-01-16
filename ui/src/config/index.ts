export default {
    api: {
        baseURL: "http://127.0.0.1:8080",
        timeout: 10 * 1000,
        dataField: "data",
        messageField: "message",
        okCode: 200,
        noTokenCode: 501,
        noAuthStatus: 401,
        withCredentials: true, // 跨域请求时是否需要使用凭证
    },
};
