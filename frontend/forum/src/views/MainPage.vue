<template>
  <div id="topics-page">
    <div id="topics">
      <div class="topic" v-for="topic in this.pageContent" v-bind:key="topic.id">
        <el-avatar
          shape="square"
          class="user-avatar"
          src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
        ></el-avatar>
        <router-link :to="topic.url"> <h3 class="topic-title">{{topic.title}}</h3></router-link>
        
        <div class="meta">
          <span class="username">  作者:{{topic.author_name}}</span>
          <span class="date">{{ topic.date}}</span>
        </div>
      </div>
    </div>

    <div id="pagination">
      <el-pagination
        :current-page.sync="curpage"
        background
        layout="prev, pager, next"
        :total="1000"
        @current-change="handleChange(curpage)"
      ></el-pagination>
    </div>
  </div>
</template>

<script>
const PAGE_SIZE = 18;
export default {
  data() {
    return {
      curpage: 1,
      pageContent: ""
    };
  },
  methods: {
    handleChange(curpage) {
      console.log("current page " + curpage, "  " + this.curpage);
      let data = {
        method: "GET",
        offset: PAGE_SIZE * (curpage - 1),
        limit: PAGE_SIZE
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
          
          for(let i=0; i<res.length; i++){
            let date = new Date(res[i].date*1000)
            let datestring  = date.toLocaleDateString("zh-CN")
            let timestring = date.toLocaleTimeString("zh-CN")
            res[i].date  = datestring + " " + timestring
            res[i].url = "/t/" + res[i].tid
          }
          this.pageContent = res;
        })
        .catch(e => {
          console.log(e);
        });
      console.log("content : " + this.pageContent);
    }
  },
  mounted() {
    this.handleChange(1);
    console.log(" !!! " + this.pageContent);
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
  float: left ;
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
.meta{
   font-weight: bold;
  font-size: 15px;
  color: #888;
}
.username {
 
  float:right bottom;;
}
.date{
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

