<template>
    <div class="navbar">
        <div class="left-side">
            <a-space>
                <a-typography-title style="margin: 0; font-size: 18px" :heading="5">{{ props.title }}</a-typography-title>
            </a-space>
        </div>
        <ul class="right-side">
            <li>
                <a-tooltip :content="'点击切换为' + (theme === 'light' ? '暗黑模式' : '亮色模式')">
                    <a-button class="nav-btn" type="outline" :shape="'circle'" @click="toggleTheme">
                        <template #icon>
                            <icon-moon-fill v-if="theme === 'dark'" />
                            <icon-sun-fill v-else />
                        </template>
                    </a-button>
                </a-tooltip>
            </li>
        </ul>
    </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { AppStore } from "@/stores";

const props = defineProps({
    title: {
        type: String,
        default: "GameTools",
    },
});

const appStore = AppStore();
const theme = computed(() => {
    return appStore.theme;
});

function toggleTheme() {
    appStore.toggleTheme();
}
</script>

<style scoped lang="less">
.navbar {
    display: flex;
    justify-content: space-between;
    height: 100%;
    background-color: var(--color-bg-2);
}

.left-side {
    display: flex;
    align-items: center;
    padding-left: 20px;
}

.right-side {
    display: flex;
    padding-right: 20px;
    list-style: none;

    :deep(.locale-select) {
        border-radius: 20px;
    }

    li {
        display: flex;
        align-items: center;
        padding: 0 10px;
    }

    .nav-btn {
        border-color: rgb(var(--gray-2));
        color: rgb(var(--gray-8));
        font-size: 16px;
    }
}
</style>
