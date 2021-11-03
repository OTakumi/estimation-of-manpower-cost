var gulp = require('gulp');
var typescript = require('gulp-typescript');

// src ts path
const srcPath = {
    ts: './ts/*.ts',
}

// dist js path
const distPath = {
    js: './dist/js',
}

// Default task
gulp.task('default', () => {
    return gulp.src('./ts/index.ts')
                    .pipe(typescript())
                    .js
                    .pipe(gulp.dest('./dist/js'));
});
