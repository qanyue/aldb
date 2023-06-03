<template>
  <el-upload   class="avatar-uploader" action=""  :show-file-list="true" :http-request="uploadImg" multiple
    list-type="picture"
    :before-upload="beforeUpload">
    <el-icon class="avatar-uploader-icon">
    <Plus /></el-icon>
  </el-upload>
</template>

<script lang="ts"  setup>
import { ref } from 'vue'
import type { UploadFile, UploadFiles, UploadProps, UploadRequestHandler } from 'element-plus'
import axios from 'axios'
import { Plus, Upload } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const imageUrl = ref('')
const imageName = ref('')
// 父组件传值时，须有图片的url；其次可选择图片的宽高（默认都为180）
const props = defineProps({
  url: String,
  width: {
    type: Number,
    default: 180
  },
  height: {
    type: Number,
    default: 180
  }
})
const emit = defineEmits(['upload'])


const beforeUpload = (rawFile) => {
  if (rawFile.type !== 'image/jpg' && rawFile.type !== 'image/png' && rawFile.type !== 'image/jpeg') {
    // console.log(rawFile.type)
    ElMessage.error('图片格式应该是png或jpg')
    return false
  } else if (rawFile.size / 1024 / 1024 > 5) {
    ElMessage.error('图片大小应该小于5MB')
    return false
  }
  return true
}

/**
 * 上传图片至华为obs
 * @param {*} req
 */
const afterUpload: UploadProps['onSuccess'] = (response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) => {
  if (response.code != 200) {
    ElMessage.error("上传失败")
    return
  }
}
const uploadImg = function (req) {
  const config = {
    headers: { 'Content-Type': 'multipart/form-data' }
  }
  // 重命名要上传的文件
  const uploadUrl = import.meta.env.VITE_APP_BASE_API + "/upload"
  const formdata = new FormData()
  formdata.append('file', req.file)
   axios.post(uploadUrl, formdata, config).then(res => {
    if (res.data.code != 200) {
      ElMessage.error("上传图片失败")
      return
    }
    imageUrl.value = res.data.data.url as string
    imageName.value = res.data.data.name as string
    console.log("Image:infomation "+imageName.value,imageUrl.value)
    emit('upload',  imageName.value,imageUrl.value) // 向父组件传递图片的url
  })
}
</script>

<style lang="scss" scoped>
.avatar-uploader .avatar {
  width: 360px;
  height: 180px;
  display: block;
}

.avatar-uploader :deep(.el-upload) {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader :deep(.el-upload:hover) {
  border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
}
</style>
