<template>
  <div id="pagebody">
    <div id="topic">
      <el-divider></el-divider>
      <router-link :to="topic.user_url">
        <el-avatar
        shape="square"
        class="user-avatar"
        src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
      ></el-avatar>
      </router-link>
      
      <h3 class="topic-title">{{topic.title}}</h3>
      <div class="meta">
        <span class="username">作者:{{topic.author_name}}</span>
        <span class="date">{{ topic.date}}</span>
      </div>
      <el-divider></el-divider>
      <div class="content" v-html="topic.content"></div>
      <el-divider></el-divider>
    </div>
    <div id="comments">
      <div class="comment" v-for="(comment, index) in  this.comments" v-bind:key="comment">
        <div class="meta">
          <router-link :to="comment.user_url">
            <el-avatar
            shape="square"
            class="user-avatar"
            src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
          ></el-avatar>
          </router-link>
          
          <span class="cid"> #{{index+1}}</span>
          <span class="username">作者:{{comment.author_name}}</span>
          
          <span class="date">{{ comment.date}}</span>
        </div>
        <el-divider></el-divider>
        <div class="content" v-html="comment.content"></div>
        <el-divider></el-divider>
      </div>
    </div>
    <div class="input">
      <el-input id="content-area" type="textarea" v-model="input_content" placeholder="输入内容"></el-input>

      <el-button type="primary" @click.prevent="submit">提交</el-button>
    </div>
  </div>
</template>
<script>

import router from "@/route";
import {dateFormate , md2Html} from "@/util"
export default {
  props: ["tid"],
  data() {
    return {
      topic: "",
      comments: "",
      input_content: ""
    };
  },
  methods: {
    submit() {
      let author_id = Number(localStorage.getItem("uid"));
      let author_name = localStorage.getItem("username");
      if (!author_id || !author_name) {
        router.push("/login");
      }
      let data = {
        method: "POST",
        tid: Number(this.tid),
        author_id: author_id,
        author_name: author_name,
        content: this.input_content
      };
      console.log(data);
      fetch("http://localhost:8080/apiv2/comment", {
        method: "POST",
        body: JSON.stringify(data),
        headers: new Headers({
          "Content-Type": "application/json"
        })
      })
        .then(res => {
          return res.json();
        })
        .then(() => {
          
          window.location.reload()
        });
    }
  },
  mounted() {
    this.tid = Number(this.tid);

    let data = {
      method: "GET",
      tid: this.tid
    };

    fetch("http://localhost:8080/apiv2/topic", {
      method: "POST",
      body: JSON.stringify(data),
      headers: new Headers({
        "Content-Type": "application/json"
      })
    })
      .then(res => {
        return res.json();
      })
      .then(res => {
        this.topic = res;

        this.topic.date = dateFormate(this.topic.date)

     
        this.topic.content = md2Html(this.topic.content);
        this.topic.user_url = "/u/" + this.topic.author_id
        
      })
      .catch(e => {
        console.log(e)
      });
    fetch("http://localhost:8080/apiv2/comment", {
      method: "POST",
      body: JSON.stringify(data),
      headers: new Headers({
        "Content-Type": "application/json"
      })
    })
      .then(res => {
        return res.json();
      })
      .then(res => {
        this.comments = res;
        for(let i=0; i<this.comments.length; i++ ){
            this.comments[i].date = dateFormate(this.comments[i].date)
            this.comments[i].content = md2Html(this.comments[i].content)
            this.comments[i].user_url = "/u/" + this.comments[i].author_id
        }
      })
      .catch(e => {
        console.log(e)
      });
  }
};
</script>

<style>
#pagebody {
  width: 50%;
  height: 1000px;

  left: 25%;
  position: relative;
}

.user-avatar {
  float: left;
  bottom: 10px;
  position: relative;
}

.topic-title {
  font-size: 20px;
  font-weight: normal;
  line-height: 21px;
  color: black;
  margin: 0 0 3px 0;
  letter-spacing: 0;
  text-shadow: 0 1px 0 #fff;
}
.meta {
  font-weight: bold;
  font-size: 15px;
  color: #888;
}
.username {
  float: right bottom;
}
.date {
  right: 0px;
  bottom: 0px;
  position: absolute;
}
.meta {
  position: relative;
}
.content {
  text-align: center;
}
.cid{
    float: right top;
}
</style>
