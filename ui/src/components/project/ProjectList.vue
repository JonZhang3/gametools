<template>
    <a-spin :loading="loading" class="container">
        <template v-if="data.length > 0">
            <a-space direction="vertical" fill>
                <template v-if="data.length > 0">
                    <a-row>
                        <a-col style="text-align: right">
                            <a-space>
                                <a-input-search
                                    v-model="searchForm.name"
                                    @search="loadData"
                                    placeholder="输入项目名称查找"
                                    search-button
                                    style="width: 280px"
                                />
                                <a-tooltip content="刷新" position="top">
                                    <a-button @click="loadData">
                                        <template #icon>
                                            <IconRefresh />
                                        </template>
                                    </a-button>
                                </a-tooltip>
                            </a-space>
                        </a-col>
                    </a-row>
                    <a-row>
                        <a-col>
                            <a-table :columns="columns" :data="data" :pagination="pageData"></a-table>
                        </a-col>
                    </a-row>
                </template>
            </a-space>
        </template>
        <div v-else class="container is-layout-center">
            <a-empty description="暂无相关项目" />
        </div>
    </a-spin>
</template>

<script lang="ts" setup>
import { onMounted, reactive } from "vue";
import type { Project } from "@/apis/project";
import ProjectApi from "@/apis/project";
import useLoading from "@/composables/useLoading";

const { loading, setLoading } = useLoading();
const columns = [
    { title: "名称", dataIndex: "name" },
    { title: "创建时间", dataIndex: "createdAt" },
    { title: "更新时间", dataIndex: "updatedAt" },
];
let data = reactive<Array<Project>>([]);
const pageData = reactive({
    current: 1,
    total: 0,
    defaultPageSize: 30,
});
const searchForm = reactive({
    name: "",
});

async function loadData() {
    setLoading(true);
    try {
        const result = await ProjectApi.listProjects(pageData.current, {
            name: searchForm.name,
        });
        data = result.data.data;
    } finally {
        setLoading(false);
    }
}

onMounted(() => {
    loadData();
});
</script>
