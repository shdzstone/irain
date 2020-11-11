<template>
    <a-form id="components-form-demo-normal-login"
        :form="form"
        class="login-form"
        @submit="handleSubmit">
        <a-form-item>
            <a-input
                class="login-input"
                v-decorator="[
                'userName',
                { rules: [{ required: true, message: 'Please input your username!' }] },
                ]"
                placeholder="Username"
            >
                <a-icon slot="prefix" type="user" style="color: rgba(0,0,0,.25)" />
            </a-input>
        </a-form-item>
        <a-form-item>
            <a-input-password
                class="login-input"
                v-decorator="[
                'password',
                { rules: [{ required: true, message: 'Please input your Password!' }] },
                ]"
                type="password"
                placeholder="Password"
            >
                <a-icon slot="prefix" type="lock" style="color: rgba(0,0,0,.25)" />
            </a-input-password>
        </a-form-item>
        <a-form-item>
            <a-button type="primary" html-type="submit" class="login-form-button" style="width: 100%; height: 50px;">
                Log in
            </a-button>
        </a-form-item>
    </a-form>
</template>
<script>
import { Form, Input, Icon, Button } from 'ant-design-vue'
const Item = Form.Item
const PassWord = Input.Password
export default {
    name: 'LoginForm',
    components: {
        'a-form': Form,
        'a-form-item': Item,
        'a-input': Input,
        'a-icon': Icon,
        'a-button': Button,
        'a-input-password': PassWord
    },
    beforeCreate() {
        this.form = this.$form.createForm(this, { name: 'normal_login' });
    },
    methods: {
        handleSubmit(e) {
            e.preventDefault();
            this.form.validateFields((err, values) => {
                if (!err) {
                    this.$emit('on-success-valid', {
                        userName: values.userName,
                        password: values.password
                    })
                }
            });
        }
    }
}
</script>
<style lang="less">
.login-input {
    height: 50px;
    .ant-input {
        height: 100%;
    }
}
</style>