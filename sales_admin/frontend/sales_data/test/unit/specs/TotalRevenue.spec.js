import Vue from 'vue'
import TotalRevenue from '@/components/TotalRevenue'

describe('Upload.vue', () => {
  it('should render correct contents', () => {
    const Constructor = Vue.extend(TotalRevenue)
    const vm = new Constructor().$mount()
    expect(vm.$el.querySelector('h1').textContent)
      .toEqual('Total sales revenue: $')
  })
})
