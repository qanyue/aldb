<template>
    <el-dialog title="新建数据集">
        <el-form :model="riverForm">
            <el-form-item label="名称" :label-width="formLabelWidth">
                <el-input v-model="riverForm.name" />
            </el-form-item>
            <el-form-item label="描述" :label-width="formLabelWidth">
                <el-input v-model="riverForm.address" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="$emit('dialog-hide')">取消</el-button>
                <el-button type="primary" @click="submitRiverForm(riverForm)">提交</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script lang="ts" setup>
import { ElMessage } from 'element-plus'
import { reactive, ref } from 'vue'
import { addRiver } from '~/api/algae'
import { River } from '~/api/algae'

defineEmits(['dialog-hide'])

const props = defineProps({
    userEmail: { type: String, required: true }
})
const formLabelWidth = '140px'

const riverForm = reactive({
    name: '',
    address: ''
}) as River

const submitRiverForm = (river: River) => {
    let userEmail = props.userEmail as string
    if (userEmail.length <= 0) {
        ElMessage.error("请先登录")
        return
    }
    addRiver(userEmail, river).then((res) => {
        if (res.code != 200) {
            ElMessage.error(res.msg)
            return
        }
        ElMessage.success('添加数据集成功')
        riverForm.name = ''
        riverForm.address = ''
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
