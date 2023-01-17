import { defineStore } from "pinia";
import UserApi from "@/apis/user";
import { clearToken, setToken } from "@/auth";

export interface UserState {
    id?: number;
    username: string;
    nickname: string;
    createdAt: string;
    updatedAt: string;
}

const useUserStore = defineStore("user", {
    state(): UserState {
        return {
            id: undefined,
            username: "",
            nickname: "",
            createdAt: "",
            updatedAt: "",
        };
    },
    getters: {
        userInfo(state: UserState): UserState {
            return { ...state };
        },
    },
    actions: {
        async login(username: string, password: string) {
            try {
                const result = await UserApi.login(username, password);
                setToken(result.data.token);
            } catch (e) {
                clearToken();
                throw e;
            }
        },
    },
});

export default useUserStore;
