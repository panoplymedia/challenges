<template>
    <div class="col-sm-4 col-sm-offset-4">
      <h2>Register Account</h2>
      <p>Register an account to upload sales data.</p>
      <div class="alert alert-danger" v-if="error">
        <p>{{ error }}</p>
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
      <button class="btn btn-primary" @click="submit()">Register</button>
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

        // Check if form is valid
        if (!(username && password)) {
            return;
        }
        this.loading = true;

        var credentials = {
          username: username,
          password: password
        }
        auth.register(this, credentials, 'login')
      }
    }

  }
  </script>
