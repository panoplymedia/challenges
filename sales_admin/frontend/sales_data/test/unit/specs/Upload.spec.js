import Vue from 'vue'
import Upload from '@/components/Upload'

describe('Upload.vue', () => {
  it('should render correct contents', () => {
    const Constructor = Vue.extend(Upload)
    const vm = new Constructor().$mount()
    expect(vm.$el.querySelector('.label p').textContent)
      .toEqual(' Select csv sales data file to upload: ')
  })
})
