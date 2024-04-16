"use strict";
const common_vendor = require("../../../common/vendor.js");
const _sfc_main = {
  methods: {
    goToLogin() {
      common_vendor.index.navigateTo({
        url: "/pages/login/login"
      });
    },
    goToMyCar() {
      common_vendor.index.navigateTo({
        url: "/pages/my_car/my_car"
      });
    },
    goToConfigCar() {
      common_vendor.index.navigateTo({
        url: "/pages/config_car/config_car"
      });
    },
    goToGameRecords() {
      common_vendor.index.navigateTo({
        url: "/pages/game_records/game_records"
      });
    }
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.o((...args) => $options.goToLogin && $options.goToLogin(...args)),
    b: common_vendor.o((...args) => $options.goToMyCar && $options.goToMyCar(...args)),
    c: common_vendor.o((...args) => $options.goToConfigCar && $options.goToConfigCar(...args)),
    d: common_vendor.o((...args) => $options.goToGameRecords && $options.goToGameRecords(...args))
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__file", "C:/Users/lenovo/Desktop/活 4-7 要 - uniapp 小程序/pages/tabbar/mine/mine.vue"]]);
wx.createPage(MiniProgramPage);
