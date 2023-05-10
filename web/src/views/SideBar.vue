<script setup lang="ts"></script>
<template>
    <el-aside >
        <el-menu unique-opened router     background-color="#333744" text-color="#fff" 
          active-text-color="#409BFF"  default-active="/index/DataSets">
          <a href="/#/index" class="logo">
            <el-icon color="#409EFF" size="50"><IEp-ElementPlus /></el-icon>
            <h1>藻类标注系统</h1>
          </a>
            <el-menu-item index="/index/user">
                <el-icon><IEp-User /></el-icon>
                <span>个人中心</span>
            </el-menu-item>
            <el-menu-item index="/index/DataSet">
                <el-icon> <IEp-EditPen /> </el-icon>
                数据集管理
            </el-menu-item>
            <el-menu-item index="/index/TagSet">
                <el-icon><IEp-PriceTag /></el-icon>
                标签管理
            </el-menu-item>
        </el-menu>
    </el-aside>
</template>

<style scoped>
.el-aside {
    width: 200px;
    background-color: #333744;
    border-right: black;
    line-height: 100vh;

}

.logo{
    display: flex;
    justify-content: center;
    align-items:center;
    height: 60px;
    text-decoration: none;
    color: #eee;
    
}
.el-menu {
    width: 200px;
    border-right: none;
}
</style>