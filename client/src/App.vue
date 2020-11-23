<template>
  <div id="app">
    <div>
      <b-navbar type="dark" variant="dark">
        <b-navbar-brand tag="h1" href="#">Facial Detection Project</b-navbar-brand>
        <b-collapse is-nav>
          <b-navbar-nav>
            <b-nav-item @click='$router.push({name:"Home"})'>Home</b-nav-item>
            <b-nav-item @click='$router.push({name:"UploadMine"})'>Upload My Image</b-nav-item>
          </b-navbar-nav>
        </b-collapse>
      </b-navbar>
    </div>
    <!-- <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
    </div> -->
    <router-view />
  </div>
</template>

<script>
export default {
  name: "facial-app",
  components: {},
  data() {
    return {};
  },
};
</script>

<style scoped>
#app {
  font-family: Lato, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

</style>
