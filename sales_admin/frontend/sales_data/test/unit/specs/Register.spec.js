import Vue from 'vue'
import Register from '@/components/Register'

describe('Register.vue', () => {
  it('should render correct contents', () => {
    const Constructor = Vue.extend(Register)
    const vm = new Constructor().$mount()
    expect(vm.$el.querySelector('.col-sm-4 p').textContent)
      .toEqual('Register an account to upload sales data.')
  })
})
