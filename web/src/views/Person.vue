<template>
  <Header :user-name="userInfo.name" :display="false"/>
  <el-main>
    <el-row>
      <el-col :span="6">
        <div class="fl-left avatar-box">
          <div class="user-card">
            <div class="user-headpic-update"></div>
            <div class="user-personality">
              <p class="nickName">{{ userInfo.name }}</p>
              <p class="person-info">欢迎使用本系统</p>
            </div>
            <div class="user-information">
              <ul>
                <li>
                  <i-ep-user />
                  {{ role }}
                </li>
                <el-tooltip class="item" effect="light" :content="userInfo.createAt" placement="bottom">
                  <li>
                    <i-ep-timer />
                    注册时间
                  </li>
                </el-tooltip>
              </ul>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :span="18">
        <div class="user-addcount">
          <el-tabs v-model="activeName">
            <el-tab-pane v-if="perm" label="用户管理" name="control">
              <Grant />
            </el-tab-pane>
            <el-tab-pane v-else label="标注管理" name="control">
              <Anno :user-email="userEmail" />
            </el-tab-pane>
            <el-tab-pane label="账号管理" name="info">
              <ul>
                <li>
                  <p class="title">绑定邮箱</p>
                  <p class="desc">
                    已绑定邮箱：{{ userInfo.email }}
                  </p>
                </li>
                <li>
                  <p class="title">修改密码</p>
                  <p class="desc">
                    修改个人密码
                    <a href="javascript:void(0)" @click="showPassword = true">修改密码</a>
                  </p>
                </li>
              </ul>
            </el-tab-pane>
          </el-tabs>
        </div>
      </el-col>
    </el-row>
    <el-dialog v-model="showPassword" title="修改密码" width="360px" @close="">
      <el-form :model="pwdModify" label-width="80px" :rules="rules">
        <el-form-item :minlength="5" label="原密码" prop="password">
          <el-input v-model="pwdModify.password" show-password />
        </el-form-item>
        <el-form-item :minlength="5" label="新密码" prop="newPassword">
          <el-input v-model="pwdModify.newPassword" show-password />
        </el-form-item>
        <el-form-item :minlength="5" label="确认密码" prop="confirmPassword">
          <el-input v-model="pwdModify.confirmPassword" show-password />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="formCancel">取 消
          </el-button>
          <el-button size="small" type="primary" @click="formSubmit">确 定
          </el-button>
        </div>
      </template>
    </el-dialog>
  </el-main>
</template>

<script lang="ts" setup>
import { changePassword, getUser } from "~/api/user"
import { useRouter } from "vue-router"
import Anno from "~/components/Anno.vue"

const router = useRouter()
const userEmail = sessionStorage.getItem("UserEmail") as string
if (userEmail == null) {
  ElMessage.error("未登录")
  router.push({ name: "Login" })
}
const userInfo = ref({})
const role = ref('普通用户')
const perm = ref(false)
const fetchUser = () => {
  getUser(userEmail).then((res) => {
    userInfo.value = res.data
    if (res.data.access != 0) {
      role.value = '系统管理员'
      perm.value = true
    }
  })
}
fetchUser()
const activeName = ref('control')
// 用户管理
const rules = reactive({
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 5, message: '最少5个字符', trigger: 'blur' },
  ],
  newPassword: [
    { required: true, message: '请输入新密码', trigger: 'blur' },
    { min: 5, message: '最少5个字符', trigger: 'blur' },
  ],
  confirmPassword: [
    { required: true, message: '请输入确认密码', trigger: 'blur' },
    { min: 5, message: '最少5个字符', trigger: 'blur' },
    {
      validator: (rule: any, value: string, callback: any) => {
        if (value !== pwdModify.newPassword) {
          callback(new Error('两次密码不一致'))
        } else {
          callback()
        }
      },
      trigger: 'blur',
    },
  ],
})
const showPassword = ref(false)
const pwdModify = reactive({
  email: "",
  password: "",
  newPassword: "",
  confirmPassword: "",
})
const formCancel = () => {
  showPassword.value = false
  pwdModify.password = ""
  pwdModify.newPassword = ""
  pwdModify.confirmPassword = ""
}
const formSubmit = () => {
  pwdModify.email = userEmail
  if (pwdModify.password == "" || pwdModify.newPassword == "" || pwdModify.confirmPassword == "") {
    ElMessage.warning("请补充完整信息")
    return
  }
  changePassword(pwdModify).then((res) => {
    if (res.code != 200) {
      ElMessage.error("修改失败")
      return
    }
    ElMessage.success("修改成功")
    formCancel()
  })
}
</script>

<style lang="scss">
@import "../styles/person";
</style>