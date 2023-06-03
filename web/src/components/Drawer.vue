<template>
  <el-drawer :model-value="drawer" :title="alga.name" size="50%" :before-close="handleClose" destroy-on-close>
    <el-table :data="gridData" height="250">
      <el-table-column property="createAt" label="创建时间" />
      <el-table-column property="updateAt" label="修改时间" />
      <el-table-column property="description" label="描述" />
      <el-table-column label="标签">
        <el-table-column prop="tag.name" label="名称" />
        <el-table-column prop="tag.resourceName" label="resourceName" width="120" />
      </el-table-column>
      <el-table-column label="操作">
        <template #default="scope">
          <el-button size="small" type="success" @click="deleteAnnotation(scope.row)">删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <br>
    <template #title>
      <h3>添加标注</h3>
    </template>
    <el-form :model="form">
      <el-form-item label="描述">
        <el-input v-model="form.description" autosize type="textarea" clearable placeholder="请输入标注描述" />
      </el-form-item>
      <el-form-item label="标签">
        <div class="wrapper">
          <image-label :url="algaStore.algaInfo.url" :containerProps="iContainer" :get-points="getPoints" />
        </div>
        <div class="znlh-tab-elect">
          <br>
          <el-input v-model="searchTagText" placeholder="搜索标签" class="input-with-select">
            <template #append>
              <el-button type="success" :icon="Search" @click="searchTag()" />
            </template>
          </el-input>
          <el-tag v-for="(item, index) in tagset" :key="item.name" closable @click="handleRadio(1, item, index)"
            @close="handleTagClose(item)" :class="{ active: index == activeIndex }" effect="plain">
            {{ item.name.length > 10 ? item.name.substring(0, 10) + '...' : item.name }}
          </el-tag>
        </div>
      </el-form-item>
    </el-form>
    <el-button size="large" @click="cancelForm">取消</el-button>
    <el-button type="primary" size="large" @click="submitForm">提交</el-button>
    <el-button size="large">下一个图片</el-button>
  </el-drawer>
</template>

<script lang="ts" setup>
import { Tag, addAnno, deleteAnno } from "~/api/algae"
import { UploadFile, UploadFiles, UploadProps } from "element-plus/es"
import { Annotation } from "~/api/algae";
import { deleteTag, getAllTags } from "~/api/tag";
import { useAlgaStore } from "~/store/alga";
import { Search } from '@element-plus/icons-vue'
import { useTagStore } from "~/store/tagSet";


const iContainer = ref({
  width: 500,
  height: 500,
})

const searchTagText = ref("")
const tagStore = useTagStore()
interface Point {
  x: number;
  y: number
}

const points = ref<Point[]>([])

//push points to form.segmentation
const getPoints = (params: Point[]) => {
  let pointNums: number[] = []
  params.forEach((value: Point) => {
    pointNums.push(value.x)
    pointNums.push(value.y)
  })
  form.segmentation = pointNums
}
const algaStore = useAlgaStore()
const props = defineProps(['alga', 'drawer', 'user', 'gridData'])
const emit = defineEmits(['update'])

console.log("url:   " + algaStore.algaInfo.url)
let activeIndex = ref(-1);
let labelCodeList: number[] = reactive([]);

// 单选选中
function handleRadio(type: number, tag: Tag, index: number) {
  if (type == 1) {
    activeIndex.value = index;
    form.tag = tag;
    // console.log("单选",form.tag);
  }
  if (type == 2) {
    activeIndex.value = index;
    // console.log("自定义添加activeIndex",this.activeIndex);
  }
}


type checkTag = {
  name: string,
  resourceName: string,
  checked: boolean
}

function useForm() {
  // 删除标注
  const deleteAnnotation = (row: any) => {
    let annotaion: Annotation = {
      description: row.description,
      segmentation: row.segmentation,
      tag: row.tag
    }
    deleteAnno(algaStore.algaId, annotaion).then((res) => {
      console.log("deleteAnno", res)
      if (res.data == true) {
        console.log("删除成功")
        emit('update')
      }
    })
  }
  const formats: Tag[] = reactive([])
  const getTags = () => {
    getAllTags().then((res) => {
      let tags = Array.from(res.data)
      tags.forEach((value: any) => {
        let tag: Tag = {
          name: value.name,
          resourceName: value.resourceName
        }
        tagStore.$patch((state) => {
          state.tags.push(tag)
        })
        formats.push(tag)
      })
    })
  }
  getTags()
  const form = reactive({
    user: props.user,
    algaId: props.alga.id,
    description: '',
    segmentation: [] as number[],
    tag: {
      name: '',
      resourceName: ''
    },
  })
  const submitForm = () => {
    form.algaId = props.alga.id
    form.user = props.user
    console.log("form", form)
    if (form.algaId == "" || form.tag.name == "" || form.segmentation.length == 0) {
      ElMessage.warning("请补充完整信息")
      return
    }
    addAnno(form.algaId, form.segmentation, form.tag, form.description).then((res) => {
      if (res.code != 200) {
        console.log(res)
        return
      }
      ElMessage.success("添加成功")
      cancelForm()
      location.reload()
    })
    // ElMessage.success("添加成功")
    // cancelForm()
  }
  const cancelForm = () => {
    form.description = ""
    form.tag.name = ""
    form.tag.resourceName = ""
    emit('update', false)
  }
  // 上传文件
  const uploadUrl = import.meta.env.VITE_APP_BASE_API + "/upload"
  const afterUpload: UploadProps['onSuccess'] = (response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) => {
    if (response.code != 200) {
      ElMessage.error("上传失败")
      return
    }
    form.description = response.data.url
  }
  // 关闭回调
  const handleClose = () => {
    cancelForm()
  }

  const handleTagClose = (tag: Tag) => {
    deleteTag(tag).then((res) => {
      if (res.code != 200) {
        return
      }
      ElMessage.success("删除成功")
      cancelForm()
      location.reload()
    })
  }

  const handleTagClick = (tag: Tag) => {
    form.tag = tag
  }
  return {
    deleteAnnotation,
    formats,
    form,
    handleTagClose,
    handleTagClick,
    submitForm,
    cancelForm,
    uploadUrl,
    afterUpload,
    handleClose
  }
}

const {
  deleteAnnotation,
  formats: tagset,
  form,
  handleTagClose,
  handleTagClick,
  submitForm,
  cancelForm,
  uploadUrl,
  afterUpload,
  handleClose
} = useForm()

const searchTag = () => {
  if (searchTagText.value == "") {
    return
  }
  tagset.length = 0
  tagStore.tags.forEach((value: Tag) => {
    if (value.name.includes(searchTagText.value)) {
      tagset.push(value)
    }
  })
}



</script>

<style scoped lang="less">
.wrapper {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100vw;
  height: 600px;
  padding: 24px;
  background-color: #000;
  box-sizing: border-box;
}

.submit {
  width: 500px;
}

.znlh-tab-elect {
  width: 500px;

  // float: left;
  .el-tag--plain {
    color: black;
    border-color: #efefef;
    margin-right: 10px;
  }

  .active {
    color: #3e79ff;
    background-color: #ecf5ff;
    border: 1px solid #9eb5ed;
  }
}
</style>