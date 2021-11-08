<template>
  <div class = "login_container">
    <div class="top_menu">
    <el-menu :default-active="activeIndex" style="float: right;margin-right: 5%" class="el-menu-demo" mode="horizontal" v-on:click="handleSelect" background-color="#fff"
             text-color="#000000FF"
             active-text-color="#000000FF">
      <el-menu-item index="1">加入我们</el-menu-item>
      <el-submenu index="2">
        <template slot="title">我的工作台</template>
        <el-menu-item index="2-1">选项1</el-menu-item>
        <el-menu-item index="2-2">选项2</el-menu-item>
        <el-menu-item index="2-3">选项3</el-menu-item>
        <el-submenu index="2-4">
          <template slot="title">选项4</template>
          <el-menu-item index="2-4-1">选项1</el-menu-item>
          <el-menu-item index="2-4-2">选项2</el-menu-item>
          <el-menu-item index="2-4-3">选项3</el-menu-item>
        </el-submenu>
      </el-submenu>
      <el-menu-item index="3" disabled>消息中心</el-menu-item>
      <el-menu-item index="4"><a href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
    </el-menu>
    </div>
    <div class="word_box" v-show = "start_use">
      <p style="float: left;font-size:35px;font-weight:bolder" height="100%" width="100%">敏捷开发,我们更高效</p>
      <p style="float: left;font-size:30px;font-weight:bold" height="100%" width="100%">但,不仅如此</p>
      <Divider style = "margin-top: 10px;color: rgb(0,0,0)"></Divider>
    </div>
    <div class="img_box">
      <img :src="imgurl" height="100%" width="100%">
    </div>
    <div class = "start_box">
        <Button type="primary" shape="circle" v-show = "start_use" size="large" v-on:click="startclick" style="margin-right: 20px">开始使用</Button>
        <Button type="primary" shape="circle" v-show = "start_register" size="large" v-on:click="startclick" style="margin-right: 20px" ghost> 注册 </Button>
        <Button type="primary" shape="circle" v-show = "github" size="large" v-on:click="startclick" icon="logo-github"> GitHub </Button>
    </div>
    <div class = "login_box" v-show="isstart">
      <p style="margin-top:1px;font-size:20px;font-weight:normal" >欢迎使用</p>
      <Divider />
      <el-form ref="loginFormRef" :model="loginForm" :rules = "loginFormRules" label-width = "0px" class="login_form">
        <el-form-item prop = "username">
          <el-input v-model="loginForm.username" prefix-icon="iconfont icon-user"></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input v-model="loginForm.password" prefix-icon="iconfont icon-3702mima" type="password"></el-input>
        </el-form-item>
        <el-form-item class="btns">
          <el-button type="primary" v-on:click="login">登录</el-button>
          <el-button type="info" v-on:click="resetLoginForm">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script>
export default {
  data () {
    return {
      activeIndex: '1',
      imgurl: require('../assets/login.png'),
      start_use: true,
      start_register: true,
      github: true,
      isstart: false,
      loginForm: {
        username: 'admin',
        password: '123456'
      },
      loginFormRules: {
        username: [
          {
            required: true,
            message: '请输入登录名称',
            trigger: 'blur'
          },
          {
            min: 3,
            max: 10,
            message: '长度在 3 到 10 个字符',
            trigger: 'blur'
          }
        ],
        password: [
          {
            required: true,
            message: '请输入密码',
            trigger: 'blur'
          },
          {
            min: 6,
            max: 15,
            message: '长度在 6 到 15 个字符',
            trigger: 'blur'
          }
        ]
      }
    }
  },
  methods: {
    handleSelect (key, keyPath) {
      console.log(key, keyPath)
    },
    startclick () {
      // eslint-disable-next-line no-unused-expressions,no-sequences
      this.start_use = false,
      this.start_register = false,
      this.isstart = true,
      this.github = false
    },
    resetLoginForm () {
      return 0
    },
    login () {
      // 表单预验证
      this.$refs.loginFormRef.validate(async valid => {
        // if (!valid) return // 如果表单校验不通过，就直接返回
        // // { data: res } ：将 data 重命名为 res
        // const { data: res } = await this.$http.post('login', this.loginForm)
        // console.log(res.meta)
        // if (res.meta.status !== 200) return this.$message.error(res.meta.msg)
        // this.$message.success('登录成功')
        // // 1.将登录成功之后的token 保存到客户端的 sessionStorage 中
        // //  1.1 项目中除了登录之外的其他接口，必须登录后才能访问
        // //  1.2 token 只应在当前网站打开期间生效，所以将 token 保存到 sessionStorage 中.
        // //  sessionStorage 是会话保存机制 而 localStorage 是持久化的存储机制
        // console.log(res)
        // window.sessionStorage.setItem('token', res.data.token)
        // // 2.通过编程式导航跳转到后台主页，路由地址是 /home
        this.$router.push('/home')
        // git push -u origin login 将本地分支推送到码云上
      })
    }
  }
}
</script>
<style>
.login_container {
  background-color: #2b4b6b;
  height: 100%;
}
.top_menu{
  background-color: #fff;
  position: absolute;
  top: 0%;
  width: 100%;
}
.img_box{
  width: 50%;
  height: 80%;
  background-color: #fff;
  right:5%;
  top: 10%;
  position: absolute;
  z-index: -1;
}
.word_box{
  width: 350px;
  height: 350px;
  background-color: #fff;
  border-radius: 3px;
  position: absolute;
  left: 20%;
  top: 55%;
  transform: translate(-50%, -50%);
}
.start_box{
  width: 350px;
  height: 20px;
  background-color: #fff;
  border-radius: 3px;
  position: absolute;
  left: 19%;
  top: 70%;
  transform: translate(-50%, -50%);
}
.login_box {
  width: 300px;
  height: 200px;
  background-color: #fff;
  border-radius: 3px;
  position: absolute;
  left: 20%;
  top: 45%;
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
