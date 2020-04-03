<template>
  <el-row>
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
    >
      <el-menu-item index="1">首页</el-menu-item>
      <el-menu-item index="2">发现</el-menu-item>
      <el-menu-item index="3">热榜</el-menu-item>
    </el-menu>
  </el-row>
</template>
<style>
.el-menu-demo {
  border: 0ch !important;
  margin: 0ch !important;
}
</style>

<script>
import router from "@/route";
export default {
  data() {
    return {
      input: "",
      activeIndex: "1"
    };
  },
  methods: {
    handleSearch() {
      console.log(this.input);
    },
    handleSelect(key, keypath) {
      console.log(key, "          ", keypath);
      switch (key) {
        case "1":
          router.push("/mainPage");
          break;
        case "2":
          router.push("/explore");
          break;
        case "3":
          router.push("/hot");
          break;
      }
    }
  }
};
</script>
