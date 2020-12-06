<template>
  <div class="center">
    <div class="results md-layout">
      <md-card md-with-hover class="md-layout-item md-size-15 md-alignment-center-center" v-for="r in searchResults" :key="r.id">
        <md-card-media v-if="type == 'image'">
          <img :src="r.webformatURL" class="image" />
        </md-card-media>
      </md-card>
    </div>
  </div>
</template>
<script>
export default {
  name: "results",
  props: {
    type: String
  },
  computed: {
    searchResults() {
      return this.$store.state.searchResults;
    }
  },
  data() {
    return {};
  }
};
</script>

<style scoped lang="scss">
.md-card {
  width: 30vw;
  margin: 4px;
  display: inline-block;
  vertical-align: top;
}
</style>