<template>
  <div class="home container-fluid" 
  style="background-color:black;
  background-size:cover;
  height: 100vh;
  width: 100%;
  background-image: url('https://pixabay.com/get/57e9d64b4a51b114f6d1847bcf21347c083ed8e55550724b75277a.jpg');">

    <b-jumbotron bg-variant="transparent" text-variant="light" lead="And perform Facial Detection!">
      <template #header>Get Image from Pixabay.com</template>
      <hr class="my-4" style="height: 2px;color: black;">
      <b-input-group style="width: 70%;margin:auto;" class="col-8 align-center" size="md">
        <b-input-group-prepend is-text>
          <b-icon icon="search"></b-icon>
        </b-input-group-prepend>
        <b-form-input id="searchBox" :type="search" v-model="searchText" placeholder="Search"></b-form-input>
        <b-input-group-postpend>
          <b-button class="mx-2" variant="success">Go!</b-button>
        </b-input-group-postpend>
      </b-input-group>
      
    </b-jumbotron>

  </div>
</template>

<script>
// @ is an alias to /src
import axios from 'axios';
export default {
  name: 'Home',
  components: {
      
  },
  props: {
    
  },
  created() {
    axios.get()
  }
}
</script>

<style scoped>
  .home {
    text-align: center;
    margin: 0 auto;
  }
  b-jumbotron {
    background-size: cover;
  }
  
</style>
