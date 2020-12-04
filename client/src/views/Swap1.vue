<template>
  <div
    class="Swap1 container-fluid"
    style="
      background-attachment: default;
      background-size: cover;
      height: 100vh;
    "
    :style="{ backgroundImage: `url('${bgimage}')` }"
  >
    <b-jumbotron
      text-variant="light"
      class="mb-1 pb-1"
      :fluid="true"
      header-level="4"
      header="Upload Images to Perform Face Swap!"
      bg-variant="transparent"
    >
      <p class="my-5 text-left font-weight-light pl-4" style="font-size: 150%">
        Max image size: 10MB
      </p>
    </b-jumbotron>
    <v-uploader
      class="w-50 align-middle d-inline-block"
      @done="vUploadDone"
      file-size-limit="10MB"
      :multiple="true"
      button-text="Upload for Face Swap"
      language="en"
      :before-upload="beforeUpload"
      :item-limit=2
    />
  </div>
</template>

<script>
import Vue from "vue";
import vUploader from "../components/uploader_index";
import Axios from "axios";
import imageDisplay from "../components/imageDisplay";
export default {
  components: {},
  data() {
    return {
      receiveURL: null,
      imageList: [],
      bgimage: require("../assets/bg1.jpg"),
      swapResponse: null,
    };
  },
  methods: {
    openResultModal() {
      this.$dlg.modal(imageDisplay, {
        title : "Have Fun 😉",
        width: 800,
        height: 700,
        params: {
          url: this.receiveURL
        }
      }
        
      )
    },
    loadingImg() {
      const key = this.$dlg.mask("Processing...  ", () => (this.openResultModal()));
      Axios.post("http://127.0.0.1:5555/uploadMinePost/swap", {
        file1: this.imageList[0],
        file2: this.imageList[1],
      }).then((response) => (this.swapResponse = response,
        this.receiveURL = response.data.url,
        console.log(this.receiveURL),
        this.$dlg.close(key)
      ));
      
    },
    vUploadDone(f) {
      if (f.length == 2) {
        this.$dlg.alert(
          "You Are Ready to Swap Faces. Do you CONFIRM to Swap? ",
          () => {this.loadingImg()},
          {
            messageType: "confirm",
            language: "en",
            cancelCallback: function () {
              window.location.reload();
            },
          }
        );
      }
    },
    beforeUpload: function (id, name) {
      var fe = /(?:\.([^.]+))?$/;
      var fileExt = fe.exec(name)[1];
      if (fileExt === undefined) {
        this.$dlg.alert("Cannot detect uploaded image's file extension", null, {
          messageType: "error",
          language: "en",
        });
      }
      if (id === 0 || id === 1) {
        this.imageList.push(name);
        // console.log(this.imageList)
      }
      // if (id === 2) {
      //   this.$dlg.alert(
      //     "Exceeded maximum images: 2 \n Uploading interrupted.",
      //     null,
      //     {
      //       messageType: "error",
      //       language: "en",
      //     }
      //   );
      //   return false;
      // }
      // console.log("id is:"+id+" file ext is:"+fileExt);
      return true;
    },
  },
  computed: {},
  created() {
    const uploaderConfig = {
      // For Vagrant
      uploadFileUrl: "http://127.0.0.1:5555/uploadMinePost/",
      deleteFileUrl: "",
      showMessage: (vue, message) => {
        // using v-dialogs to show message
        vue.$dlg.alert(message, null, { messageType: "error" });
      },
    };
    Vue.use(vUploader, uploaderConfig);
  },
};
</script>

<style>
</style>