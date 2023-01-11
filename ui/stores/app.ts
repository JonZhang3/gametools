import { defineStore } from "pinia";

export interface AppState {
    theme: string;
    colorWeek: boolean;
    hideMenu: boolean;
    menuCollapse: boolean;
    themeColor: string;
    menuWidth: number;
    globalSettings: boolean;
    tabBar: boolean;

    [key: string]: unknown;
}

export const useAppState = defineStore("app", {
    state(): AppState {
        return {
            colorWeek: false,
            globalSettings: false,
            hideMenu: false,
            menuCollapse: false,
            menuWidth: 0,
            tabBar: false,
            themeColor: "#165DFF",
            theme: "light"
        };
    },
    actions: {
        toggleTheme() {
            if (this.theme === "light") {
                this.theme = "dark";
                document.body.setAttribute("arco-theme", "dark");
            } else {
                this.theme = "light";
                document.body.removeAttribute("arco-theme");
            }
        }
    }
});
