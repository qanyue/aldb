<template>
  <div>
    <div class="big">
      <div class="inner">
        <img src="../assets/notFound.png" alt="not" />
        <p>页面被神秘力量吸走了</p>
        <p>↓</p>
        <el-space class="leftPic">
          <router-link to="/">
            <el-button type="primary" size="large">返回首页</el-button>
          </router-link>
          <el-button class="error-btn" type="primary" size="large" @click="goBack">返回上一页</el-button>
        </el-space>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { useRouter } from "vue-router"

export default {
  name: "404",
  setup() {
    const router = useRouter()
    const goBack = () => {
      router.go(-1)
    };
    return {
      goBack,
    };
  },
};
</script>


<style lang="scss">
.big {
  width: 100%;
  height: calc(100vh);
  background-color: rgb(244, 244, 244);
  position: relative;
}

.inner {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

.inner p {
  text-align: center;
  font-size: 24px;
}

.inner .leftPic {
  display: flex;
  justify-content: center;
}
</style>