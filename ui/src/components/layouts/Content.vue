<template>
    <a-layout class="main-content">
        <a-layout-header class="content-header">
            <a-typography-title class="content-header-title" :heading="5">{{ title }}</a-typography-title>
            <a-divider v-if="$slots.extra" direction="vertical" />
            <div class="content-header-extra">
                <slot name="extra"></slot>
            </div>
        </a-layout-header>
        <a-layout-content class="content-content">
            <slot></slot>
        </a-layout-content>
    </a-layout>
</template>

<script lang="ts" setup>
defineProps({
    title: {
        type: String,
        required: true,
    },
});
</script>

<style lang="less">
.main-content {
    background-color: var(--color-bg-2);
    height: 100%;
}

.content-header {
    display: flex;
    align-items: center;
    height: 64px;
}

h5.arco-typography.content-header-title {
    margin: 0;
    padding: 0 24px;
    font-weight: 700;
    font-size: 18px;
}

.content-header-extra {
    flex: 1;
    display: flex;
    align-items: center;
    padding-right: 24px;
}

.content-content {
    overflow: auto;
    padding: 5px;
}
</style>
