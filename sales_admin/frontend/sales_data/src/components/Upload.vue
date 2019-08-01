<template>
  <html>
    <body>
      <form class="label "@submit.prevent="Submit" enctype="multipart/form-data">
          <p> Select csv sales data file to upload: </p>
        <input type="file" ref="file" @change="onSelect">
      </form>
      <div class="fields">
        <button class="btn btn-primary" @click="submit()">Upload Sales Data</button>
      </div>
    </body>
  </html>
</template>


<script>
import auth from '../auth'
export default {
  name: "FileUpload",
  data() {
    return {
      file: '',
      message: '',
      error: ''
    }
  },
  methods: {
    onSelect() {
      const file = this.$refs.file.files[0]
      this.file = file
    },
    async submit() {
      const formData = new FormData()
      formData.append('file', this.file)

      try{
        auth.upload(this, formData, 'sales_numbers');
      }
      catch(err) {
        consol.log(err);
      }
    }
  }

}

</script>
