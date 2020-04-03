<template>
    <div >
        <el-avatar shape="square" :size="40" :src="avatar_url"></el-avatar>
    </div>
</template>

<script>
export default {
    data(){
        return{
            avatar_url:""
        }
    },
    methods:{
        getAvatar(){
            return "https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png"
        }
    },
    mounted(){
        this.avatar_url = this.getAvatar()
    }

}
</script>
<style>
.el-avatar{
    margin-top: 10px;
    margin-right: 0px
}
</style>


