import Vue from "vue";
import Router from "vue-router";

import Login from '@/views/Login.vue'
import SignUp from '@/views/Signup.vue'
import createTopic from '@/views/CreateTopic.vue'
import mainPage from '@/views/MainPage.vue'
import TopicPage from "@/views/TopicPage.vue"
import SearchPage from "@/views/SearchTopic.vue"
import Expore from "@/views/Explore.vue"
import UserProfile from "@/views/UserProfile.vue"
import Reply from '@/views/Reply.vue'
import Hot from '@/views/Hot.vue'
Vue.use(Router);

export default new Router({
  mode: "history",
  routes: [
    {
      path: "/login",
      component:Login
    },
    {
        path:"/signup",
        component:SignUp
    },{
      path:"/createTopic",
      component:createTopic
    },{
      path:"/mainPage",
      component:mainPage
    },{
      path:"/t/:tid",
      component:TopicPage,
      props: true
    },{
      path:"/search",
      component:SearchPage
    },{
      path:"/explore",
      component:Expore
    },{
      path:"/u/:uid",
      component:UserProfile,
      props: true
    },{
      path:"/reply",
      component:Reply
    },
    {
      path:"/hot",
      component:Hot
    }
  ]
});
