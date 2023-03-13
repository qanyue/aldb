<template>
  <div>
    <div>
      <el-table :data="tableData" :default-sort="{ prop: 'createAt', order: 'descending' }">
        <el-table-column align="left" label="创建时间" prop="createAt" sortable />
        <el-table-column align="left" label="修改时间" prop="updateAt" />
        <el-table-column align="left" label="描述" prop="description" />
        <el-table-column align="left" label="格式" prop="format" />
        <el-table-column align="left" label="下载">
          <template #default="scope">
            <el-button :icon="Download" type="success" circle @click="download(scope.row.url)" />
          </template>
        </el-table-column>
        <el-table-column label="操作" min-width="250" fixed="right">
          <template #default="scope">
            <el-button type="primary" :icon="Edit" circle @click="editButton(scope.row)" />
            <el-button type="danger" :icon="Delete" circle @click="deleteButton(scope.row)" />
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="annoDialog" title="编辑标注" center destroy-on-close>
      <el-form ref="userForm" :model="annoInfo" label-width="80px">
        <el-form-item label="描述" prop="description">
          <el-input v-model="annoInfo.description" />
        </el-form-item>
        <el-form-item label="格式">
          <el-select v-model="annoInfo.format">
            <el-option v-for="item in formats" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="链接" prop="url">
          <el-input v-model="annoInfo.url" disabled />
        </el-form-item>
        <el-form-item>
          <el-upload class="upload-demo" :action="uploadUrl" :on-success="afterUpload" :show-file-list="false">
            <el-button type="primary">上传新文件</el-button>
          </el-upload>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="annoDialog = false">取消</el-button>
          <el-button type="primary" @click="editSubmit">确定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { Delete, Download, Edit } from '@element-plus/icons-vue'
import { getAnno } from "~/api/user"
import { ElMessageBox } from "element-plus"
import { UploadFile, UploadFiles, UploadProps } from "element-plus/es"
import { deleteAnno, updateAnno } from "~/api/algae"

const props = defineProps(['userEmail'])

const tableData = ref([])
const formats = ref(['XML', 'COCO', 'VOC', 'CSV', 'JSON'])
const annoDialog = ref(false)
const annoInfo = reactive({
  description: '',
  format: '',
  url: '',
  id: '',
})
function useAnnoForm() {
  // 获取个人标注
  const fetchAnno = () => {
    getAnno(props.userEmail).then((res) => {
      tableData.value = res.data
    })
  }
  fetchAnno()
  // 编辑
  const editButton = (anno: any) => {
    annoInfo.description = anno.description
    annoInfo.format = anno.format
    annoInfo.url = anno.url
    annoInfo.id = anno.id
    annoDialog.value = true
  }
  const editSubmit = () => {
    updateAnno(annoInfo).then((res) => {
      if (res.code != 200) {
        return
      }
      ElMessage.success("修改成功")
      annoDialog.value = false
      fetchAnno()
    })
  }
  // 删除
  const deleteButton = (obj: any) => {
    ElMessageBox.confirm(
      '确定删除此标注？',
      '删除 ' + obj.id,
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )
      .then(() => {
        deleteAnno(obj.id).then((res) => {
          if (res.code == 200)
            ElMessage.success("删除成功")
          else ElMessage.error("删除失败")
          fetchAnno()
        })
      })
      .catch(() => {
        ElMessage.info("删除撤销")
      })
  }
  // 下载标注
  const download = (row: any) => {
    window.open(row)
  }
  // 上传文件
  const uploadUrl = import.meta.env.VITE_APP_BASE_API + "/upload"
  const afterUpload: UploadProps['onSuccess'] = (response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) => {
    if (response.code != 200) {
      ElMessage.error("上传失败")
      return
    }
    ElMessage.success("上传成功")
    annoInfo.url = response.data.url
  }
  return {
    editButton,
    editSubmit,
    deleteButton,
    download,
    uploadUrl,
    afterUpload
  }
}
const { editButton, editSubmit, deleteButton, download, uploadUrl, afterUpload } = useAnnoForm()
</script>

<style scoped>
</style>