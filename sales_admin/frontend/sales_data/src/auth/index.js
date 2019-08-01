import router from '../router'
import axios from 'axios'

// URL and endpoint constants
const API_URL = 'http://localhost:5000'
const LOGIN_URL = API_URL + '/auth/login'
const REGISTER_URL = API_URL + '/auth/register'
const UPLOAD_URL = API_URL + '/sales_data/upload'
const TOTAL_REVENUE = API_URL + '/sales_data/total_revenue'

export default {

  // User object will let us check authentication status
  user: {
    authenticated: false
  },

  // Send a request to the login URL and save the returned JWT
  login(context, creds, redirect) {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(creds)
    };

    axios.post(LOGIN_URL, requestOptions)
    .then(result => {
      this.user.authenticated = true
      localStorage.setItem('access_token', result.data.access_token)

      // Redirect to a specified route
      if(redirect) {
        router.push(redirect)
      }
    })
    .catch(error => {
      console.log(error)
    })
  },

  register(context, creds, redirect) {
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(creds)
    };

    axios.post(REGISTER_URL, requestOptions)
    .then(result => {
      // Redirect to a specified route
      if(redirect) {
        router.push(redirect)
      }
    })
    .catch(error => {
      console.log(error)
    })
  },

  // Not used but to log out, we just need to remove the token
  logout() {
    localStorage.removeItem('access_token')
    this.user.authenticated = false
  },

  // The object to be passed as a header for authenticated requests
  getAuthHeader() {
    return {
      'Authorization': 'Bearer ' + localStorage.getItem('access_token'),
    }
  },

  // Uploads csv sales data file
  upload(context, salesData, redirect) {
      var h = this.getAuthHeader()
      return axios.post(UPLOAD_URL, salesData, {headers: h})
      .then(result => {
        console.log("Sales data uploaded successfully")
        if(redirect) {
          router.push(redirect)
        }
      })
      .catch(error => {
        console.log(error)
      })
  },
}
