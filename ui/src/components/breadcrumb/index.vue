<template>
    <a-breadcrumb class="container-breadcrumb">
        <a-breadcrumb-item>
            <span v-if="back" @click="handleBack" class="back"><icon-left /> 返回</span>
            <icon-apps v-else />
        </a-breadcrumb-item>
        <a-breadcrumb-item v-for="item in items" :key="item">
            {{ item }}
        </a-breadcrumb-item>
    </a-breadcrumb>
</template>

<script lang="ts" setup>
import { PropType, defineEmits } from "vue";

defineProps({
    items: {
        type: Array as PropType<string[]>,
        required: true,
        default: () => [],
    },
    back: {
        type: Boolean,
        default: false,
    },
});
const emit = defineEmits(["back"]);

function handleBack() {
    emit("back");
}
</script>

<style lang="less" scoped>
.container-breadcrumb {
    margin: 16px 0;

    :deep(.arco-breadcrumb-item) {
        color: rgb(var(--gray-6));

        &:last-child {
            color: rgb(var(--gray-8));
        }
    }

    :deep(.back) {
        cursor: pointer;
    }
}
</style>
