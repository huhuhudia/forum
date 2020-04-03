<template>
  <div id="app">
    <el-container>
      <el-header>
        <el-row>
          <el-col :span="6">
            <img src alt>
          </el-col>
          <el-col :span="3">
            <navbar></navbar>
          </el-col>
          <el-col :span="5">
            <div id="input">
              <el-input
                suffix-icon="el-icon-search"
                v-model="input"
                @change="handleSearch"
                placeholder="输入感兴趣的内容"
              ></el-input>
            </div>
          </el-col>
          <el-col :span="2">
            <img src alt>
          </el-col>
          <el-col :span="4" v-show="!isLogin">
            <div id="link">
              <router-link to="/login">
                <el-button size="small">登录</el-button>
              </router-link>

              <router-link to="/signup">
                <el-button id="join-us-button" size="small">加入我们</el-button>
              </router-link>
            </div>
          </el-col>

          <el-col :span="1" v-show="isLogin">
            <router-link to="/createTopic">
              <el-button type="primary" icon="el-icon-edit" circle></el-button>
            </router-link>
          </el-col>
          <el-col :span="1" v-show="isLogin">
            <router-link to="/reply"><notice></notice></router-link>
            
          </el-col>
          <el-col :span="4" v-show="isLogin">
            <avatar></avatar>
          </el-col>
          <el-col :span="1" v-show="isLogin">
            <el-button @click="handleLogout">注销</el-button>
          </el-col>
        </el-row>
      </el-header>
      <el-main>
        <router-view :pageContent="pageContent"></router-view>
      </el-main>
    </el-container>
  </div>
</template>



<script>
import navbar from "@/components/Navbar.vue";
import Vue from "vue";
import { mapState } from "vuex";
import notice from "@/components/Notice.vue";
import avatar from "@/components/Avatar.vue";
import store from "@/store";
import router from "@/route"
Vue.component("navbar", navbar);
Vue.component("notice", notice);
Vue.component("avatar", avatar);
export default {
  data() {
    return {
      input: "",
      pageContent:""
    };
  },
  methods: {
    handleLogout(){
      localStorage.clear()
      location.reload()
    }
    ,
    f() {
      console.log("!!");
    },
    handleSearch() {
      let data = {
        method: "GET",
        target: this.input,
        offset: 0,
        limit: 100
      };
      fetch("http://localhost:8080/apiv2/mainpage", {
        method: "POST",
        body: JSON.stringify(data),
        headers: new Headers({
          "Content-Type": "application/json"
        })
      })
        .then(response => {
          return response.json();
        })
        .then(res => {
          for (let i = 0; i < res.length; i++) {
            let date = new Date(res[i].date * 1000);
            let datestring = date.toLocaleDateString("zh-CN");
            let timestring = date.toLocaleTimeString("zh-CN");
            res[i].date = datestring + " " + timestring;
            res[i].url = "/t/" + res[i].tid;
          }
          this.pageContent = res;
        })
        .catch(e => {
          console.log(e);
        });
        router.push("/search")
    }
  },
  mounted() {
    if (localStorage.getItem("uid")) {
      store.commit("login");
    }
  },
  computed: mapState(["isLogin"])
};
</script>
<style>
body {
  margin: 0;
  padding: 0;
}
.el-main {
  padding: 0;
}
#main-page {
  position: absolute;
  height: 100%;
  width: 100%;
}

#input {
  margin-top: 10px;
}

.el-button {
  margin-top: 15px;
}
</style>
