<template>
   <div>
     <h1>Total sales revenue: ${{ sales_revenue }}</h1>
   </div>
</template>



<script>
import auth from '../auth'
import axios from 'axios'
export default {
  data() {
    return {
      sales_revenue: '',
      error: ''
    }
  },
  created: function(){
    return axios.get('http://localhost:5000/sales_data/total_revenue',
     {headers: auth.getAuthHeader()})
    .then(result => {
      this.sales_revenue = result.data.total_revenue
    })
    .catch(error => {
      console.log(error)
    })
  }
  }
</script>
