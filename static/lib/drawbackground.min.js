/**
 * drawbackground 1.0.2
 * created at Sun Feb 04 2018 12:23:03 GMT+0800 (CST)
 */

(function (global, factory) {
	typeof exports === 'object' && typeof module !== 'undefined' ? module.exports = factory() :
	typeof define === 'function' && define.amd ? define(factory) :
	(global.drawbackground = factory());
}(this, (function () { 'use strict';

function __$styleInject(css, returnValue) {
  if (typeof document === 'undefined') {
    return returnValue;
  }
  css = css || '';
  var head = document.head || document.getElementsByTagName('head')[0];
  var style = document.createElement('style');
  style.type = 'text/css';
  head.appendChild(style);
  
  if (style.styleSheet){
    style.styleSheet.cssText = css;
  } else {
    style.appendChild(document.createTextNode(css));
  }
  return returnValue;
}

function drawBackground(option) {
    if ( option === void 0 ) option = {};

    var config = {
        zindex: option.zindex || 11,
        opacity: option.opacity || 0.5,
        color: option.color || "0.0.0",
        count: option.count || 99
    };
    //查找元素
    var $ = function(name) {
        return document.querySelectorAll(name)
    };

    //设置canvas宽高
    var setCanvasSize = function() {
        canvas_width = the_canvas.width = window.innerWidth || document.documentElement.clientWidth || document.body.clientWidth;
        canvas_height = the_canvas.height = window.innerHeight || document.documentElement.clientHeight || document.body.clientHeight;
    };

    //绘制canvas的过程
    var drawCanvas = function() {
        //每次开始绘制前  清理屏幕
        context.clearRect(0, 0, canvas_width, canvas_height);
        /*
        	思路：
        		1.双层循环点的集合
        */
        var e, x_dist, y_dist, dist, line_width, i;
        random_points.forEach(function (item, index) {
            item.x += item.xa;
            item.y += item.ya;
            item.xa *= item.x > canvas_width || item.x < 0 ? -1 : 1;
            item.ya *= item.y > canvas_height || item.y < 0 ? -1 : 1;
            context.fillRect(item.x - 0.5, item.y - 0.5, 1, 1);
            for (i = index + 1; i < allPoint.length; i++) {
                e = allPoint[i];
                if (!!e.x && !!e.y) {
                    x_dist = item.x - e.x;
                    y_dist = item.y - e.y;
                    dist = x_dist * x_dist + y_dist * y_dist;
                    dist < e.max && (e == current_point && dist >= e.max / 2 && (item.x -= 0.03 * x_dist, item.y -= 0.03 * y_dist), line_width = (e.max - dist) / e.max, context.beginPath(), context.lineWidth = line_width, context.strokeStyle = "rgba(" + config.color + "," + (line_width + 0.2) + ")", context.moveTo(item.x, item.y), context.lineTo(e.x, e.y), context.stroke());
                }
            }
        });

        frame_func(drawCanvas);

    };

    var the_canvas = document.createElement('canvas'),
        canvas_id = 'canvas_id_' + config.l,
        context = the_canvas.getContext("2d"),
        canvas_width, canvas_height,
        frame_func = window.requestAnimationFrame || window.webkitRequestAnimationFrame || window.mozRequestAnimationFrame || window.oRequestAnimationFrame || window.msRequestAnimationFrame || function(func) {
            window.setTimeout(func, 1000 / 60);
        },
        current_point = {
            x: null, //当前鼠标x
            y: null, //当前鼠标y
            max: 20000 // 圈半径的平方
        },
        allPoint,
        random_points = [];
    the_canvas.id = canvas_id;
    the_canvas.style.cssText = "position:fixed;top:0;left:0;z-index:" + config.zindex + ";opacity:" + config.opacity;
    $("body")[0].appendChild(the_canvas);

    //初始化 画布大小
    setCanvasSize();
    window.onresize = setCanvasSize;

    window.onmousemove = function(e) {
        e = e || window.event;
        current_point.x = e.clientX;
        current_point.y = e.clientY;
    };
    window.onmouseout = function() {
        current_point.x = null;
        current_point.y = null;
    };

    //生成随机点
    for (var i = 0; i < config.count; i++) {
        var x = Math.random() * canvas_width,
            y = Math.random() * canvas_height,
            xa = Math.random() * 2 - 1,
            ya = Math.random() * 2 - 1;

        random_points.push({
            x: x,
            y: y,
            xa: xa,
            ya: ya,
            max: 6000
        });
    }

    allPoint = random_points.concat([current_point]);
    //0.1s后绘制
    setTimeout(function () {
        frame_func(drawCanvas);
    }, 100);
}

return drawBackground;

})));