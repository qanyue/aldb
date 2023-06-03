<template>
  <el-button circle @click="dialogFormVisible = true" :icon="Upload" plain />
  <el-dialog v-model="dialogFormVisible" title="藻类图像上传" width="30%" destroy-on-close>
    <el-form :model="form">
      <el-form-item label="前缀">
        <el-input v-model="form.preSuffix" />
      </el-form-item>
   
      <el-form-item label="图像">
        <!--将图像上传后会返回OBS存储链接-->        
        <UploadDialog @upload="addAlgaForm" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="formCancel">取消</el-button>
        <el-button type="primary" @click="formSubmit">确认</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import {  addAlgaMore } from "~/api/algae"
import type { UploadFile, UploadFiles, UploadProps } from 'element-plus'
import { Plus, Upload } from '@element-plus/icons-vue'
import UploadDialog from "./UploadDialog.vue";
import { useRoute } from "vue-router";
const props = defineProps({
  riverId:{type:String,required:true}
  }
)
const emit = defineEmits(['refresh'])
const dialogFormVisible = ref(false)
const riverId  = ref(props.riverId as string)
const form = reactive({
  preSuffix: "",
  src: "",
})

class algaForm {
  preSuffix: string = ""
  src: string = ""
  constructor(name: string, src: string) {
    this.preSuffix = name
    this.src = src
  }
  toJson() {
    return {
      name: this.preSuffix,
      src: this.src,
    }
  }

}
let algaForms = new Array<algaForm>

const clearForm = () => {
  form.preSuffix = ""
  form.src = ""
}

function useUpload() {

  const formCancel = () => {
    dialogFormVisible.value = false
    clearForm();
  }

  const addAlgaForm = (name:string, url:string) => {
    console.log(name,url)
    algaForms.push(new algaForm(name, url))
  }


  const formSubmit = () => {
    let formdata = new FormData()
    algaForms.forEach((value) => {
      formdata.append("algaforms", JSON.stringify(value))
    })
    addAlgaMore(formdata,riverId.value,form.preSuffix).then((res) => {
      if (res.code != 200) {
        ElMessage.error("上传失败")
        return
      }
      ElMessage.success("上传成功")
      formCancel() //清除已经上传的表单,使得表单为空
      emit('refresh', true)
    })
  }
  const afterUpload: UploadProps['onSuccess'] = (response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) => {
    if (response.code != 200) {
      ElMessage.error("上传失败")
      return
    }
    form.src = response.data.url
  }
  return {
    formCancel, addAlgaForm, formSubmit, afterUpload
  }
}
const { formCancel, addAlgaForm, formSubmit, afterUpload } = useUpload()
</script>

<style scoped>
.avatar-uploader .avatar {
  width: 100%;
  height: 100%;
  display: block;
}
</style>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
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
