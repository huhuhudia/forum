<template>
  <div id="topics-page">
    <div id="topics">
      <div class="topic" v-for="comment in this.pageContent" v-bind:key="comment">
        <el-avatar
          shape="square"
          class="user-avatar"
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
        ></el-avatar>
        <router-link :to="comment.topic_url"><div class="content" v-html="comment.content"></div></router-link>
        

        <div class="meta">
          <span class="username">作者:{{comment.author_name}}</span>
          <span class="date">{{comment.date}}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { dateFormate, md2Html } from "@/util";

export default {
  data() {
    return {
      pageContent: []
    };
  },

  mounted() {
    
    let data = {
      uid: Number(localStorage.getItem("uid"))
    };

    fetch("http://localhost:8080/apiv2/reply", {
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
        this.pageContent = res;
        for (let i = 0; i < this.pageContent.length; i++) {
          this.pageContent[i].topic_url = "/t/" + this.pageContent[i].tid;

          this.pageContent[i].date = dateFormate(this.pageContent[i].date);
          this.pageContent[i].content = md2Html(this.pageContent[i].content);
        }
        console.log(this.pageContent)
      });
  }
};
</script>

<style>
#topics-page {
  width: 40%;

  left: 25%;
  position: relative;
}
#pagination {
  position: relative;
  bottom: 0px;
  margin: 0px;
}
.topic {
  position: relative;
  height: 5%;
  left: auto;
  top: 5%;
  border: 1px solid rgb(189, 197, 189);
}
#topics {
  height: 1000px;
}
.topic-title {
  float: left bottom;
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

a {
  text-decoration: none;
}

.router-link-active {
  text-decoration: none;
}
</style>

