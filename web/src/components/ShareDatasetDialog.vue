<template>
    <el-dialog title="共享数据集">
        <el-form :model="riverForm">
            <el-form-item label="用户邮箱" :label-width="formLabelWidth">
                <el-input v-model="riverForm.userEmail" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="$emit('dialog-hide')">取消</el-button>
                <el-button type="primary" @click="submitRiverForm(riverForm.riverId,riverForm.userEmail)">提交</el-button>
            </span>
        </template>
    </el-dialog>

</template>

<script lang="ts" setup>
import { ElMessage } from 'element-plus'
import { reactive, ref } from 'vue'
import { shareRiver } from '~/api/algae'
import { useShareStore } from '~/store/shareDatasetInfo'
defineEmits(['dialog-hide'])
const formLabelWidth = '140px'
const riverForm = reactive({
    riverId: '',
    userEmail: ''
})

const shareDatasetInfo = useShareStore()
riverForm.riverId= shareDatasetInfo.riverId
const submitRiverForm = (riverId:string,userEmail:string) => {
    shareRiver(userEmail,riverId).then((res) => {
        if (res.code != 200) {
            ElMessage.error(res.msg)
            return
        }
        ElMessage.success('共享数据集成功')
        riverForm.riverId = ''
        riverForm.userEmail = ''
    })
}

</script>
<style scoped>
.el-button--text {
    margin-right: 15px;
}

.el-select {
    width: 300px;
}

.el-input {
    width: 300px;
}

.dialog-footer button:first-child {
    margin-right: 10px;
}
</style>
