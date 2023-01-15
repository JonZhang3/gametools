<template>
    <a-table :columns="columns" :data="data" :loading="loading"></a-table>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";

const loading = ref(false);
const columns = [
    { title: "名称", dataIndex: "name" },
    { title: "创建者", dataIndex: "a" },
];
const data = reactive([]);
</script>
