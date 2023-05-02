<template>
  
    <el-container >
      <el-container >
        <SideBar/>
        <el-main>
          <router-view/>
        </el-main>
      </el-container>
    </el-container>
  
</template>


<script lang="ts" setup>
import SideBar  from './SideBar.vue'
</script>

<style scoped>
.el-container{
        /*设置内部填充为0，几个布局元素之间没有间距*/
        padding: 0px;
         /*外部间距也是如此设置*/
        margin: 0px;
        /*统一设置高度为100%*/
        height: 100%;
    }
</style>
