<template>
  <div class="register">
    <b-row class="mt-5">
      <b-col
        md="8"
        offset-md="2"
        lg="6"
        offset-lg="3"
      >
        <b-card title="注册">
          <b-form>
            <b-form-group label="姓名">
              <b-form-input
                id="input-1"
                v-model="$v.user.name.$model"
                type="text"
                placeholder="输入您的名称（选填）"
              ></b-form-input>
            </b-form-group>

            <b-form-group label="手机号">
              <b-form-input
                id="input-2"
                v-model="$v.user.telephone.$model"
                type="number"
                placeholder="输入您的电话号码"
                :state="validateState('telephone')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('telephone')">
                手机号不符合要求
              </b-form-invalid-feedback>
              <b-form-valid-feedback :state="validation">
                Good Job
              </b-form-valid-feedback>
            </b-form-group>

            <b-form-group label="密码">
              <b-form-input
                id="input-3"
                v-model="$v.user.password.$model"
                type="password"
                placeholder="输入您的密码"
                :state="validateState('password')"
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('password')">
                密码不可小于六位
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group>
              <b-button
                @click="register"
                variant="outline-primary"
                block
              >注册</b-button>
            </b-form-group>

          </b-form>
        </b-card>
      </b-col>
    </b-row>
  </div>
</template>

<script>
import { required, minLength } from 'vuelidate/lib/validators';

import customValidator from '@/helper/validator';
import { mapActions } from 'vuex';

export default {
  data() {
    return {
      user: {
        name: '',
        telephone: '',
        password: '',
      },
      validation: null,
    };
  },
  // validation验证
  validations: {
    user: {
      name: {

      },
      telephone: {
        required,
        telephhone: customValidator.telephoneValidator,
      },
      password: {
        required,
        minLength: minLength(6),

      },
    },
  },

  methods: {
    ...mapActions('userModule', { userRegister: 'register' }),
    validateState(name) {
      // es6的解构赋值
      const { $dirty, $error } = this.$v.user[name];
      return $dirty ? !$error : null;
    },
    register() {
      // 验证数据
      this.$v.user.$touch();
      if (this.$v.user.$anyError) {
        return;
      }

      // 请求Api
      this.userRegister(this.user).then(() => {
        // 跳转主页
        this.$router.replace({ name: 'Home' });
      }).catch((err) => {
        this.$bvToast.toast(err.response.data.msg, {
          title: '数据验证错误',
          variant: 'danger',
          solid: true,
        });
      });

      console.log('register');
    },

  },
};
</script>

<style lang="scss" scoped>
</style>
