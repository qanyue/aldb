<template>
    <el-dialog title="新建标签">
        <el-form :model="tagForm">
            <el-form-item label="Name" :label-width="formLabelWidth">
                <el-input v-model="tagForm.name" />
            </el-form-item>
            <el-form-item label="ResourceName" :label-width="formLabelWidth">
                <el-input v-model="tagForm.resourceName" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="$emit('dialog-hide')">取消</el-button>
                <el-button type="primary" @click="submitTagForm(tagForm)">提交</el-button>
            </span>
        </template>
    </el-dialog>

</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { Tag} from '~/api/algae'
import { addTag } from '~/api/tag'
defineEmits(['dialog-hide'])
const formLabelWidth = '140px'
const tagForm = reactive({
    name: '',
    resourceName: ''
}) as Tag

const submitTagForm = (tag: Tag) => {
    addTag(tag).then((res) => {
        if (res.code != 200) {
            ElMessage.error(res.msg)
            return
        }
        ElMessage.success('添加标签成功')
        tagForm.name = ''
        tagForm.resourceName = ''
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
