<template>

  <div id="createTopic-page">
    <el-form label-width="100px">
      <el-form-item lable="标题">
        <el-input id="title-area" v-model="title" placeholder="输入标题"></el-input>
      </el-form-item>
      <el-form-item>
        <el-input id="content-area" type="textarea" v-model="content" placeholder="输入内容"></el-input>
      </el-form-item>
      <el-form-item>
          <el-button type="primary" @click.prevent="submit">提交</el-button>
      </el-form-item>
    </el-form>
  </div>
  
</template>

<script>
import router from "@/route"
export default {
    data(){
        return {
            title:"",
            content:""
        }
    },
    methods:{
        submit(){
            let author_id = Number(localStorage.getItem("uid"))
            let author_name = localStorage.getItem("username")
            if( !author_id  || !author_name ){
                router.push("/login")
            }

            let data = {
                method:"POST",
                author_id:author_id,
                author_name:author_name,
                title:this.title,
                content:this.content

            }
            fetch("http://localhost:8080/apiv2/topic",{
                method:"POST",
                body:JSON.stringify(data),
                headers:new Headers({
                    'Content-Type': "application/json"
                })
            }).then(res => {
                return res.json()
            }).then(res=>{
                console.log(res)
                router.push("/t/" + res.tid)
            })
        }
    }
};
</script>

<style>
#createTopic-page{
  margin-top: 5%;
  margin-left: 30%;
  width: 40%;
}
#content-area{
    height: 300px;
}
</style>
