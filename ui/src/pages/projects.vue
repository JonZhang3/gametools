<template>
    <Content title="项目">
        <template #extra>
            <div class="header-tabs">
                <a-tabs @change="handleTabsChange">
                    <a-tab-pane v-for="item in tabs" :key="item.key" :title="item.title"></a-tab-pane>
                </a-tabs>
            </div>
            <div>
                <a-button type="primary" @click="handleToCreateProject">
                    <template #icon>
                        <icon-plus />
                    </template>
                    创建项目
                </a-button>
            </div>
        </template>
        <template #default>
            <ProjectList :type="currentKey" />
            <CreateProjectDialog v-model="createDialogVisible" />
        </template>
    </Content>
</template>

<script lang="ts" setup>
import Content from "@/components/layouts/Content.vue";
import ProjectList from "@/components/project/ProjectList.vue";
import CreateProjectDialog from "@/components/project/CreateProjectDialog.vue";
import { ref } from "vue";

const createDialogVisible = ref(false);

const tabs = [
    { key: 1, title: "我参与的" },
    { key: -1, title: "已归档的" },
];

const currentKey = ref(1);

function handleTabsChange(key: number) {
    currentKey.value = key;
}

function handleToCreateProject() {
    createDialogVisible.value = true;
}
</script>

<style lang="less" scoped>
.header-tabs {
    flex: 1;

    :deep(.arco-tabs-content) {
        display: none;
    }

    :deep(.arco-tabs-nav::before) {
        display: none;
    }
}
</style>