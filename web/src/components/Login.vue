<template>
<div class="login">
  <div class="form">
    <el-input class="input" v-model="username" placeholder="请输入用户名" v-on:keyup.enter.native="login"></el-input>
    <el-input class="input" placeholder="请输入密码" v-model="password" show-password v-on:keyup.enter.native="login"></el-input>
    <el-button class="login-btn" type="primary" @click="login">登录</el-button>
    <el-link class="to-regist" href="regist">还没有账号？去注册</el-link>
  </div>
</div>
</template>

<script>
import {LoginAPI} from "@/common/api";

export default {
  name: "Login",
  data() {
    return {
      username: '',
      password: '',
    }
  },
  methods: {
    warning(msg) {
      this.$notify({
        title: '警告',
        message: msg,
        type: 'warning'
      });
    },
    login(){
      if (this.username === '') {
        this.warning('账号不能为空')
        return
      }
      if (this.password === '') {
        this.warning('密码不能为空')
        return
      }
      const data = {
        'username': this.username,
        'password': this.password,
      }
      LoginAPI(data).then(resp=>{
        if (resp.err != null){
          this.warning(resp.msg)
        }else {
          this.$router.push('/')
        }
      })
    }
  }
}
</script>

<style scoped>
.form{
  width: 300px;
  margin: auto;
  padding: 30px;
  background-color: #000000;
  background-color: rgba(0,0,0,0.2);
  border-radius: 5px;
}
.input{
  margin-bottom: 20px;
}
.login-btn{
  width: 100%;
}
.to-regist{
  margin-top: 20px;
  font-size: 10px;
}
</style>