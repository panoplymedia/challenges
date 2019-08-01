<template>
  <div id="app">
    <ul>
        <router-link to="/register">Register</router-link>
        <router-link to="/login">Login</router-link>
        <router-link to="/sales_data/upload">Upload</router-link>
        <router-link to="/sales_data/sales_numbers">Get Sales</router-link>
    </ul>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
