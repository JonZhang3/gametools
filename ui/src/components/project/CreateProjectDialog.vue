<template>
    <a-modal
        v-model:visible="visible"
        title="创建项目"
        @ok="handleAddProject"
        @cancel="handleCancel"
        :mask-closable="false"
        :ok-loading="loading"
    >
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
import { reactive, ref, defineEmits, computed } from "vue";
import type { Project } from "@/apis/project";
import type { ValidatedError } from "@arco-design/web-vue/es/form/interface";
import ProjectApi from "@/apis/project";
import useLoading from "@/composables/useLoading";
import events from "@/stores/events";
import { Message } from "@arco-design/web-vue";

const props = defineProps({
    modelValue: {
        type: Boolean,
        default: false,
    },
});
const emit = defineEmits(["update:modelValue"]);

const { loading, setLoading } = useLoading();
const visible = computed({
    get() {
        return props.modelValue;
    },
    set(v) {
        emit("update:modelValue", v);
    },
});
const formRef = ref();
const projectForm = reactive<Partial<Project>>({
    name: "",
    description: "",
});

function handleAddProject() {
    formRef.value?.validate(async (errors: undefined | Record<string, ValidatedError>) => {
        if (!errors) {
            setLoading(true);
            try {
                await ProjectApi.createProject({ ...projectForm });
                Message.success("项目创建成功");
                events.project.setCreateOk();
            } finally {
                setLoading(false);
            }
        }
    });
}

function handleCancel() {
    formRef.value?.resetFields();
}
</script>