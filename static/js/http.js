"use strict";

var baseUrl = "http://www.xaxin.cn";
var Axios = axios.create({
  baseURL: baseUrl,
  timeout: 1000 * 60,
  withCredentials: false //带上 cookie

}); //添加一个请求拦截器

Axios.interceptors.request.use(function (config) {
  //POST传参序列化
  if (config.method === 'post') {
    config.hideError = config.data ? config.data.hideError : false; // 数据序列化成表单

    if (config.data && config.data.data) {
      var formData = new FormData();
      Object.keys(config.data.data).forEach(function (key) {
        return formData.append(key, config.data.data[key]);
      });
      config.data = formData;
    }
  }

  if (config.method === 'put' || config.method == 'delete') {
    var url = '';

    if (config.data && config.data.data) {
      for (var key in config.data.data) {
        url += key + '=' + config.data.data[key] + '&';
      }

      url = url.slice(0, url.length - 1);
    }

    config.data = url;
  }

  if (config.method === 'get' && config.url.includes('/api/') && config.data) {
    var _url = '?';

    for (var _key in config.data) {
      _url += _key + '=' + config.data[_key] + '&';
    }

    _url = _url.slice(0, _url.length - 1);
    config.url += _url;
  }

  return config;
}, function (error) {
  alert('参数错误！');
}); //添加一个返回拦截器

Axios.interceptors.response.use(function (res) {
  var data = res.data;
  return data;
}, function (error) {
  return error;
});

Axios.Get = function (url, data) {
  return Axios.get(url, {
    data: data
  });
};

Axios.Post = function (url, data) {
  return Axios.post(url, data);
};