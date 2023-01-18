<template>
    <a-modal @ok="handleAddProject" @cancel="handleCancel">
        <a-form ref="formRef" :model="projectForm" layout="vertical">
            <a-form-item label="项目名" field="name" :validate-trigger="['change', 'blur']">
                <a-input v-model="projectForm.name" />
            </a-form-item>
            <a-form-item label="描述" field="description" :validate-trigger="['change', 'blur']">
                <a-input v-model="projectForm.description" />
            </a-form-item>
        </a-form>
    </a-modal>
</template>

<script lang="ts" setup>
import { reactive, ref } from "vue";
import type { Project } from "@/apis/project";
import ProjectApi from "@/apis/project";
import type { ValidatedError } from "@arco-design/web-vue/es/form/interface";

const props = defineProps({
    visible: Boolean,
});

const formRef = ref();
const projectForm = reactive<Partial<Project>>({
    name: "",
    description: "",
});

function handleAddProject() {
    formRef.value?.validate(async (errors: undefined | Record<string, ValidatedError>) => {
        if (!errors) {
            await ProjectApi.createProject({ ...projectForm });
        }
    });
}

function handleCancel() {
    formRef.value?.resetFields();
}
</script>
