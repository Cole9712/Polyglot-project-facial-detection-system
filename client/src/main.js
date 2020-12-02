import Vue from 'vue'
import App from './App.vue'
import router from './router'
import VueResource from 'vue-resource'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
// vUploader is Adapted from https://github.com/TerryZ/v-uploader
import vUploader from './components/uploader_index'

const uploaderConfig = {
  // For Vagrant
  // uploadFileUrl: 'http://localhost:5555/uploadMine',
  // For Dev purposes
  uploadFileUrl: 'http://localhost:8082/uploadMine',
  // uploadFileUrl: 'https://api.imgur.com/3/image',
  deleteFileUrl: '',
  showMessage: (vue, message) => {
    // using v-dialogs to show message
    vue.$dlg.alert(message, null, { messageType: 'error' })
  }
}

Vue.use(VueResource)
// Vue.http.headers.common['Content-Type'] = 'application/json'
// Vue.http.headers.common['Access-Control-Allow-Origin'] = '*'
// Vue.http.headers.common['Accept'] = 'application/json, text/plain, */*'
// Vue.http.headers.common['Access-Control-Request-Method'] = '*'
// Vue.http.headers.common['Access-Control-Allow-Headers'] = 'Origin, Accept, Content-Type, Authorization, Access-Control-Allow-Origin'
// Vue.http.options.emulateJSON = true
Vue.config.productionTip = false
// Install BootstrapVue
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
Vue.use(vUploader, uploaderConfig);

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
