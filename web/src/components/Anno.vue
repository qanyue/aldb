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
                    <el-input v-model="annoInfo.anno.description" />
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
import { ElMessage, ElMessageBox } from "element-plus"
import { UploadFile, UploadFiles, UploadProps } from "element-plus/es"
import { deleteAnno, addAnno, getAnno, Annotation } from "~/api/algae"
import { useAnnotationStore } from "~/store/anno";

const props = defineProps(['algaId'])

const tableData = ref([])
const formats = ref(['XML', 'COCO', 'VOC', 'CSV', 'JSON'])
const annoDialog = ref(false)
const annoInfo = useAnnotationStore()


function useAnnoForm() {
    // 获取图片标注
    const fetchAnno = () => {
        getAnno(props.algaId).then((res) => {
            tableData.value = res.data
        })
    }
    fetchAnno()
    // 编辑
    const editButton = (anno: Annotation) => {
        annoInfo.anno.description = anno.description
        annoDialog.value = true
    }
    const editSubmit = () => {
        addAnno(annoInfo.algaId, annoInfo.anno).then((res) => {
            if (res.code != 200) {
                return
            }
            ElMessage.success("修改成功")
            annoDialog.value = false
            fetchAnno()
        })
    }
    // 删除
    const deleteButton = (algaId: string, anno: Annotation) => {
        ElMessageBox.confirm(
            '确定删除此标注？',
            //TODO 加入图片名称
            '删除......',
            {
                confirmButtonText: '确认',
                cancelButtonText: '取消',
                type: 'warning',
            }
        )
            .then(() => {
                deleteAnno(algaId, anno).then((res) => {
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
    // TODO  添加标注界面
    const download = (row: any) => {
        window.open(row)
    }
    const uploadUrl = import.meta.env.VITE_APP_BASE_API + "/upload"
    //TODO 修改为添加成功记录
    const afterUpload: UploadProps['onSuccess'] = (response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) => {
        if (response.code != 200) {
            ElMessage.error("上传失败")
            return
        }
        ElMessage.success("上传成功")
        // annoInfo.url = response.data.url
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

<style scoped></style>