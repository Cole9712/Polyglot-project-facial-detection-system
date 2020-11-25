<template>
  <div class="UploadMine container-fluid" 
    style="background-attachment:default;
    background-size:cover;
    height: 100vh;"
    :style="{ backgroundImage: `url('https://pixabay.com/get/52e0d4424350ae14f6d1867dda293e7e1337d7e6514c704f752979d39e4ac55d_1920.jpg')` }">

    <b-jumbotron text-variant="light" class="mb-1 pb-1" fluid="true" header-level="4" header="Upload your own image file!" bg-variant="transparent">
      <p class="my-5 text-left font-weight-light pl-4" style="font-size:150%;">Max image size: 10MB</p>
    </b-jumbotron>
    <v-uploader
    @done="vUploadDone"
    file-size-limit='10MB'
    :preview-width="800"
    :preview-height="600"
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
      
    };
  },
  methods: {
    vUploadDone(f) {
      console.log(f)
    }
  },
};
</script>

