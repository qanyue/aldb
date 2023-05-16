<template>
    <el-table :data="tableData" style="width: 100%">
        <el-table-column fixed prop="name" label="数据集名称" width="200"/>
        <el-table-column prop="address" label="描述"/>
        <!--        TODO 待通过pinia 实现数量统计-->
        <!--        <el-table-column prop="size" label="已标记/图片"/>-->
        <el-table-column fixed="right" label="操作" width="150">
            <template #default="scope">
                <el-button link type="primary" size="small" @click="showAlgae(scope.row.id)">标注</el-button>
                <el-button link type="primary" size="small">删除</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue'
import {useOperatorStore} from "~/store/operator";
import {getRiverInfo, getRivers} from "~/api/algae";
import {emitChangeFn} from "element-plus";
import {useRouter} from "vue-router";

const router = useRouter()
type riverWithoutAlgae = {
    id: string
    name: string
    address: string

}
const handleClick = (hello: string) => {
    console.log('click' + hello)
}

let tableData: any = reactive([])
const userEmail = sessionStorage.getItem("UserEmail") as string
const operater = useOperatorStore()
const getAllRiverInfo = (userEmail: string) => {
    getRivers(userEmail).then((res) => {
        let rivers = Array.from(res.data)
        rivers.forEach((value: any) => {
            tableData.push({
                name: value.name,
                address: value.address,
                id: value.id
            })
        })
    })
}
getAllRiverInfo(userEmail)
const showAlgae = (riverId: string) => {
    router.push({
        name: "Main",
        params: {
            riverId: riverId
        }
    })
}

</script>
