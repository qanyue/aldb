<template>
  <el-drawer :model-value="drawer" :title="alga.name" size="50%" :before-close="handleClose" destroy-on-close>
    <el-table :data="gridData" height="250">
      <el-table-column property="createAt" label="创建时间" />
      <el-table-column property="updateAt" label="修改时间" />
      <el-table-column property="description" label="描述" />
      <el-table-column property="format" label="格式" />
      <el-table-column label="操作">
        <template #default="scope">
          <el-button size="small" type="success" @click="download(scope.row.url)">下载
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <br>
    <el-collapse>
      <el-collapse-item name="1">
        <template #title>
          <h3>添加标注</h3>
        </template>
        <el-form :model="form">
          <el-form-item label="描述">
            <el-input v-model="form.description" autosize type="textarea" clearable placeholder="请输入标注描述" />
          </el-form-item>
          <el-form-item label="格式">
            <el-select v-model="form.format">
              <el-option v-for="item in formats" :key="item" :label="item" :value="item" />
            </el-select>
          </el-form-item>
          <el-form-item label="上传标注">
            <el-upload drag :action="uploadUrl" multiple :on-success="afterUpload">
              <el-icon class="el-icon--upload">
                <i-ep-upload-filled />
              </el-icon>
              <div class="el-upload__text">
                Drop file here or <em>click to upload</em>
              </div>
            </el-upload>
          </el-form-item>
        </el-form>
        <el-button size="large" @click="cancelForm">取消</el-button>
        <el-button type="primary" size="large" @click="submitForm">提交</el-button>
      </el-collapse-item>
    </el-collapse>
  </el-drawer>
</template>

<script lang="ts" setup>
import { addAnno } from "~/api/algae"
import { UploadFile, UploadFiles, UploadProps } from "element-plus/es"

const props = defineProps(['alga', 'drawer', 'user', 'gridData'])
const emit = defineEmits(['update'])

function useForm() {
  // 下载标注
  const download = (row: any) => {
    window.open(row)
  }
  // 添加标注
  const formats = ref(['XML', 'COCO', 'VOC', 'CSV', 'JSON'])
  const form = reactive({
    user: props.user,
    alga: props.alga.name,
    description: '',
    format: '',
    url: '',
  })
  const submitForm = () => {
    form.alga = props.alga.name
    form.user = props.user
    if (form.user == "" || form.alga == "" || form.description == "" || form.format == "" || form.url == "") {
      ElMessage.warning("请补充完整信息")
      return
    }
    addAnno(form).then((res) => {
      if (res.code != 200) {
        return
      }
      ElMessage.success("添加成功")
      cancelForm()
      location.reload()
    })
  }
  const cancelForm = () => {
    form.description = ""
    form.format = ""
    form.url = ""
    emit('update', false)
  }
  // 上传文件
  const uploadUrl = import.meta.env.VITE_APP_BASE_API + "/upload"
  const afterUpload: UploadProps['onSuccess'] = (response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) => {
    if (response.code != 200) {
      ElMessage.error("上传失败")
      return
    }
    form.url = response.data.url
  }
  // 关闭回调
  const handleClose = (done: any) => {
    cancelForm()
    done()
  }
  return {
    download,
    formats,
    form,
    submitForm,
    cancelForm,
    uploadUrl,
    afterUpload,
    handleClose
  }
}
const { download, formats, form, submitForm, cancelForm, uploadUrl, afterUpload, handleClose } = useForm()
</script>

<style scoped>
.submit {
  width: 500px;
}
</style>