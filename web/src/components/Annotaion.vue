<script setup lang="ts">
import { ref } from 'vue'
import VueLabelImage from "vue-image-label";

const IContainer = reactive({
  width: 500,
  height: 500
})

interface Point { x: number; y: number }

const points = ref<Point[]>([])

const getPoints = (params: Point[]) => {
  points.value = params
}
</script>

<template>
  <div class="wrapper">
    <image-label url="https://aldb.obs.cn-east-3.myhuaweicloud.com/img/l.jpg" :containerProps="IContainer"
      :get-points="getPoints" />
    <div v-if="points.length" className="footer">
      <h3>坐标点信息:</h3>
      <ul className="points">
        <li v-for="point in points" :key="point.x + point.y">
          {{ point.x }}, {{ point.y }}
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
.wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100vw;
  height: 100vh;
  padding: 24px 0;
  background-color: #000;
  box-sizing: border-box;
}

.footer {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 1000px;
  background: #333;
  height: 60px;
  margin-top: 24px;
}

h3 {
  color: aliceblue;
}

.points {
  display: flex;
  list-style: none;
  min-width: 600px;
}

.points li {
  margin-left: 24px;
  color: aliceblue;
}
</style>