<template>
    <el-table :data="tableData" style="width: 100%">
        <el-table-column fixed prop="name" label="name" width="200"/>
        <el-table-column prop="resourceName" label="resourceName"/>
        <!--        TODO 待通过pinia 实现数量统计-->
        <!--        <el-table-column prop="size" label="已标记/图片"/>-->
        <el-table-column fixed="right" label="操作" width="150">
            <template  #default="scope">
<!--                <el-button link type="primary" size="small" @click="">新建</el-button>-->
                <el-button link type="primary" size="small">删除</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue'
import {useOperatorStore} from "~/store/operator";
import {emitChangeFn} from "element-plus";
import {getRivers} from "~/api/algae";
import {getAllTags} from "~/api/tag";

const handleClick = (hello:string) => {
    console.log('click'+hello)
}

let tableData: any = reactive([])
const userEmail = sessionStorage.getItem("UserEmail") as string
const operater = useOperatorStore()
const getAllTag = () => {
    getAllTags().then((res) => {
        let Tags = Array.from(res.data)
        Tags.forEach((value: any) => {
            tableData.push({
                name: value.name,
                resourceName: value.resourceName
            })
        })
    })
}
getAllTag()
const makeAnnotation = () => {

}

</script>