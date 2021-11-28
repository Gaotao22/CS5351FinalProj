<template>
    <div class="login_container">
        <div class="word_box" v-show="start_use">
            <p
                style="float: left; font-size: 35px; font-weight: bolder;"
                height="100%"
                width="100%"
            >
                敏捷开发,我们更高效
            </p>
            <p
                style="float: left; font-size: 30px; font-weight: bold;"
                height="100%"
                width="100%"
            >
                但,不仅如此
            </p>
            <Divider style="margin-top: 10px; color: rgb(0, 0, 0);"></Divider>
        </div>
        <div class="img_box">
            <img :src="imgurl" height="100%" width="100%" />
        </div>
        <div class="start_box">
            <Button
                type="primary"
                shape="circle"
                v-show="start_use"
                size="large"
                v-on:click="startclick"
                style="margin-right: 20px;"
                >开始使用</Button
            >
            <Button
                type="primary"
                shape="circle"
                v-show="start_register"
                size="large"
                v-on:click="startclick"
                style="margin-right: 20px;"
                ghost
            >
                注册
            </Button>
            <Button
                type="primary"
                shape="circle"
                v-show="github"
                size="large"
                v-on:click="startclick"
                icon="logo-github"
            >
                GitHub
            </Button>
        </div>
        <div class="login_box" v-show="isstart">
            <p style="margin-top: 1px; font-size: 20px; font-weight: normal;">
                欢迎使用
            </p>
            <Divider />
            <el-form
                ref="loginFormRef"
                :model="loginForm"
                :rules="loginFormRules"
                label-width="0px"
                class="login_form"
            >
                <el-form-item prop="username">
                    <el-input
                        v-model="loginForm.username"
                        prefix-icon="iconfont icon-user"
                    ></el-input>
                </el-form-item>
                <el-form-item prop="password">
                    <el-input
                        v-model="loginForm.password"
                        prefix-icon="iconfont icon-3702mima"
                        type="password"
                    ></el-input>
                </el-form-item>
                <el-form-item class="btns">
                    <el-button type="primary" v-on:click="login"
                        >登录</el-button
                    >
                    <el-button type="info" v-on:click="resetLoginForm"
                        >重置</el-button
                    >
                </el-form-item>
            </el-form>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            activeIndex: '1',
            imgurl: require('../../assets/login.png'),
            start_use: true,
            start_register: true,
            github: true,
            isstart: false,
            loginForm: {
                username: 'admin',
                password: '123456',
            },
            loginFormRules: {
                username: [
                    {
                        required: true,
                        message: '请输入登录名称',
                        trigger: 'blur',
                    },
                    {
                        min: 3,
                        max: 10,
                        message: '长度在 3 到 10 个字符',
                        trigger: 'blur',
                    },
                ],
                password: [
                    {
                        required: true,
                        message: '请输入密码',
                        trigger: 'blur',
                    },
                    {
                        min: 6,
                        max: 15,
                        message: '长度在 6 到 15 个字符',
                        trigger: 'blur',
                    },
                ],
            },
        }
    },
    methods: {
        handleSelect(key, keyPath) {
            console.log(key, keyPath)
        },
        startclick() {
            // eslint-disable-next-line no-unused-expressions,no-sequences
            ;(this.start_use = false),
                (this.start_register = false),
                (this.isstart = true),
                (this.github = false)
        },
        resetLoginForm() {
            return 0
        },
        login() {
            // 表单预验证
            this.$refs.loginFormRef.validate(async () => {
                this.$router.push('/home')
                // git push -u origin login 将本地分支推送到码云上
            })
        },
    },
}
</script>
<style>
.login_container {
    background-color: #2b4b6b;
    height: 100%;
}
.top_menu {
    background-color: #fff;
    position: absolute;
    top: 0%;
    width: 100%;
}
.img_box {
    width: 50%;
    height: 80%;
    background-color: #fff;
    right: 5%;
    top: 10%;
    position: absolute;
    z-index: -1;
}
.word_box {
    width: 350px;
    height: 350px;
    background-color: #fff;
    border-radius: 3px;
    position: absolute;
    left: 20%;
    top: 200px;
    transform: translate(-50%, -50%);
}
.start_box {
    width: 350px;
    height: 20px;
    background-color: #fff;
    border-radius: 3px;
    position: absolute;
    left: 19%;
    top: 150px;
    transform: translate(-50%, -50%);
}
.login_box {
    width: 300px;
    height: 200px;
    background-color: #fff;
    border-radius: 3px;
    position: absolute;
    left: 20%;
    top: 100px;
    transform: translate(-50%, -50%);
}
.btns {
    display: flex;
    justify-content: flex-end;
}

.login_form {
    position: absolute;
    top: 40%;
    width: 100%;
    padding: 0 20px;
    box-sizing: border-box;
}
</style>
