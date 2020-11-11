<template>
    <div class="irain-login">
        <div class="login-con">
            <h3 class="title">irain-admin</h3>
            <login-form @on-success-valid="handleSubmit"></login-form>
            <p class="login-tip">输入任意用户名和密码即可</p>
        </div>
    </div>
</template>
<script>
import LoginForm from './componetns/login-form'
export default {
    name: 'Login',
    components: {
        LoginForm
    },
    methods: {
        handleSubmit(state) {
            const {userName, password} = state;
            console.log(userName, password)
        }
    }
}
</script>
<style lang="less">
.irain-login {
    min-height: 100vh;
    min-width: 100vw;
    overflow: hidden;
    background-color: #2d3a4b;
    display: flex;
    align-items: center;
    justify-content: center;
    .login-con {
        width: 520px;
        box-sizing: border-box;
        .title {
            font-size: 26px;
            color: #eee;
            margin: 0 auto 40px auto;
            text-align: center;
            font-weight: 700;
        }
        .login-tip{
            font-size: 10px;
            text-align: center;
            color: #c3c3c3;
        }
    }
}
</style>