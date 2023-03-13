<template>
  <el-header class="header">
    <router-link to="/">
      <div class="logo">数据标注存储系统</div>
    </router-link>

    <div v-if="display" class="input">
      <el-input v-model="search" placeholder="请输入..." size="large" clearable @change="searchData">
        <template #prepend>
          <el-select v-model="select" placeholder="类型" style="width: 80px">
            <el-option label="图像" value="图像" />
            <el-option label="标注" value="标注" />
          </el-select>
        </template>
        <template #append>
          <el-button :icon="Search" @click="searchData" />
        </template>
      </el-input>
    </div>


    <div class="header-right">
      <div class="header-user-con">
        <Upload @refresh="refreshData" />
        <el-dropdown class="user-name" trigger="click" @command="handleCommand">
          <span class="el-dropdown-link">
            <el-avatar> {{ userName }} </el-avatar>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="user">个人中心
              </el-dropdown-item>
              <el-dropdown-item divided command="out">退出登录
              </el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>
  </el-header>
</template>

<script lang="ts" setup>
import { useRouter } from "vue-router"
import { Search } from '@element-plus/icons-vue'
import { searchAlga } from "~/api/algae";
const router = useRouter()
const props = defineProps(['userName', 'display'])
const emit = defineEmits(['search', 'refresh'])

const handleCommand = (command: string | number | object) => {
  if (command == "out") {
    sessionStorage.clear()
    router.push({ name: "Login" })
  } else if (command == "user") {
    router.push({ name: "User" })
  }
}

const search = ref('')
const select = ref('图像')

const searchData = () => {
  emit('search', { key: search.value, type: select.value })
}
const refreshData = (data: any) => {
  emit('refresh', data)
}
</script>

<style scoped>
.header {
  position: relative;
  box-sizing: border-box;
  width: 100%;
  height: 70px;
  font-size: 22px;
  border-bottom: solid 1px #dfe1e5;
}

.header .logo {
  float: left;
  width: 250px;
  line-height: 70px;
}

.header .input {
  float: left;
  width: 500px;
  height: 70px;
  display: flex;
  align-items: center;
  padding-left: 50px;
}

.header-right {
  float: right;
  padding-right: 50px;
}

.header-user-con {
  display: flex;
  height: 70px;
  align-items: center;
}

.user-name {
  margin-left: 10px;
}

.el-dropdown-link {
  /* color: #fff; */
  cursor: pointer;
}
</style>
