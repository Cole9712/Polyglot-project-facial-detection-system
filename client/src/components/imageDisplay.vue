<template>
  <div>
    <img
      class="mx-auto"
      :src="url"
      alt="Image here."
      @load="scaleImg"
      style="height: 600px; max-width: 800px"
    />
  </div>
</template>
<script>
export default {
  props: {
    url: String,
  },
  data() {
    return {
      imgHeight: null,
      imgWidth: null,
    };
  },
  methods: {
    scaleImg(event) {
      // let img = event.target;
      // this.imgHeight =
    },
  },
};
</script>