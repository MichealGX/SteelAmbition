"use strict";
const common_vendor = require("../../../common/vendor.js");
const _sfc_main = {
  data() {
    return {
      listVm: [
        { id: "1", name: "嗨(｡･∀･)ﾉﾞ嗨" },
        { id: "2", name: "嗨(｡･∀･)ﾉﾞ嗨" },
        { id: "3", name: "嗨(｡･∀･)ﾉﾞ嗨" },
        { id: "4", name: "嗨(｡･∀･)ﾉﾞ嗨" },
        { id: "5", name: "嗨(｡･∀･)ﾉﾞ嗨" },
        { id: "6", name: "嗨(｡･∀･)ﾉﾞ嗨" },
        { id: "7", name: "嗨(｡･∀･)ﾉﾞ嗨" },
        { id: "8", name: "嗨(｡･∀･)ﾉﾞ嗨" },
        { id: "9", name: "嗨(｡･∀･)ﾉﾞ嗨" },
        { id: "10", name: "嗨(｡･∀･)ﾉﾞ嗨" },
        { id: "11", name: "嗨(｡･∀･)ﾉﾞ嗨" }
      ],
      currentPage: 1,
      pageSize: 5
    };
  },
  computed: {
    totalPages() {
      return Math.ceil(this.listVm.length / this.pageSize);
    },
    displayedList() {
      const start = (this.currentPage - 1) * this.pageSize;
      const end = start + this.pageSize;
      return this.listVm.slice(start, end);
    }
  },
  methods: {
    prevPage() {
      if (this.currentPage > 1) {
        this.currentPage--;
      }
    },
    nextPage() {
      if (this.currentPage < this.totalPages) {
        this.currentPage++;
      }
    }
  }
};
function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
  return {
    a: common_vendor.f($options.displayedList, (item, index, i0) => {
      return {
        a: common_vendor.t(item.id),
        b: common_vendor.t(item.name),
        c: index
      };
    }),
    b: common_vendor.o((...args) => $options.prevPage && $options.prevPage(...args)),
    c: $data.currentPage === 1,
    d: common_vendor.t($data.currentPage),
    e: common_vendor.o((...args) => $options.nextPage && $options.nextPage(...args)),
    f: $data.currentPage === $options.totalPages
  };
}
const MiniProgramPage = /* @__PURE__ */ common_vendor._export_sfc(_sfc_main, [["render", _sfc_render], ["__file", "C:/Users/lenovo/Desktop/活 4-7 要 - uniapp 小程序/pages/tabbar/rank/rank.vue"]]);
wx.createPage(MiniProgramPage);
