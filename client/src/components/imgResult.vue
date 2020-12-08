<template>
  <div class="center">
    <div class="results md-layout">
      <md-card
        md-with-hover
        class="md-layout-item md-size-15"
        v-for="r in searchResults"
        :key="r.id"
      >
        <md-card-media md-solid v-if="type == 'image'">
          <md-card-media md-ratio="4:3">
            <img :src="r.webformatURL" class="image" />
          </md-card-media>
        </md-card-media>
        <md-card-expand>
          <md-card-actions>
            <md-button @click="popupFullImg(r.largeImageURL)"><b-icon icon="zoom-in"></b-icon></md-button>
            <md-button @click="startDetection(r.largeImageURL, r.id)">Detect Faces</md-button>
          </md-card-actions>
        </md-card-expand>
      </md-card>
    </div>
  </div>
</template>
<script>
import imageDisplay from "./imageDisplay";
import Axios from "axios";
export default {
  name: "results",
  props: {
    type: String,
  },
  computed: {
    searchResults() {
      return this.$store.state.searchResults;
    },
  },
  data() {
    return {
      facialUrl: null
    };
  },
  methods: {
    popupFullImg(imgUrl) {
      this.$dlg.modal(imageDisplay, {
        title: "Lager Image",
        width: 900,
        height: 700,
        params: {
          url: imgUrl,
        },
      });
    },
    startDetection(url, inputID) {
      var extension = url.split(/[#?]/)[0].split('.').pop().trim();
      const key = this.$dlg.mask("Please Wait...  ", () => (
          this.$dlg.modal(imageDisplay, {
          title: "Detection Result",
          width: 900,
          height: 700,
          params: {
            url: this.facialUrl,
          },
        })));
      Axios.post("http://localhost:5555/homeDetection", {
        imgUrl: url,
        id: inputID,
        ext: extension,
      }).then((response) => (
        this.facialUrl = response.data.url,
        console.log(this.receiveURL),
        this.$dlg.close(key)
      ));
    }
  },
};
</script>

<style scoped lang="scss">
.md-card {
  width: 30vw;
  margin: 4px;
  display: inline-block;
  vertical-align: center;
}
</style>