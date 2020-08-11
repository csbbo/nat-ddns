<template>
  <div class="regist">
    <div class="form">
      <el-input class="input" v-model="username" placeholder="请输入用户名" v-on:keyup.enter.native="regist"></el-input>
      <el-input class="input" placeholder="请输入密码" v-model="password1" show-password v-on:keyup.enter.native="regist"></el-input>
      <el-input class="input" placeholder="请再次输入密码" v-model="password2" show-password v-on:keyup.enter.native="regist"></el-input>
      <el-button class="login-btn" type="primary" @click="regist">注册</el-button>
      <el-link class="to-login" href="/login">已有账号？去登录</el-link>
    </div>
  </div>
</template>

<script>
import {RegistAPI} from "@/common/api";

export default {
  name: "Regist",
  data() {
    return {
      username: '',
      password1: '',
      password2: '',
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
    succeed(msg) {
      this.$notify({
        title: '成功',
        message: msg,
        type: 'success'
      });
    },
    regist(){
      if (this.username === '') {
        this.warning('账号不能为空')
        return
      }
      if (this.password1 === '' || this.password2 === '') {
        this.warning('密码不能为空')
        return
      }
      if (this.password1 != this.password2) {
        this.warning('两次密码不一致')
        return
      }
      const data = {
        'username': this.username,
        'password': this.password1,
      }
      RegistAPI(data).then(resp=>{
        if (resp.err != null){
          this.warning(resp.msg)
        }else {
          this.succeed('注册成功')
          this.$router.push('/login')
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
.to-login{
  margin-top: 20px;
  font-size: 10px;
}
</style>