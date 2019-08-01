<template>
    <div class="col-sm-4 col-sm-offset-4">
      <h2>Login</h2>
      <p>Login to your account to upload sales data, default login below.</p>
      <div class="alert alert-danger" v-if="error">
        <p>{{ error }}</p>
      </div>
      <div class="default">
        <p>Username = admin</p>
        <p>Password = admin</p>
      </div>
      <div class="form-group">
        <label for="username">Username</label>
        <input
          type="text"
          name="username"
          class="form-control"
          placeholder="Enter your username"
          v-model="username"
        >
      </div>
      <div class="form-group">
        <label htmlFor="password">Password</label>
        <input
          type="password"
          name="password"
          class="form-control"
          placeholder="Enter your password"
          v-model="password"
        >
      </div>
      <button class="btn btn-primary" @click="submit()">Login</button>
    </div>
  </template>

  <script>
  import auth from '../auth'
  export default {
    data() {
      return {
        error: ''
      }
    },
    methods: {
      submit() {
        this.submitted = true;
        const { username, password } = this;

        // stop here if form is invalid
        if (!(username && password)) {
            return;
        }
        this.loading = true;

        var credentials = {
          username: username,
          password: password
        }
        auth.login(this, credentials, 'sales_data/upload')
      }
    }

  }
  </script>
