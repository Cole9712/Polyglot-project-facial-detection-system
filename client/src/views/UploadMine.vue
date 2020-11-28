<template>
  <div class="UploadMine container-fluid" 
    style="background-attachment:default;
    background-size:cover;
    height: 100vh;"
    :style="{ backgroundImage: `url('${bgimage}')` }">

    <b-jumbotron text-variant="light" class="mb-1 pb-1" :fluid='true' header-level="4" header="Upload your own image file!" bg-variant="transparent">
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
};
</script>

