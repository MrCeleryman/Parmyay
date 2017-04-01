// https://github.com/shelljs/shelljs
require("./check-versions")();
var shell = require("shelljs");
var utils = require("./utils");
process.env.NODE_ENV = "production";

var path = require("path");
var config = require("../config");
var ora = require("ora");
var webpack = require("webpack");
var webpackConfig = require("./webpack.prod.conf");

utils.apiSetup();

console.log(
	"  Tip:\n" +
	"  Built files are meant to be served over an HTTP server.\n" +
	"  Opening index.html over file:// won\"t work.\n"
);

var spinner = ora("building for production...");
spinner.start();

var assetsPath = path.join(config.build.assetsRoot, config.build.assetsSubDirectory);
shell.rm("-rf", assetsPath);
shell.mkdir("-p", assetsPath);

// See if static directory exists
if (shell.test("-e", "static") && shell.test("-d", "static")) {
	shell.cp("-R", "static/*", assetsPath);
}

webpack(webpackConfig, function (err, stats) {
	spinner.stop()
	if (err) throw err
	process.stdout.write(stats.toString({
		colors: true,
		modules: false,
		children: false,
		chunks: false,
		chunkModules: false
	}) + "\n")
});
