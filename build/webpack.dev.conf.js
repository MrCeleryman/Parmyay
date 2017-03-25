var config = require("../config");
var webpack = require("webpack");
var merge = require("webpack-merge");
var utils = require("./utils");
var baseWebpackConfig = require("./webpack.base.conf");
var HtmlWebpackPlugin = require("html-webpack-plugin");
var path = require("path");
var ExtractTextPlugin = require("extract-text-webpack-plugin");
var extractCSS = new ExtractTextPlugin("[name].css");

// add hot-reload related code to entry chunks
Object.keys(baseWebpackConfig.entry).forEach(function (name) {
	baseWebpackConfig.entry[name] = ["./build/dev-client"].concat(baseWebpackConfig.entry[name]);
});

module.exports = merge(baseWebpackConfig, {
	module: {
		loaders: utils.styleLoaders({ sourceMap: config.dev.cssSourceMap, extract: true }),
		rules: [{
			test: /\.css$/,
			use: extractCSS.extract({
				use: ["css-loader"],
				fallback: "style-loader"
			})
		}]
	},
	// eval-source-map is faster for development
	devtool: "#eval-source-map",
	plugins: [
		extractCSS,
		new ExtractTextPlugin(utils.assetsPath("css/[name].[contenthash].css")),
		new webpack.DefinePlugin({
			"process.env": config.dev.env
		}),
		// https://github.com/glenjamin/webpack-hot-middleware#installation--usage
		new webpack.optimize.OccurrenceOrderPlugin(),
		new webpack.HotModuleReplacementPlugin(),
		new webpack.NoEmitOnErrorsPlugin(),
		// https://github.com/ampedandwired/html-webpack-plugin
		new HtmlWebpackPlugin({
			filename: "index.html",
			template: "src/index.html",
			inject: true
		})
	]
});