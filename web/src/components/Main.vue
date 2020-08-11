<template>
<div class="main">
  <el-menu
      class="el-menu-demo"
      mode="horizontal"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b">
    <el-menu-item class="platform-name">NAT-DDNS</el-menu-item>
    <el-submenu index="5" class="user-manager">
      <template slot="title"><i class="el-icon-user-solid"></i></template>
      <el-menu-item index="5-1">用户信息</el-menu-item>
      <el-menu-item index="5-2" @click="logout">注销</el-menu-item>
    </el-submenu>
  </el-menu>
</div>
</template>

<script>
import {LogoutAPI} from "@/common/api";

export default {
  name: "Main",
  data() {
    return {

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
    logout() {
      LogoutAPI().then(resp=>{
        if (resp.err === null) {
          console.log("logout")
          this.$router.push('/login')
        } else {
          this.warning("用户未登录")
        }
      })
    }
  }
}
</script>

<style scoped>
.el-menu-demo{
  margin: -8px;
}
.user-manager{
  float: right;
}
.platform-name{
  font-size: large;
}
</style>