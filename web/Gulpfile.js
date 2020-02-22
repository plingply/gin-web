var gulp = require('gulp'),
	less = require('gulp-less'),
    // webpack = require('gulp-webpack'),
	sourcemaps = require('gulp-sourcemaps'),
	cssmin = require('gulp-minify-css'),
	rename = require('gulp-rename'),
	imagemin = require('gulp-imagemin'),
	jshint = require('gulp-jshint'),
	uglify = require('gulp-uglify'),
	htmlmin = require('gulp-htmlmin'),
	clean = require('gulp-clean'),
	stylish = require('jshint-stylish'),
    base64 = require('gulp-base64'),
    autoprefixer = require('gulp-autoprefixer'),
	concat = require('gulp-concat'),
    htmlImport = require('gulp-html-import');



var version = '-v1.0.0',
	src = 'src/',
	out = '../static/';

var options = {   //这是配置项  路径搞不对  https://www.npmjs.com/package/gulp-base64
    baseDir: 'src/less',  //  
    extensions: ['svg', 'png', 'jpg'],
    exclude:    [/\.server\.(com|net)\/dynamic\//, '--live.jpg'],
    maxImageSize: 8*1024, // bytes 
    debug: true
}

// var webpack_config = require('./webpack.config.js');


gulp.task('less',function(){

    gulp.src(src + 'less/*.less')
    	// .pipe(sourcemaps.init())
        .pipe(less())
        // .pipe(sourcemaps.write('./map'))  //生成map文件报错
        .pipe(base64(options))//base64编码
        .pipe(autoprefixer({
            browsers: ['last 2 versions','last 3 Safari versions', 'Android >= 4.0'],
            cascade: false, //是否美化属性值 默认：true 像这样：
            //-webkit-transform: rotate(45deg);
            //        transform: rotate(45deg);
            remove:true //是否去掉不必要的前缀 默认：true 
        }))
        // .pipe(cssmin())
        // .pipe(rename({ suffix: version + '.min' }))
        .pipe(gulp.dest(out+'css'));
})



gulp.task('css',function(){    
    gulp.src(src + 'less/*.css')
        .pipe(cssmin())
        // .pipe(rename({ suffix: version + '.min' }))
        .pipe(gulp.dest(out+'css'));
})

gulp.task('img', function(){
    gulp.src(src + 'img/**/*')
        .pipe(imagemin())
        .pipe(gulp.dest(out + 'img'));
})

gulp.task('js',function(){
    gulp.src([src + 'js/*.js'])
        // .pipe(webpack(webpack_config))
    	.pipe(jshint())
    	.pipe(jshint.reporter(stylish))
    	// .pipe(rename({ suffix: version + '.min' }))
    	// .pipe(uglify())
    	.pipe(gulp.dest(out + 'js'));
})

gulp.task('lib',function(){
    gulp.src(src+'lib/' + '*.js')
        .pipe(rename({ suffix: '.min' }))
        // .pipe(uglify())
        .pipe(gulp.dest(out + 'lib'));
})

gulp.task('html',function(){
	var htmlSrc = './src/*.html',
        htmlDst = out,
        options = {
	        removeComments: false,//清除HTML注释
	        collapseWhitespace: false,//压缩HTML
	        collapseBooleanAttributes: true,//省略布尔属性的值 <input checked="true"/> ==> <input />
	        removeEmptyAttributes: true,//删除所有空格作属性值 <input id="" /> ==> <input />
	        removeScriptTypeAttributes: true,//删除<script>的type="text/javascript"
	        removeStyleLinkTypeAttributes: true,//删除<style>和<link>的type="text/css"
	        minifyJS: true,//压缩页面JS
	        minifyCSS: true//压缩页面CSS
    	};
        
    gulp.src(htmlSrc)
        .pipe(htmlImport(src + '/components/'))
        // .pipe(gulp.dest('dist'))
    	.pipe(htmlmin(options))
    	.pipe(gulp.dest(htmlDst));
})

//清除js
gulp.task('cjs', function(){
	gulp.src(out+'js/')
    	.pipe(clean());
})
//清除css
gulp.task('ccss', function(){
    gulp.src(out+'css/')
        .pipe(clean());
})
//清除css
gulp.task('cimg', function(){
    gulp.src(out+'images/')
        .pipe(clean());
})
//清除lib
gulp.task('clib', function(){
    gulp.src(out+'lib/')
        .pipe(clean());
})
gulp.task('chtml', function(){
    gulp.src(out+'*.html')
        .pipe(clean());
})


gulp.task('clean', ['cjs', 'ccss', 'cimg', 'clib', 'chtml']);

gulp.task('default', ['css', 'img', 'less', 'js', 'lib', 'html']);

gulp.task('w', function () {
    gulp.watch(src + 'less/*.less', ['less']); //当所有less文件发生改变时，调用testLess任务
    gulp.watch(src + 'css/*.css', ['css']);
    gulp.watch(src + 'js/*.js', ['js']);
    gulp.watch(src + 'img/**/*', ['img']);
    gulp.watch(src + 'lib/*.js', ['lib']);
    gulp.watch(src + '*.html', ['html']);
});
