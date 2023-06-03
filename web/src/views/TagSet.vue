<template>
    <el-row>
        <el-button size="large" @click="dialogFormVisiable = true">新建标签</el-button>
        <el-col :span="6" style="float:left; margin-left: 45vw;width: 270px; hight: 40px;">
            <el-input type="text" :prefix-icon="Search" v-model="searchTagText" placeholder="请输入"
                style="width: 270px; cursor: pointer" @input="searchTag"></el-input>
        </el-col>
    </el-row>

    <el-table :data="tagset" style="width: 100%">
        <el-table-column fixed prop="name" label="name" width="200" />
        <el-table-column prop="resourceName" label="resourceName" />
        <el-table-column fixed="right" label="操作" width="150">
            <template #default="scope">
                <el-button link type="primary" size="small" @click="deleteSelectTag(scope.row)">删除</el-button>
            </template>
        </el-table-column>
    </el-table>
    <NewTagsetDialog v-model="dialogFormVisiable" @dialog-hide="dialogFormVisiable = false"></NewTagsetDialog>
</template>

<script lang="ts" setup>
import { deleteTag, getAllTags } from "~/api/tag";
import { useTagStore } from "~/store/tagSet";
import { Search } from '@element-plus/icons-vue'
import { Tag } from "~/api/algae";


let tagset: Tag[] = reactive([])
const tagStore = useTagStore()
const searchTagText = ref('')

//控制新建标签集对话框的显示
const dialogFormVisiable = ref(false)
const TagSearchText = ref('')
const getAllTag = () => {
    getAllTags().then((res) => {
        let Tags = Array.from(res.data)
        Tags.forEach((value: any) => {
            tagset.push({
                name: value.name,
                resourceName: value.resourceName
            }as Tag)
        })
        if (tagStore.tags.length <= 0) {
            tagStore.tags = JSON.parse(JSON.stringify(tagset));
        }
    })
}
getAllTag()
const searchTag = () => {
    if (searchTagText.value == "") {
        tagStore.tags.forEach((value: Tag) => {
            tagset.push(value)
        })
        return
    }
    tagset.length = 0
    tagStore.tags.forEach((value: Tag) => {
        if (value.name.includes(searchTagText.value) || value.resourceName.includes(searchTagText.value)) {
            tagset.push(value)
        }
    })
}

const deleteSelectTag =  (tag: Tag)=>{
    deleteTag(tag).then((res) => {
        if (res.code != 200) {
            ElMessage.error("删除标签失败")
            return
        }
        ElMessage.success("删除标签成功")
        //重新加载标签集
        getAllTag()
    })
}
</script>