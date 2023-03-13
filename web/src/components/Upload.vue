<template>
  <el-button circle @click="dialogFormVisible = true" :icon="Upload" plain />
  <el-dialog v-model="dialogFormVisible" title="藻类图像上传" width="30%" destroy-on-close>
    <el-form :model="form">
      <el-form-item label="名称">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="所属河流">
        <el-space>
          <el-select v-model="form.river" placeholder="选择河流">
            <el-option v-for="item in rivers" :key="item.name" :label="item.name" :value="item.name" />
          </el-select>
          <el-button type="primary" size="small" @click="riverFormVisible = true">添加
          </el-button>
        </el-space>
      </el-form-item>
      <el-form-item label="图像">
        <el-upload drag :action="uploadUrl" class="avatar-uploader" :show-file-list="false" :on-success="afterUpload">
          <img v-if="form.src" :src="form.src" class="avatar" />
          <el-icon v-else class="avatar-uploader-icon">
            <Plus />
          </el-icon>
        </el-upload>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="formCancel">取消</el-button>
        <el-button type="primary" @click="formSubmit">确认</el-button>
      </span>
    </template>
  </el-dialog>
  <el-dialog v-model="riverFormVisible" title="添加河流" width="30%" center destroy-on-close>
    <el-form :model="river">
      <el-form-item label="名称">
        <el-input v-model="river.name" />
      </el-form-item>
      <el-form-item label="位置">
        <el-input v-model="river.address" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="riverFormCancel">取消</el-button>
        <el-button type="primary" @click="riverFormSubmit">添加</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import { addAlga, addRiver, getRivers } from "~/api/algae"
import type { UploadFile, UploadFiles, UploadProps } from 'element-plus'
import { Plus, Upload } from '@element-plus/icons-vue'

const emit = defineEmits(['refresh'])

const rivers = ref([])
const river = reactive({
  name: "",
  address: "",
});
const riverFormVisible = ref(false)
function useRiverForm() {
  // 获取河流

  const fetchRivers = () => {
    getRivers({}).then((res) => {
      rivers.value = res.data;
    });
  };
  fetchRivers();
  // 添加河流
  const riverFormSubmit = () => {
    if (river.name == "" || river.address == "") {
      ElMessage.warning("请填写完整信息")
      return;
    }
    addRiver(river).then((res) => {
      console.log(res)
      if (res.code != 200) {
        return;
      }

      ElMessage.success("添加河流成功")
      fetchRivers();
      riverFormVisible.value = false
    });
  };
  const riverFormCancel = () => {
    river.name = ""
    river.address = ""
    riverFormVisible.value = false
  };
  return {
    riverFormSubmit,
    riverFormCancel
  }
}
const { riverFormSubmit, riverFormCancel } = useRiverForm()
const dialogFormVisible = ref(false)
const form = reactive({
  name: "",
  src: "",
  river: "",
})
const clearForm = () => {
  form.name = ""
  form.src = ""
  form.river = ""
}
function useUpload() {
  const formCancel = () => {
    dialogFormVisible.value = false
    clearForm();
  }
  const formSubmit = () => {
    addAlga(form).then((res) => {
      if (res.code != 200) {
        ElMessage.error("上传失败")
        return
      }
      ElMessage.success("上传成功")
      formCancel()
      emit('refresh',true)
    })
  }
  // 上传图片
  const uploadUrl = import.meta.env.VITE_APP_BASE_API + "/upload"
  const afterUpload: UploadProps['onSuccess'] = (response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) => {
    if (response.code != 200) {
      ElMessage.error("上传失败")
      return
    }
    form.src = response.data.url
  }

  return {
    formCancel, formSubmit, uploadUrl, afterUpload
  }
}
const { formCancel, formSubmit, uploadUrl, afterUpload } = useUpload()
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