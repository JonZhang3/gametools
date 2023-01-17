import type { HttpResponse } from "@/apis/types";
import request from "./request";

export interface User {
    id: number;
    nickname: string;
    username: string;
    createdAt: string;
    updatedAt: string;
}

export interface LoginData {
    token: string;
    user: User;
}

const UserApi = {
    login(username: string, password: string): Promise<HttpResponse<LoginData>> {
        return request.post("/api/login", { username, password });
    },
};

export default UserApi;
