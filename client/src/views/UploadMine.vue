<template>
  <div class="UploadMine container-fluid" 
    style="background-attachment:default;
    background-size:cover;
    height: 100vh;"
    :style="{ backgroundImage: `url('${bgimage}')` }">

    <b-jumbotron text-variant="light" class="mb-1 pb-1" :fluid='true' header-level="4" header="Upload Image to Perform Facial Detection" bg-variant="transparent">
      <p class="my-5 text-left font-weight-light pl-4" style="font-size:150%;">Max image size: 10MB</p>
    </b-jumbotron>
    <v-uploader
    @done="vUploadDone"
    file-size-limit='10MB'
    :preview-width="700"
    :preview-height="525"
    button-text='Upload for Facial Detection'
    language='en'
    />
    
  </div>
</template>

<script>
  import Vue from 'vue'
  // vUploader is Adapted from https://github.com/TerryZ/v-uploader
  import vUploader from '../components/uploader_index'
  
  export default {
    components: {

    },
    data() {
      return {
        bgimage: require('../assets/bg1.jpg')
      };
    },
    methods: {
      vUploadDone(f) {
        console.log(f)
      }
    },
    created() {
      const uploaderConfig = {
        // For Vagrant
        uploadFileUrl: 'http://localhost:5555/uploadMinePost',
        // For Dev purposes
        // uploadFileUrl: 'http://localhost:8082/uploadMinePost',
        deleteFileUrl: '',
        showMessage: (vue, message) => {
          // using v-dialogs to show message
          vue.$dlg.alert(message, null, { messageType: 'error' })
        }
      };
      Vue.use(vUploader, uploaderConfig);
    }
  };
</script>

