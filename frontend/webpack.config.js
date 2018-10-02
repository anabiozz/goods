const path = require("path")
const webpack = require("webpack")

module.exports = (env, argv) => {
  mode = argv.mode == "production" ? true : false
  return {
    entry: {
      bundle: "./src/scripts/main.js"
    },
    plugins: [
      new webpack.DefinePlugin({
        PRODUCTION: JSON.stringify(mode)
      })
    ],
    output: {
      path: path.resolve(__dirname, "../backend/static/scripts/"),
      filename: "bundle.js",
      chunkFilename: "[id].js"
    },
    module: {
      rules: [{
        test: /\.js$/,
        exclude: [
          path.resolve(__dirname, "node_modules"),
        ],
        loader: "babel-loader",
        query: {
          presets: ["es2015"],
        },
      }]
    },
    optimization: {
      minimize: true
    }
  }
}