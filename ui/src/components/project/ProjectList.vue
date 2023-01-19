<template>
    <a-spin :loading="loading" class="container direction-vertical">
        <a-space direction="vertical" fill>
            <a-row>
                <a-col style="text-align: right">
                    <a-space>
                        <a-input-search
                            v-model="searchForm.name"
                            @search="loadData"
                            allow-clear
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
                    <a-table
                        v-if="data.length > 0"
                        :columns="columns"
                        :data="data"
                        :pagination="pageData"
                        @row-click="handleToDetail"
                    >
                        <template #operation="{ record }">
                            <a-popconfirm content="确定要归档该项目吗？" @ok="handleToArchive(record.id)">
                                <a-link status="danger">归档</a-link>
                            </a-popconfirm>
                        </template>
                    </a-table>
                    <a-empty v-else description="暂无相关项目" />
                </a-col>
            </a-row>
        </a-space>
    </a-spin>
</template>

<script lang="ts" setup>
import { onMounted, reactive, watch } from "vue";
import { useRouter } from "vue-router";
import type { Project } from "@/apis/project";
import ProjectApi from "@/apis/project";
import useLoading from "@/composables/useLoading";
import events from "@/stores/events";
import { Message } from "@arco-design/web-vue";

const props = defineProps({
    type: {
        type: Number,
        default: 1,
    },
});

const router = useRouter();
const { loading, setLoading } = useLoading();
const columns = [
    { title: "名称", dataIndex: "name" },
    { title: "描述", dataIndex: "description" },
    { title: "创建时间", dataIndex: "createdAt" },
    { title: "更新时间", dataIndex: "updatedAt" },
    { title: "操作", slotName: "operation" },
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
            state: props.tpe,
        });
        data = result.data.data;
    } finally {
        setLoading(false);
    }
}

async function handleToArchive(id: number) {
    setLoading(true);
    try {
        await ProjectApi.archiveProject(id);
        Message.success("项目归档成功");
        await loadData();
    } finally {
        setLoading(false);
    }
}

function handleToDetail(record: Project) {
    console.log(record);
    router.push({
        name: "project-root",
        params: {
            id: recordid,
        },
    });
}

onMounted(() => {
    loadData();
});

watch(
    () => props.type,
    () => loadData()
);

events.project.listenCreateOk(() => loadData());
events.project.listenEditOk(() => loadData());
</script>