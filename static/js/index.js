"use strict";

drawbackground({
  zindex: -1,
  //canvas 定位 
  opacity: 0.5,
  //canvas 透明度 
  color: "0.0.0",
  //画线颜色 
  count: 30 //点的数量 

});
Vue.prototype.http = Axios;
Vue.filter("yyyy_mm_dd_H_M_S", function (val) {
  //1989-01-11 11:05:12
  if (!val || val == 0 || val == '') return "--";
  var date = new Date(val);
  return date.getFullYear() + "-" + (date.getMonth() + 1 < 10 ? "0" + (date.getMonth() + 1) : date.getMonth() + 1) + '-' + (date.getDate() < 10 ? "0" + date.getDate() : date.getDate()) + " " + (date.getHours() < 10 ? "0" + date.getHours() : date.getHours()) + ":" + (date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes()) + ":" + (date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds());
});
var app = new Vue({
  el: "#app",
  data: {
    tabIndex: "",
    art_id: "",
    page: 1,
    limit: 10,
    count: 0,
    list: [],
    type: [],
    info: {}
  },
  created: function created() {
    this.getTypeList();
    this.getList();
  },
  mounted: function mounted() {},
  methods: {
    tabClick: function tabClick(code) {
      this.tabIndex = code;
      this.getList();
    },
    // 菜单
    clickList: function clickList(id) {
      if (id == this.art_id) return;
      this.art_id = id;
      this.getActInfo();
    },
    // 获取文章详情
    getActInfo: function getActInfo() {
      var _this = this;

      this.http.Get("/api/article/info/".concat(this.art_id)).then(function (res) {
        if (res.code == 200) {
          _this.info = res.data;
        }
      });
    },
    getList: function getList() {
      var _this2 = this;

      this.http.Get('/api/article/list', {
        page: this.page,
        limit: this.limit,
        type: this.tabIndex
      }).then(function (res) {
        if (res.code == 200) {
          _this2.list = res.data.item;
          _this2.count = res.data.count;

          if (res.data.item && res.data.item.length > 0) {
            _this2.art_id = res.data.item[0].id;

            _this2.getActInfo();
          }
        }
      })["catch"](function (error) {
        console.log(error);
      });
    },
    getTypeList: function getTypeList() {
      var _this3 = this;

      this.http.Get("/api/article/type").then(function (res) {
        if (res.code == 200) {
          _this3.type = res.data;
        }
      })["catch"](function (error) {
        console.log(error);
      });
    }
  }
});