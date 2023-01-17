<template>
    <a-table :columns="columns" :data="data" :loading="loading" :pagination="pageData"></a-table>
</template>

<script lang="ts" setup>
import { onMounted, reactive, ref } from "vue";
import type { Project } from "@/apis/project";
import ProjectApi from "@/apis/project";

const props = defineProps({
    name: {
        type: String,
        default: null,
    },
    state: {
        type: Number,
        default: 1,
    },
});

const loading = ref(false);
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

async function loadData() {
    loading.value = true;
    try {
        const result = await ProjectApi.listProjects(pageData.current);
        data = result.data.data;
    } finally {
        loading.value = false;
    }
}

onMounted(() => {
    loadData();
});
</script>
