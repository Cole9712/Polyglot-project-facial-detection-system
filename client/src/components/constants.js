/**
 *
 * @param fileSize
 * @param thousand
 */
let fileSize2Bytes = (fileSize, thousand) => {
  if (!fileSize) return null;
  let baseUnit = thousand ? 1000 : 1024;
  let kb = baseUnit,
    mb = kb * baseUnit,
    gb = mb * baseUnit;
  let tmpCode = fileSize.substring(fileSize.length - 2).toUpperCase();
  let sizeNumber = fileSize.substring(0, fileSize.length - 2);
  let num = Number.parseInt(sizeNumber);
  let result = 0;
  switch (tmpCode) {
    case "KB":
      result = num * kb;
      break;
    case "MB":
      result = num * mb;
      break;
    case "GB":
      result = num * gb;
      break;
  }
  return result;
};

const i18n = {
  en: {
    ui: {
      thumbnail: "thumbnail",
      choseFileButton: "select file",
      dropHere: "drop files here",
      done: "done",
      fileTypes: "file extensions",
      fileSizeLimit: "file size limit",
    },
  },
};

const buildOptions = function () {
  let p = {
    multiple: this.multiple,
    request: {
      endpoint: this.uploadFileUrl,
      inputName: this.uploadFileObjName,
      //server side validate file info
      params: {
        fileSizeLimit: this.fileSizeLimit,
        fileTypeExts: this.fileTypeExts,
      },
    },
    deleteFile: {
      enabled: false,
      method: "POST", //,
      //'endpoint' : $webroot + 'upload/deleteUploadFile'
    },
    debug: true,
    validation: {
      allowedExtensions: this.fileTypeExts.split(","),
      sizeLimit: fileSize2Bytes(this.fileSizeLimit, true),
      sizeLimitStr: this.fileSizeLimit,
      image: {
        maxHeight: this.imageMaxHeight,
        maxWidth: this.imageMaxWidth,
        minHeight: this.imageMinHeight,
        minWidth: this.imageMinWidth,
      },
    },
    callbacks: {
      //the callback when file upload finish
      onComplete: function (id, name, json, xhr) {},
      //the callback before delete file, return false can stop it.
      onSubmitDelete: function (id) {},
    },
    cors: {
      expected: true,
      sendCredentials: false,
    },
  };
  if (this.language && this.language === "cn") {
    p.messages = i18n.cn.messages;
    p.text = i18n.cn.text;
  }
  if (this.callback && typeof this.callback === "function") {
    p.callbacks.onComplete = (id, name, json, xhr) => {
      if (json) this.callback(json);
    };
  }
  return p;
};

const getI18n = (language) =>
  !language || language !== "en" ? i18n.cn.ui : i18n.en.ui;

export { buildOptions };
export { getI18n };
