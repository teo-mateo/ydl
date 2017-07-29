var path = require('path');
var HtmlWebpackPlugin = require('html-webpack-plugin');
var UglifyJsPlugin = require('uglifyjs-webpack-plugin');

var exps = {
	entry: './app/index.js', 
	output: {
		path: path.resolve(__dirname, 'dist'),
		filename: 'index_bundle.js',
		publicPath: ''
	},
	module: {
		rules: [
			{ test: /\.(js)$/, use: 'babel-loader' },
			{ test: /\.css$/, use: ['style-loader', 'css-loader']}
		]
	},
	devServer: {
		historyApiFallback: true
	},
	plugins: [
		new HtmlWebpackPlugin({
			template: 'app/index.html'
		}), 
		new UglifyJsPlugin({
			minimize: true,
			extractComments: true
		})
	], 
	//devtool: 'eval-source-map'
}

module.exports = exps;