<template>
  <Header :user-name="userName" :display="true" @search="search" @refresh="fetchData"/>
  <el-main>
    <div class="flex-auto overflow-y-auto">
      <Waterfall :list="list" :row-key="options.rowKey" :gutter="options.gutter"
        :has-around-gutter="options.hasAroundGutter" :width="options.width" :breakpoints="options.breakpoints"
        :img-selector="options.imgSelector" :background-color="options.backgroundColor"
        :animation-effect="options.animationEffect" :animation-duration="options.animationDuration"
        :animation-delay="options.animationDelay" :lazyload="options.lazyload" :load-props="options.loadProps">
        <template #item="{ item, url, index }">
          <div
            class="bg-gray-900 rounded-lg shadow-md overflow-hidden transition-all duration-300 ease-linear hover:shadow-lg hover:shadow-gray-600 group">
            <div class="overflow-hidden">
              <LazyImg :url="url" class="cursor-pointer transition-all duration-300 ease-linear group-hover:scale-105"
                @click="handlePreview(item)" />
            </div>
            <div class="px-4 pt-2 pb-4 flex justify-between items-center border-t border-t-gray-800">
              <h2 class="pb-4 text-gray-50 group-hover:text-yellow-500">
                {{ item.name }}
              </h2>
              <button
                class="px-3 h-7 rounded-full bg-red-500 text-sm text-white shadow-lg transition-all duration-300 hover:bg-red-600"
                @click.stop="showDrawer(item)">
                标注
              </button>
            </div>
          </div>
        </template>
      </Waterfall>
      <Drawer :alga="alga" :drawer="drawer" :user="userEmail" :grid-data="gridData" @update="closeDrawer" />
      <!-- 大图预览 -->
      <el-dialog v-model="previewVisible" :title="previewData.title" width="500px" center>
        <el-image style="width:100%" :src="previewData.url" />
      </el-dialog>
    </div>

    <el-backtop />
  </el-main>
</template>

<script lang="ts" setup>
import { LazyImg, Waterfall } from "vue-waterfall-plugin-next"
import "vue-waterfall-plugin-next/style.css"
import loading from "~/assets/loading.png"
import error from "~/assets/error.png"
import { getAnno, getData, searchAlga } from "~/api/algae"
import { useRouter } from "vue-router"
import { getUser } from "~/api/user"

const router = useRouter()
// 获取用户信息
function userInfo() {
  const userEmail = sessionStorage.getItem("UserEmail") as string
  if (userEmail == null) {
    ElMessage.error("未登录")
    router.push({ name: "Login" })
  }
  const userName = ref("用户")
  const fetchUser = () => {
    getUser(userEmail).then((res) => {
      userName.value = res.data.name
    })
  }
  return {
    userEmail,
    userName,
    fetchUser
  }
}
const { userEmail, userName, fetchUser } = userInfo()
fetchUser()
// 瀑布流图片设置
function useWaterfall() {
  const options = reactive({
    // 唯一key值
    rowKey: "id",
    // 卡片之间的间隙
    gutter: 10,
    // 是否有周围的gutter
    hasAroundGutter: true,
    // 卡片在PC上的宽度
    width: 320,
    // 自定义行显示个数，主要用于对移动端的适配
    breakpoints: {
      1200: {
        // 当屏幕宽度小于等于1200
        rowPerView: 4,
      },
      800: {
        // 当屏幕宽度小于等于800
        rowPerView: 3,
      },
      500: {
        // 当屏幕宽度小于等于500
        rowPerView: 2,
      },
    },
    // 动画效果
    animationEffect: "animate__fadeInUp",
    // 动画时间
    animationDuration: 1000,
    // 动画延迟
    animationDelay: 300,
    // 背景色
    backgroundColor: "#ffffff",
    // imgSelector
    imgSelector: "src",
    // 加载配置
    loadProps: {
      loading,
      error,
    },
    // 是否懒加载
    lazyload: true,
  });
  return options;
}
const options = useWaterfall()
// 预览图片
function usePreview() {
  const previewVisible = ref(false)
  const previewData = reactive({
    title: "",
    url: "",
  })
  const handlePreview = (item: any) => {
    previewData.title = item.river + " · " + item.name
    previewData.url = item.src
    previewVisible.value = true
  }
  return {
    previewVisible,
    previewData,
    handlePreview,
  }
}
const { previewVisible, previewData, handlePreview } = usePreview()
// 数据获取
function useData() {
  const list = ref([])
  const fetchData = () => {
    getData({}).then((res) => {
      list.value = res.data
    })
  }
  const search = (data: any) => {
    if (data.key == "") {
      fetchData()
      return
    }
    if (data.type == "图像" || data.type == "") {
      searchAlga(data.key).then((res) => {
        list.value = res.data
      })
    }
  }
  return {
    list,
    fetchData,
    search,
  }
}
const { list, fetchData, search } = useData()
fetchData()
// 控制抽屉
function useDrawer() {
  // 抽屉打开
  const alga = ref({})
  const drawer = ref(false)
  // 已有标注
  const gridData = ref([])
  const fetchAnno = (item: any) => {
    getAnno(item.name).then((res) => {
      gridData.value = res.data
    });
  };
  const showDrawer = (item: any) => {
    fetchAnno(item)
    alga.value = item
    drawer.value = true
    console.log(gridData.value)
  }
  // 抽屉关闭
  const closeDrawer = (update: any) => {
    drawer.value = false
  }
  return {
    alga,
    drawer,
    gridData,
    showDrawer,
    closeDrawer
  }
}
const { alga, drawer, gridData, showDrawer, closeDrawer } = useDrawer()
</script>

<style scoped>
@import "../styles/index.css";
</style>