var gulp = require('gulp');
var del = require('del')
var sourcemaps = require('gulp-sourcemaps');
var ts = require('gulp-typescript');
var systemjs = require('gulp-systemjs-builder');
var connect = require('gulp-connect');

var distRoot = '../../release/nerv/webui';
var dist = distRoot + '/console';


gulp.task('clean', function () {
    return del([dist, 'js'], { force: true });
});

gulp.task('compile:ts', ['clean'], function () {
    var tsProject = ts.createProject('tsconfig.json');
    return tsProject.src()
        //.pipe(sourcemaps.init())
        .pipe(tsProject()).js
        //.pipe(sourcemaps.write('.'))
        .pipe(gulp.dest('js'));
});

gulp.task('compile', ['compile:ts'], function () {
    var builder = systemjs()
    builder.loadConfigSync('./systemjs.config.js')
    builder.bundle('js/**/*', {
        minify: false,
        mangle: false
    })
        .pipe(gulp.dest(dist + '/js'));
});

gulp.task('copy:libs', ['clean'], function () {
    return gulp.src([
        'node_modules/core-js/client/shim.min.js',
        'node_modules/zone.js/dist/zone.js',
        'node_modules/reflect-metadata/Reflect.js',
        'node_modules/systemjs/dist/system.src.js',
        'systemjs.config.js'],
        {
            base: './'
        })
        .pipe(gulp.dest(dist))
});

gulp.task('copy:assets', ['clean'], function () {
    return gulp.src([
        'app/**/*.html',
        'index.html',
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
        root:['.'],
        port: 3000,
        livereload: true
    });
});

gulp.task('reload', function () {
    gulp.src(['app/**/*.html', 'js/**/*', 'css/**/*', 'images/**/*', 'fonts/**/*'], { base: './' })
        .pipe(gulp.dest('./'))
        .pipe(connect.reload());
})

gulp.task('build', ['compile', 'copy:libs', 'copy:assets']);

gulp.task('default', ['build']);

gulp.task('serve', ['compile:ts', 'connect'],function () {
    gulp.watch('app/**/*', ['compile:ts'])
    gulp.watch('index.html', ['reload'])
    gulp.watch('css/**/*.css', ['reload']);
    gulp.watch('images/**/*', ['reload']);
    gulp.watch('fonts/**/*', ['reload']);
    gulp.watch('js/**/*', ['reload']);
});