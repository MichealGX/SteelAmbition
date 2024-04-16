"use strict";
const common_vendor = require("../../common/vendor.js");
const _sfc_main = {
  data() {
    return {
      activeTab: "login",
      loginData: {
        username: "",
        password: ""
      },
      registerData: {
        username: "",
        password: "",
        emit: ""
      }
    };
  },
  methods: {
    showLogin() {
      this.activeTab = "login";
    },
    showRegister() {
      this.activeTab = "register";
    },
    // 点击登录按钮
    logins() {
      let thit = this;
      if (!thit.loginData.username) {
        return;
      }
      if (!thit.loginData.password) {
        return;
      }
      common_vendor.index.request(`http://127.0.0.1:4523/m1/4225974-0-default/platform/users/login`, {
        method: "POST",
        data: {
          username: thit.loginData.username || "",
          psw: thit.loginData.password || ""
        }
      }).then((res) => {
        console.log(res, "res");
        if (res.code === 200) {
          common_vendor.index.switchTab({
            url: "/pages/tabbar/index/index"
          });
        }
      }).catch((err) => {
        console.error(err);
      });
    },
    // 点击注册按钮
    registers() {
      let thit_ = this;
      if (!thit_.registerData.username) {
        return;
      }
      if (!thit_.registerData.password) {
        return;
      }
      common_vendor.index.request(`http://127.0.0.1:4523/m1/4225974-0-default/platform/users/register`, {
        method: "POST",
        data: {
          username: thit_.registerData.username || "",
          password: thit_.registerData.password || "",
          email: thit_.registerData.emit ? thit_.registerData.emit : ""
        }
      }).then((res) => {
        console.log(res, "res");
        if (res.code === 200) {
          common_vendor.index.navigateBack({ delta: 1 });
        }
      }).catch((err) => {
        console.error(err);
      });
    }
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return common_vendor.e({
    a: common_vendor.o((...args) => $options.showLogin && $options.showLogin(...args)),
    b: $data.activeTab === "login" ? 1 : "",
    c: common_vendor.o((...args) => $options.showRegister && $options.showRegister(...args)),
    d: $data.activeTab === "register" ? 1 : "",
    e: $data.activeTab === "login"
  }, $data.activeTab === "login" ? {
    f: $data.loginData.username,
    g: common_vendor.o(($event) => $data.loginData.username = $event.detail.value),
    h: $data.loginData.password,
    i: common_vendor.o(($event) => $data.loginData.password = $event.detail.value),
    j: common_vendor.o((...args) => $options.logins && $options.logins(...args))
  } : {}, {
    k: $data.activeTab === "register"
  }, $data.activeTab === "register" ? {
    l: $data.registerData.username,
    m: common_vendor.o(($event) => $data.registerData.username = $event.detail.value),
    n: $data.registerData.password,
    o: common_vendor.o(($event) => $data.registerData.password = $event.detail.value),
    p: $data.registerData.emit,
    q: common_vendor.o(($event) => $data.registerData.emit = $event.detail.value),
    r: common_vendor.o((...args) => $options.registers && $options.registers(...args))
  } : {});
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__file", "C:/Users/lenovo/Desktop/活 4-7 要 - uniapp 小程序/pages/login/login.vue"]]);
wx.createPage(MiniProgramPage);
