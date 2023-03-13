<template>
  <div id="userLayout">
    <div class="login_panle">
      <div class="login_panle_form">
        <div class="login_panle_form_title">
          <!-- <img class="login_panle_form_title_logo" :src="" alt /> -->
          <p class="login_panle_form_title_p">数据标注存储系统</p>
        </div>
        <el-form ref="loginForm" :model="loginFormData" :rules="rules" @keyup.enter="submitForm">
          <el-form-item prop="email">
            <el-input v-model="loginFormData.email" placeholder="请输入邮箱">
              <template #suffix>
                <span class="input-icon">
                  <i-ep-user />
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input v-model="loginFormData.password" type="password" placeholder="请输入密码" show-password="true">
              <template #suffix>
                <span class="input-icon">
                  <i-ep-lock />
                </span>
              </template>
            </el-input>
          </el-form-item>
          <el-form-item prop="captcha">
            <div class="vPicBox">
              <el-input v-model="loginFormData.captchaValue" placeholder="请输入验证码" style="width: 60%" />
              <div class="vPic">
                <img v-if="picPath" :src="picPath" alt="请输入验证码" @click="loginVerify()" />
              </div>
            </div>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" style="width: 46%" size="large" @click="dialogFormVisible = true">注册
            </el-button>
            <el-button type="primary" size="large" style="width: 46%; margin-left: 8%" @click="submitForm">登 录
            </el-button>
          </el-form-item>
        </el-form>
      </div>
      <el-dialog v-model="dialogFormVisible" title="注册" center width="30%">
        <el-form :model="form" :rules="rules">
          <el-form-item label="姓名">
            <el-input v-model="form.name" />
          </el-form-item>
          <el-form-item label="密码" prop="password">
            <el-input v-model="form.password" type="password" show-password="true" />
          </el-form-item>
          <el-form-item label="邮箱" prop="email">
            <el-input v-model="form.email" />
          </el-form-item>
        </el-form>
        <template #footer>
          <el-button @click="cancelSubmit">取消
          </el-button>
          <el-button type="primary" @click="registerSubmit">确定
          </el-button>
        </template>
      </el-dialog>
      <div class="login_panle_right" />
      <div class="login_panle_foot">
        <div class="copyright">
          <CopyRight />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { getCaptcha, loginSubmit } from "~/api/auth"
import { registerUser } from "~/api/user"
import { useRouter } from "vue-router"

const router = useRouter()

const picPath = ref("")
// 获取验证码
const loginVerify = () => {
  getCaptcha({}).then((ele) => {
    if (ele.code != 200) {
      ElMessage.error("无法获取验证码")
      return
    }
    picPath.value = ele.data.captchaValue
    loginFormData.captchaId = ele.data.captchaId
  });
};
loginVerify()

// 验证函数
const checkEmail = (rule: any, value: string, callback: any) => {
  const regEmail = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/
  if (regEmail.test(value)) return callback()
  else return callback(new Error("请输入正确的邮箱"))
};
const checkPassword = (rule: any, value: string | any[], callback: any) => {
  if (value.length < 1) {
    return callback(new Error("请输入密码"))
  } else {
    callback()
  }
};
const rules = reactive({
  email: [{ validator: checkEmail, trigger: "blur" }],
  password: [{ validator: checkPassword, trigger: "blur" }],
  captcha: [{ required: true, message: "请输入验证码", trigger: "blur" }],
});
// 登录
const loginFormData = reactive({
  email: "",
  password: "",
  captchaValue: "",
  captchaId: "",
});
const submitForm = () => {
  if (
    loginFormData.email === "" ||
    loginFormData.password === "" ||
    loginFormData.captchaValue === ""
  ) {
    ElMessage.warning("请填写完整信息")
    return;
  }
  loginSubmit(loginFormData).then((ele) => {
    if (ele.code != 200) {
      loginVerify()
      return
    }
    sessionStorage.setItem("Authorization", ele.data.token)
    sessionStorage.setItem("UserEmail", loginFormData.email)
    ElMessage.success("登录成功")
    router.push({ name: "Home" })
  });
};
// 注册
const dialogFormVisible = ref(false)
const form = reactive({
  name: "",
  password: "",
  email: "",
});
const registerSubmit = () => {
  if (form.name === "" || form.password === "" || form.email === "") {
    ElMessage.warning("请填写完整信息")
    return
  }
  registerUser(form).then((ele) => {
    if (ele.code !== 200) {
      return
    }
    ElMessage.success("注册成功")
    loginFormData.email = form.email
    loginFormData.password = form.password
    dialogFormVisible.value = false
  })
};
const cancelSubmit = () => {
  form.name = ""
  form.password = ""
  form.email = ""
  dialogFormVisible.value = false
};
</script>

<style lang="scss" scoped>
@import "../styles/login.scss";
</style>
