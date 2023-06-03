<template>
    <el-row>
        <el-button size="large" @click="dialogFormVisiable=true">
            新建数据集
        </el-button>
        <el-col :span="6" style="float:left; margin-left: 45vw;width: 270px; hight: 40px;">
            <el-input type="text" :prefix-icon="Search" v-model="searchKey" placeholder="搜索数据集"
                style="width: 270px; cursor: pointer" @keyup.enter.native="searchRiver(searchKey)"></el-input>
        </el-col>
    </el-row>
    <el-table :data="tableData" style="width: 100%">
        <el-table-column fixed prop="name" label="数据集名称" width="200" />
        <el-table-column prop="address" label="描述" />
        <!--        TODO 待通过pinia 实现数量统计-->
        <!--        <el-table-column prop="size" label="已标记/图片"/>-->
        <el-table-column fixed="right" label="操作" width="150">
            <template #default="scope">
                <el-button link type="primary" size="small" @click="showAlgae(scope.row.id)">标注</el-button>
                <el-button link type="primary" size="small" @click="shareDataSet(scope.row.id as string)">共享</el-button>
                <el-button link type="primary" size="small">删除</el-button>
            </template>
        </el-table-column>
    </el-table>

    <NewDatasetDialog v-model="dialogFormVisiable" @dialog-hide="dialogFormVisiable = false" :userEmail="user.email" />
    <ShareDatasetDialog v-model="shareDialogVisiable" @dialog-hide="shareDialogVisiable=false"></ShareDatasetDialog>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue'
import { useOperatorStore } from "~/store/operator";
import { RiverInfo, getRiverInfo, getRivers } from "~/api/algae";
import { useRouter } from "vue-router";
import { searchRiverByKey } from '~/api/algae';
import { Search } from '@element-plus/icons-vue'
import { useShareStore } from '~/store/shareDatasetInfo';


const user = useOperatorStore()
const router = useRouter()

//控制新建数据集对话框的显示
const dialogFormVisiable = ref(false)

let tableData = reactive([]) as RiverInfo[]
const userEmail = sessionStorage.getItem("UserEmail") as string
const getAllRiverInfo = (userEmail: string) => {
    getRivers(userEmail).then((res) => {
        let rivers = Array.from(res.data)
        rivers.forEach((value: any) => {
            let riverInfo: RiverInfo = {
                id: value.id,
                name: value.name,
                address: value.address
            }
            tableData.push(riverInfo)
            user.dataSet.push(riverInfo)
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

const searchKey = ref("")
const searchRiver = (key: string) => {
    tableData.length = 0
    if (key.length <= 0) {
        getAllRiverInfo(userEmail)
        return
    }
    searchRiverByKey(user.email, key).then((res) => {
        let rivers = Array.from(res.data)
        rivers.forEach((value: any) => {
            let riverInfo: RiverInfo = {
                id: value.id,
                name: value.name,
                address: value.address
            }
            tableData.push(riverInfo)
        })
    })
}
const shareRiverInfo  = useShareStore()
const shareDialogVisiable = ref(false)
const shareDataSet= (riverId:string) => {
    shareRiverInfo.$patch(
        {
            riverId:riverId,
            userEmail:''
    })  
    console.log("share",riverId,"ing")
     shareDialogVisiable.value = true
}



</script>
