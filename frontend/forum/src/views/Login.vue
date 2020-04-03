<template>
  <div id="login-page">
    <div id="form-div">
      <el-form status-icon label-width="100px">
        <el-form-item>
          <el-input type="username" placeholder="用户名" v-model="username" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item>
          <el-input type="password" placeholder="密码" v-model="password" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button class="login-button" type="primary" @click="send" block>登录</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script>
import store from "@/store";
import router from "@/route";
export default {
  data() {
    return {
      username: "",
      password: ""
    };
  },
  methods: {
    send() {
      let data = JSON.stringify({
        username: this.username,
        password: this.password,
        method: "GET"
      });

      var theUrl = "http://localhost:8080/api/user";
      fetch(theUrl,{
        method:"POST",
        body:data,
        headers:new Headers({
          'Content-Type': "application/json"
        })
      }).then(res =>{
        return res.json()
      }).then(res => {
        localStorage.setItem('username', this.username);
        localStorage.setItem('uid', res.uid);
        store.commit("login")
        router.push("/mainPage")

      })
    }
  }
};
</script>

<style>
* {
  margin: 0;
  padding: 0;
}

#login-page {
  width: 100%;
  height: 100%;

  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  background-image: url(https://qsf.fs.quoracdn.net/-3-images.home.illo_1920.png-26-5ac607d989ef8067.png);
}
#form-div {
  width: 25%;
}
.login-button {
  width: 100%;
  margin-top: 20px;
  margin-right: 0px;
}
</style>
