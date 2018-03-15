var slsw = require('serverless-webpack')
var nodeExternals = require('webpack-node-externals')
var path = require('path')

module.exports = {
  mode: 'development',
  entry: slsw.lib.entries,
  resolve: {
    extensions: [
      '.js',
      '.json',
      '.ts',
      '.tsx'
    ]
  },
  output: {
    libraryTarget: 'commonjs',
    path: path.join(__dirname, '.webpack'),
    filename: '[name].js'
  },
  target: 'node',
  module: {
    rules: [
      {
        test: /\.ts(x?)$/,
        loader: 'ts-loader'
      }
    ]
  },
  externals: [nodeExternals()]
}
