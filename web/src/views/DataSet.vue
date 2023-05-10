<template>
    <el-table :data="tableData" style="width: 100%">
        <el-table-column fixed prop="name" label="数据集名称" width="200"/>
        <el-table-column prop="address" label="描述"/>
        <!--        TODO 待通过pinia 实现数量统计-->
        <el-table-column prop="size" label="已标记/图片"/>
        <el-table-column fixed="right" label="操作" width="150">
            <template #default>
                <el-button link type="primary" size="small" @click="handleClick">统计</el-button>
                <el-button link type="primary" size="small">标注</el-button>
                <el-button link type="primary" size="small">删除</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>

<script lang="ts" setup>
const handleClick = () => {
    console.log('click')
}

const tableData = [
    {
        date: '2016-05-03',
        name: '数据集1',
        state: 'California',
        size: "20/100",
        city: 'Los Angeles',
        address: '数据集1描述',
        zip: 'CA 90036',
        tag: 'Home',
    },
    {
        date: '2016-05-02',
        name: '数据集2',
        state: 'California',
        city: 'Los Angeles',
        address: '数据集2描述',
        size: "10/20",
        zip: 'CA 90036',
        tag: 'Office',
    },
    {
        date: '2016-05-04',
        name: '数据集3',
        state: 'California',
        city: 'Los Angeles',
        address: '数据集3描述',
        size: "30/100",
        zip: 'CA 90036',
        tag: 'Home',
    },
    {
        date: '2016-05-01',
        name: '数据集4',
        state: 'California',
        city: 'Los Angeles',
        address: '数据集4描述',
        size: "60/120",
        zip: 'CA 90036',
        tag: 'Office',
    },
]
</script>
