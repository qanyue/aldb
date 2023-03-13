<template>
  <div>
    <div>
      <el-table :data="tableData" :default-sort="{ prop: 'createAt', order: 'descending' }">
        <el-table-column align="left" label="创建时间" prop="createAt" sortable />
        <el-table-column align="left" label="用户名" prop="name" />
        <el-table-column align="left" label="邮箱" prop="email" />
        <el-table-column align="left" label="用户角色">
          <template #default="scope">
            <template v-if="scope.row.access === 0">普通用户</template>
            <template v-else>管理员</template>
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
    <el-dialog v-model="userDialog" title="编辑用户" center destroy-on-close>
      <el-form ref="userForm" :model="userInfo" label-width="80px">
        <el-form-item label="用户名" prop="name">
          <el-input v-model="userInfo.name" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="userInfo.password" type="password" />
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="userInfo.email" disabled />
        </el-form-item>
        <el-form-item label="用户角色" prop="access">
          <el-select v-model="userRole" placeholder="Select">
            <el-option key="用户" value="用户" label="用户" />
            <el-option key="管理员" value="管理员" label="管理员" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="userDialog = false">取消</el-button>
          <el-button type="primary" @click="editSubmit">确定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { Delete, Edit } from '@element-plus/icons-vue'
import { deleteUser, getUsers, updateUser } from "~/api/user"
import { ElMessageBox } from "element-plus"
const tableData = ref([])
function useInfo() {
  // 获取用户信息
  const fetchUsers = () => {
    getUsers().then((res) => {
      if (res.code != 200) {
        return
      }
      tableData.value = res.data
    })
  }
  return { tableData, fetchUsers }
}
const { fetchUsers } = useInfo()
fetchUsers()
const userDialog = ref(false)
const userRole = ref('')
const userInfo = reactive({
  name: '',
  password: '',
  email: '',
  access: 0,
})
function useControl() {
  const editButton = (user: any) => {
    userInfo.name = user.name;
    userInfo.password = user.password;
    userInfo.email = user.email
    userRole.value = '用户'
    userDialog.value = true;
  }
  const editSubmit = () => {
    if (userRole.value == '用户') userInfo.access = 0
    else userInfo.access = 1
    updateUser(userInfo).then((res) => {
      if (res.code != 200) {
        return
      }
      ElMessage.success("修改成功")
      userDialog.value = false
      fetchUsers()
    })
  }
  const deleteButton = (user: any) => {
    ElMessageBox.confirm(
      '确定删除此用户？',
      '删除 ' + user.name,
      {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        type: 'warning',
      }
    )
      .then(() => {
        deleteUser(user.email).then((res) => {
          if (res.code == 200)
            ElMessage.success("删除成功")
          else ElMessage.error("删除失败")
          fetchUsers()
        })
      })
      .catch(() => {
        ElMessage.info("删除撤销")
      })
  }
  return { editButton, editSubmit, deleteButton }
}
const { editButton, editSubmit, deleteButton } = useControl()
</script>

<style lang="scss">
.nickName {
  display: flex;
  justify-content: flex-start;
  align-items: center;
}

.pointer {
  cursor: pointer;
  font-size: 16px;
  margin-left: 2px;
}
</style>