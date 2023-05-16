<template>
    <el-drawer :model-value="drawer" :title="alga.name" size="50%" :before-close="handleClose" destroy-on-close>
        <el-table :data="gridData" height="250">
            <el-table-column property="createAt" label="创建时间"/>
            <el-table-column property="updateAt" label="修改时间"/>
            <el-table-column property="description" label="描述"/>
            <el-table-column label="标签">
                <el-table-column prop="tag.name" label="名称"/>
                <el-table-column prop="tag.resourceName" label="resourceName" width="120"/>
            </el-table-column>
            <el-table-column label="操作">
                <template #default="scope">
                    <el-button size="small" type="success" @click="deleteAnnotation(scope.row)">删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <br>
        <el-collapse>
            <el-collapse-item name="1">
                <template #title>
                    <h3>添加标注</h3>
                </template>
                <el-form :model="form">
                    <el-form-item label="描述">
                        <el-input v-model="form.description" autosize type="textarea" clearable
                                  placeholder="请输入标注描述"/>
                    </el-form-item>
                    <el-form-item label="标签">
<!--                        <el-select v-model="form.tag">
                            <el-option v-for="item in formats" :key="item" :label="item" :value="item"/>
                        </el-select>-->
                         <el-tag
                v-for="(item, index) in formats"
                :key="index"
                closable
                effect="plain"
                @click="handleRadio(2, item, index)"
                :class="{ active: item.length + index == activeIndex }"
                @close="handleClose(item)"
              >
                {{item.name.length>10?item.name.substring(0,10)+'...':item.name}}
              </el-tag>
                    </el-form-item>
                    <el-form-item label="添加标注">
                        <el-upload drag :action="uploadUrl" multiple :on-success="afterUpload">
                            <el-icon class="el-icon--upload">
                                <i-ep-upload-filled/>
                            </el-icon>
                            <div class="el-upload__text">
                                Drop file here or <em>click to upload</em>
                            </div>
                        </el-upload>
                    </el-form-item>
                </el-form>
                <el-button size="large" @click="cancelForm">取消</el-button>
                <el-button type="primary" size="large" @click="submitForm">提交</el-button>
            </el-collapse-item>
        </el-collapse>
    </el-drawer>
</template>

<script lang="ts" setup>
import {addAnno, deleteAnno} from "~/api/algae"
import {UploadFile, UploadFiles, UploadProps} from "element-plus/es"
import TagSets from "~/views/TagSets.vue";
import {Annotation} from "~/api/algae";
import {getAllTags} from "~/api/tag";
const props = defineProps(['alga', 'drawer', 'user', 'gridData'])
const emit = defineEmits(['update'])
type   Tag ={
}
function useForm() {
    // 删除标注
    const deleteAnnotation = (row: any) => {

    }
    // 添加标注
    const formats:any= reactive([])
    const getTags = () =>{
        getAllTags().then((res) =>{
            let tags = Array.from(res.data)
            tags.forEach((value: any) => {
                formats.push({
                    name: value.name,
                    resourceName:value.resourceName
                })
            })
        })
    }
    getTags()
    const form = reactive({
        user: props.user,
        algaId: props.alga.id,
        description: '',
        tag: {
            name:'',
            resourceName:''
        },
    })
    const submitForm = () => {
        form.algaId = props.alga.id
        form.user = props.user
        if (form.user == "" || form.algaId == "" || form.description == ""  || form.tag.name == "") {
            ElMessage.warning("请补充完整信息")
            return
        }
        addAnno(form.algaId,form.tag,form.description).then((res) => {
            if (res.code != 200) {
                return
            }
            ElMessage.success("添加成功")
            cancelForm()
            location.reload()
        })
    }
    const cancelForm = () => {
        form.description = ""
        form.tag.name =""
        form.tag.resourceName = ""
        emit('update', false)
    }
    // 上传文件
    const uploadUrl = import.meta.env.VITE_APP_BASE_API + "/upload"
    const afterUpload: UploadProps['onSuccess'] = (response: any, uploadFile: UploadFile, uploadFiles: UploadFiles) => {
        if (response.code != 200) {
            ElMessage.error("上传失败")
            return
        }
        form.description = response.data.url
    }
    // 关闭回调
    const handleClose = (done: any) => {
        cancelForm()
        done()
    }
    return {
        deleteAnnotation,
        formats,
        form,
        submitForm,
        cancelForm,
        uploadUrl,
        afterUpload,
        handleClose
    }
}

const {deleteAnnotation, formats, form, submitForm, cancelForm, uploadUrl, afterUpload, handleClose} = useForm()
</script>

<style scoped>
.submit {
    width: 500px;
}
</style>