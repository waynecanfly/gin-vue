<template>
  <div class="register">
    <b-row>
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="注册" class="mt-5">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                v-model="$v.user.name.$model"
                type="text"
                required
                placeholder="输入您的名称（选填）"
              ></b-form-input>
            </b-form-group>
            <b-form-group label="手机号">
              <b-form-input
                v-model="$v.user.tel.$model"
                type="number"
                required
                placeholder="输入您的手机号"
                :state="validateState('tel')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('tel')">
                手机号不符合要求
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group label="密码">
              <b-form-input
                v-model="$v.user.pwd.$model"
                type="password"
                required
                placeholder="输入密码"
                :state="validateState('pwd')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('pwd')">
                密码必须大于等于6位
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-button @click="register" variant="outline-primary" block>注册</b-button>
            </b-form-group>
          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>
<script>
import { required, maxLength } from 'vuelidate/lib/validators';
import customValidator from '@/helper/validator';

export default {
  data() {
    return {
      user: {
        name: '',
        tel: '',
        pwd: '',
      },
      validation: null,
    };
  },
  validations: {
    user: {
      name: {

      },
      tel: {
        required,
        tel: customValidator.telValidator,
      },
      pwd: {
        required,
        maxLength: maxLength(6),
      },
    },
  },
  methods: {
    validateState(name) {
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }
      // 请求
      const api = 'http://localhost:1016/api/auth/register';
      this.axios.post(api, { ...this.user }).then((res) => {
        console.log(res.data);
      }).catch((err) => {
        console.log('err:', err.response.data.msg);
      });
      console.log('register');
    },
  },
};
</script>
<style scoped></style>
