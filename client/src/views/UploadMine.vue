<template>
  <div class="UploadMine container">
    <b-jumbotron fluid="true" header-level="4" header="Upload your own image file to detect!" bg-variant="white">
      <p class="my-5 text-left font-weight-light" style="font-size:150%;">Max image size: 2MB</p>
    </b-jumbotron>
    
    <vue-base64-file-upload 
        class="v1"
        accept="image/png,image/jpeg"
        image-class="v1-image"
        input-class="v1-input"
        :max-size="customImageMaxSize"
        @size-exceeded="onSizeExceeded"
        @load="onLoad" />
  </div>
</template>

<script>
// Adapted from https://github.com/dhhb/vue-base64-file-upload
import VueBase64FileUpload from "vue-base64-file-upload";
export default {
  components: {
    VueBase64FileUpload,
  },
  data() {
    return {
      // set max image size to 2 Mega Bytes
      customImageMaxSize: 2,
    };
  },
  methods: {
    onLoad(dataurl) {
      console.log(dataurl)
    },
    onSizeExceeded(size) {
      alert(`Image ${size}Mb size exceeds limits of ${this.customImageMaxSize}Mb!`);
    },
  },
};
</script>
