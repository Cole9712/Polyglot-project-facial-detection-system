<template>
  <div
    class="home container-fluid"
    style="
      background-color: white;
      background-size: cover;
      height: 100vh;
      width: 100%;
    "
  >
    <b-jumbotron
      bg-variant="transparent"
      text-variant="dark"
      lead="And perform Facial Detection!"
    >
      <template #header>Get Image from Pixabay.com</template>
      <hr class="my-4" style="height: 2px; color: black" />
      <b-input-group
        style="width: 70%; margin: auto"
        class="col-8 align-center"
        size="md"
      >
        <b-input-group-prepend is-text>
          <b-icon icon="search"></b-icon>
        </b-input-group-prepend>
        <b-form-input
          id="searchBox"
          @keyup.enter="search"
          name="q"
          :type="search"
          v-model="searchData.q"
          placeholder="Search Keyword"
        ></b-form-input>
        <div>
          <b-button @click="search" class="mx-2" type="submit" variant="success"
            >Go!</b-button
          >
        </div>
      </b-input-group>
      <!-- Adapted and modified from https://medium.com/better-programming/how-to-add-infinite-scrolling-effect-to-a-vue-js-app-164280474274 -->
      <Results class="col-11 mt-2 pt-4 mx-auto" type="image" />
    </b-jumbotron>
  </div>
</template>

<script>
// @ is an alias to /src
import axios from "axios";
import Results from "../components/imgResult";
import { photosMix } from "../components/mix/photosMix";
export default {
  name: "Home",
  components: {
    Results,
  },
  data() {
    return {
      photos: [],
      searchData: {},
    };
  },
  props: {},
  mixins: [photosMix],
  beforeMount() {
    this.$store.commit("setSearchResults", []);
  },
  methods: {
    async search(evt) {
      const response = await this.searchPhoto(this.searchData);
      this.photos = response.data.hits;
      this.$store.commit("setSearchResults", response.data.hits);
    },
  },
};
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
