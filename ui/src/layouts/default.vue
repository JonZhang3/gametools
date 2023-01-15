<template>
    <a-layout class="layout">
        <a-layout-header class="layout-navbar">
            <Header />
        </a-layout-header>
        <a-layout style="overflow: hidden">
            <a-layout-sider breakpoint="xl" collapsible :width="220" hide-trigger>
                <div class="menu-wrapper">
                    <SideMenu :name="props.name" />
                </div>
            </a-layout-sider>
            <a-layout-content class="layout-content">
                <router-view />
            </a-layout-content>
        </a-layout>
    </a-layout>
</template>

<script lang="ts" setup>
import Header from "@/components/layouts/Header.vue";
import SideMenu from "@/components/layouts/SideMenu.vue";

const props = defineProps({
    name: {
        type: String,
        default: "root",
    },
});
</script>

<style lang="less" scoped>
@nav-size-height: 60px;
@layout-max-width: 1100px;

.layout {
    width: 100%;
    height: 100%;
    overflow: hidden;
}

.layout-navbar {
    height: @nav-size-height;
}

.menu-wrapper {
    height: 100%;
    overflow: auto;
    overflow-x: hidden;

    :deep(.arco-menu) {
        ::-webkit-scrollbar {
            width: 12px;
            height: 4px;
        }

        ::-webkit-scrollbar-thumb {
            border: 4px solid transparent;
            background-clip: padding-box;
            border-radius: 7px;
            background-color: var(--color-text-4);
        }

        ::-webkit-scrollbar-thumb:hover {
            background-color: var(--color-text-3);
        }
    }
}

.layout-content {
    padding: 10px;
    overflow: auto;
    background-color: var(--color-fill-2);
    transition: padding 0.2s cubic-bezier(0.34, 0.69, 0.1, 1);
}
</style>
