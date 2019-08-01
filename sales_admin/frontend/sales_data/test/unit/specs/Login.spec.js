import Vue from 'vue'
import Login from '@/components/Login'

describe('Login.vue', () => {
  it('should render correct contents', () => {
    const Constructor = Vue.extend(Login)
    const vm = new Constructor().$mount()
    expect(vm.$el.querySelector('.col-sm-4 p').textContent)
      .toEqual('Login to your account to upload sales data, default login below.')
    expect(vm.$el.querySelector('.default').textContent)
      .toEqual("Username = admin Password = admin")
  })
})
