<template>
    <div class="login-form-wrapper">
        <div class="login-form-title">Game Tools</div>
        <div class="login-form-error-msg">{{ errorMessage }}</div>
        <a-form ref="form" :model="userForm" layout="vertical" @submit="handleLogin">
            <a-form-item field="username"
                         :rules="[{required: true, message: '用户名不能为空'}]"
                         :validate-trigger="['change', 'blur']" hide-label>
                <a-input v-model="userForm.username" size="large" placeholder="用户名">
                    <template #prefix>
                        <icon-user/>
                    </template>
                </a-input>
            </a-form-item>
            <a-form-item field="password"
                         :rules="[{required: true, message: '密码不能为空'}]"
                         :validate-trigger="['change', 'blur']" hide-label>
                <a-input-password v-model="userForm.password" size="large" placeholder="密码" allow-clear>
                    <template #prefix>
                        <icon-lock/>
                    </template>
                </a-input-password>
            </a-form-item>
            <a-button type="primary" size="large" long :loading="loading" html-type="submit">登录</a-button>
        </a-form>
    </div>
</template>

<script lang="ts" setup>
import useLoading from '~/composables/useLoading';

const {loading, setLoading} = useLoading();
const errorMessage = ref('');
const userForm: Record<string, any> = reactive({
    username: '',
    password: '',
});

function handleLogin({errors, values}: {
    errors: Record<string, ValidatedError> | undefined;
    values: object;
}) {

}

</script>

<style lang="less" scoped>
.login-form {
    &-wrapper {
        width: 320px;
    }

    &-title {
        color: var(--color-text-1);
        font-weight: 500;
        font-size: 24px;
        line-height: 32px;
    }

    &-error-msg {
        height: 32px;
        color: rgb(var(--red-6));
        line-height: 32px;
    }
}
</style>
