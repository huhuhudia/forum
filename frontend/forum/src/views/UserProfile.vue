<template>
  <div id="profile-page">
    <div>
      <el-divider></el-divider>
      <div>
        <el-avatar
          size="large"
          shape="square"
          class="user-avatar"
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
        ></el-avatar>
      </div>

      <el-button v-if="!isHisFan" @click="handleSubsribe">关注</el-button>
      <el-button v-if="isHisFan">取消关注</el-button>
      <el-divider></el-divider>
    </div>

    <div id="comments">
      <div class="comment" v-for="(comment, index) in  this.comments" v-bind:key="comment">
        <div class="meta">
          <el-avatar
            shape="square"
            class="user-avatar"
            src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
          ></el-avatar>

          <span class="cid">#{{index+1}}</span>
          <span class="username">作者:{{comment.author_name}}</span>

          <span class="date">{{ comment.date}}</span>
        </div>
        <el-divider></el-divider>
        <router-link :to="comment.topic_url">
          <div class="content" v-html="comment.content"></div>
        </router-link>

        <el-divider></el-divider>
      </div>
    </div>
  </div>
</template>


<script>
import { dateFormate, md2Html } from "@/util";

export default {
  props: ["uid"],
  data() {
    return {
      comments: [],
      isHisFan: false
    };
  },
  methods: {
    handleSubsribe() {
      let self_id = Number(localStorage.getItem("uid"));
      let other_id = Number(this.uid);
      let addFanReq = {
        method: "POST",
        self_id: other_id,
        other_id: self_id
      };
      let addFollowReq = {
        method: "POST",
        self_id: self_id,
        other_id: other_id
      };
    console.log(addFanReq)
    console.log(addFollowReq)
      fetch("http://localhost:8080/api/fans", {
        method: "POST",
        body: JSON.stringify(addFanReq),
        headers: new Headers({
          "Content-Type": "application/json"
        })
      })
        .then(() => {
              this.isHisFan = true;
        })
        .catch(e => {
        //   alert(e);
        console.log(e)
        });

      fetch("http://localhost:8080/api/follow", {
        method: "POST",
        body: JSON.stringify(addFollowReq),
        headers: new Headers({
          "Content-Type": "application/json"
        })
      })
       
        .then(() => {
          this.isHisFan = true;
        })
        .catch(e => {
          alert(e);
        });
    }
  },
  mounted() {
    let url = "http://localhost:8080/apiv2/commentsOfAuthor";
    let data = {
      author_id: Number(this.uid)
    };
    fetch(url, {
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
        for (let i = 0; i < this.comments.length; i++) {
          this.comments[i].date = dateFormate(this.comments[i].date);
          this.comments[i].content = md2Html(this.comments[i].content);
          this.comments[i].topic_url = "/t/" + this.comments[i].tid;
        }
      })
      .catch(e => {
        alert(e);
      });

    let fansreq = {
      method: "GET",
      self_id: Number(this.uid)
    };
    fetch("http://localhost:8080/api/fans", {
      method: "POST",
      body: JSON.stringify(fansreq),
      headers: new Headers({
        "Content-Type": "application/json"
      })
    })
      .then(res => {
        return res.json();
      })
      .then(res => {
        
        console.log(res);
        if (res.indexOf(localStorage.getItem("uid")) != -1) {
          this.isHisFan = true;
        }
      })
      .catch(e => {
        alert(e);
      });
  }
};
</script>
<style>
#profile-page {
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
.cid {
  float: right top;
}
.router-link-active {
  text-decoration: none;
}
</style>

