import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/Login'
import Register from '@/components/Register'
import Upload from '@/components/Upload'
import TotalRevenue from '@/components/TotalRevenue'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'login',
      component: Login
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: '/register',
      name: 'Register',
      component: Register
    },
    {
      path: '/sales_data/upload',
      name: 'Upload',
      component: Upload
    },
    {
      path: '/sales_data/sales_numbers',
      name: 'TotalRevenue',
      component: TotalRevenue
    }
  ]
})
