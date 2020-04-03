

<template>
  <div id="signup-page">
    <el-form ref="form" :model="form" label-width="80px" label-position="left">
      <el-form-item label="用户头像">
        <UploadIMG></UploadIMG>
      </el-form-item>
      <el-form-item label="用户名">
        <el-input v-model="form.username"></el-input>
      </el-form-item>
      <el-form-item label="邮箱">
        <el-input v-model="form.email"></el-input>
      </el-form-item>
      <el-form-item label="密码">
        <el-input v-model="form.password" show-password></el-input>
      </el-form-item>
      <el-form-item label="生日">
        <el-date-picker type="date" placeholder="选择日期" v-model="form.date" style="width: 100%;"></el-date-picker>
      </el-form-item>

      <el-form-item label="性别">
        <el-radio-group v-model="form.sex">
          <el-radio label="男"></el-radio>
          <el-radio label="女"></el-radio>
        </el-radio-group>
      </el-form-item>
    </el-form>
    <el-button id="submit-button" type="primary" @click.prevent="onSubmit">立即创建</el-button>
  </div>
</template>



<script>
import UploadIMG from "@/components/UploadIMG.vue";

import { BASE_URL } from "@/devconf";
import router from "@/route"

export default {
  data() {
    return {
      form: {
        username: "",
        email: "",
        password: "",
        date: "",
        sex: "",
      }
    };
  },
  methods: {
    onSubmit() {
      console.log("submit!" + BASE_URL);
      let birthday = new Date(this.form.date)
      let unixtime = birthday.getTime()/1000
      let jsonObj = {
        method:"POST",
        username: this.form.username  ,
        password: this.form.password,
        email: this.form.email,
        birthday: unixtime,
        sex:  this.form ==  1 ? true:false
      };

      
      var theUrl = "http://localhost:8080" + "/api/user"
      fetch(theUrl, {
        method: "POST",
        body: JSON.stringify(jsonObj),
        headers: new Headers({
          "Content-Type": "application/json"
        })
      })
        .then(response => {
          return response.json();
        })
        .then(() => {
          alert("!!!")
          router.push("/login")
        })
        .catch(e => {
          console.log(e);
        });
    }
  },
  components: {
    UploadIMG: UploadIMG
  }
};
</script>

<style>
#signup-page {
  margin-top: 5%;
  margin-left: 30%;
  width: 25%;
}
#submit-button {
  position: relative;
  width: 100%;
}
</style>
