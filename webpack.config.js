const pkg = require("./package.json")
const webpack = require("webpack");
const path = require("path");
const CleanPlugin = require('clean-webpack-plugin');

module.exports = {
  target: 'web',
  resolve: {
    modules: [
      'node_modules'
    ],
    extensions: ['.js', '.jsx', '.json', '.ts', '.tsx']
  },
  entry: path.resolve(__dirname, "./src/app.tsx"),
  output: {
    path: path.resolve(__dirname, "./public/js"),
    filename: "bundle.js",
    libraryTarget: "umd"
  },
  module: {
    rules: [
      {
        test: /\.tsx?$/,
        loader: "ts-loader",
        exclude: /node_modules/,
        options: {
          transpileOnly: true,
          compilerOptions: {
            target: "es5"
          }
        }
      },
      {
        test: /\.scss?$/,
        loader: "style-loader!css-loader!sass-loader"
      }
    ]
  },
  stats: "minimal",
  plugins: [
    new CleanPlugin([
      "publis/js/bundle.js"
    ], {
      verbose: true
    }),
    new webpack.DefinePlugin({
      "global.GENTLY": false,
      'process.env.BROWSER_ENV': JSON.stringify(true)
    }),
    new webpack.optimize.ModuleConcatenationPlugin()
  ]
};
