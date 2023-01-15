import { defineStore } from "pinia";

export interface AppState {
    theme: string;
    hideMenu: boolean;
    menuCollapse: boolean;
    menuWidth: number;
    globalSettings: boolean;
    tabBar: boolean;
}

const AppStore = defineStore("app", {
    state(): AppState {
        return {
            globalSettings: false,
            hideMenu: false,
            menuCollapse: false,
            menuWidth: 0,
            tabBar: false,
            theme: "light",
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
        },
    },
});

export default AppStore;
