<template>
  <div>
    <el-badge id="info-border" :value="replyCnt" class="item">
    <el-button  id="info" type="info" icon="el-icon-message" circle></el-button>  
    </el-badge>
    
  </div>
</template>
<script>


export default {
  data() {
    return {
      replyCnt: 0
    };
  },
  methods: {
    getReplyCnt() {
      return 10;
    }
  },
  mounted() {
    this.replyCnt = this.getReplyCnt();
  }
};
</script>
<style>
#info{
     
}
#info-border{
    
}
</style>


