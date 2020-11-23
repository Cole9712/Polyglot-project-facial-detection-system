<template>
  <div class="home container">
    <b-jumbotron :style="{ backgroundImage: `url(https://cdn.pixabay.com/photo/2020/11/16/00/18/columns-5747584_960_720.jpg)`}" text-variant="light" lead="And perform Facial Detection!">
      <template #header>Get Image from Pixabay.com</template>
      <hr class="my-4" style="height: 2px;color: #333;" variant="danger">
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
    max-width: 70%;
    text-align: center;
    margin: 0 auto;
  }
  b-jumbotron {
    background-size: cover;
  }
  
</style>
