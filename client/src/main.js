import Vue from 'vue'
import App from './App.vue'
import router from './router'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
// Adapted from https://github.com/TerryZ/v-uploader
import vUploader from 'v-uploader'

const uploaderConfig = {
  uploadFileUrl: 'http://localhost:8082/uploadMine',
  // For testing purposes
  // uploadFileUrl: 'https://api.imgur.com/3/image',
  deleteFileUrl: '',
  showMessage: (vue, message) => {
    // using v-dialogs to show message
    vue.$dlg.alert(message, null, { messageType: 'error' })
  }
}

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
