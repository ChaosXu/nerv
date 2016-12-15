var gulp = require('gulp');
var del = require('del')
var sourcemaps = require('gulp-sourcemaps');
var ts = require('gulp-typescript');
var systemjs = require('gulp-systemjs-builder');
var connect = require('gulp-connect');
var npm = require('gulp-npm-files');

var distRoot = '../../../release/nerv/webui';
var dist = distRoot + '/console';


gulp.task('clean', ['clean:ts'], function () {
    del([dist], { force: true });
});

gulp.task('clean:ts', function () {
    del(['js'], { force: true });
})

gulp.task('compile:ts', ['clean:ts'], function () {
    var tsProject = ts.createProject('tsconfig.json');
    return tsProject.src()
        .pipe(sourcemaps.init())
        .pipe(tsProject()).js
        .pipe(sourcemaps.write('.'))
        .pipe(gulp.dest('js'));
});

gulp.task('compile', ['compile:ts'], function () {
    // var builder = systemjs()
    // builder.loadConfigSync('./systemjs.config.js')
    // builder.bundle('js/**/*', {
    //     minify: false,
    //     mangle: false
    // })
    //     .pipe(gulp.dest(dist + '/js'));
});

gulp.task('copy:libs', ['compile'], function () {
    return gulp.src([
        // 'node_modules/core-js/client/shim.min.js',
        // 'node_modules/zone.js/dist/zone.js',
        // 'node_modules/reflect-metadata/Reflect.js',
        // 'node_modules/systemjs/dist/system.src.js',
        'js/**/*',
        'systemjs.config.js'],
        {
            base: './'
        })
        .pipe(gulp.dest(dist))
});

gulp.task('copy:npm', ['compile'], function () {
    return gulp.src(npm(null, './package.json'), { base: './' })
        .pipe(gulp.dest(dist));
})

gulp.task('copy:assets', ['compile'], function () {
    return gulp.src([
        'app/**/*.html',
        '*.html',
        'css/**/*.css',
        'images/**/*',
        'fonts/**/*'],
        {
            base: './'
        })
        .pipe(gulp.dest(dist))
});

gulp.task('connect', function () {
    connect.server({
        root: ['.'],
        port: 3000,
        livereload: true,
        middleware: function (connect, o) {
            var proxy = require('http-proxy-middleware');
            return [
                proxy('/api/objs', {
                    target: 'http://localhost:3333',
                    changeOrigin: true,
                    ws: true
                }),
                proxy('/api/scripts', {
                    target: 'http://localhost:3332',
                    changeOrigin: true,
                    ws: true
                }),
                proxy('/api/templates', {
                    target: 'http://localhost:3332',
                    changeOrigin: true,
                    ws: true
                })
            ];
        }
    });
});

gulp.task('reload:assets', function () {
    return gulp.src(['js/**/*', 'css/**/*', 'images/**/*', 'fonts/**/*'], { base: './' })
        .pipe(gulp.dest('./'))
        .pipe(connect.reload());
})

gulp.task('reload:ts', ['compile:ts'], function () {
    return gulp.src(['js/**/*'], { base: './' })
        .pipe(gulp.dest('./'))
        .pipe(connect.reload());
})

gulp.task('build', ['clean', 'compile', 'copy:libs', 'copy:assets', 'copy:npm']);

gulp.task('default', ['build']);

gulp.task('serve', ['compile:ts', 'connect'], function () {
    gulp.watch('app/**/*', ['reload:ts'])
    gulp.watch('*.html', ['reload:assets'])
    gulp.watch('css/**/*.css', ['reload:assets']);
    gulp.watch('images/**/*', ['reload:assets']);
    gulp.watch('fonts/**/*', ['reload:assets']);
});