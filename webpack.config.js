const slsw = require("serverless-webpack")
const nodeExternals = require("webpack-node-externals")

module.exports = {
        mode: "development",
        entry: slsw.lib.entries,
        module: {
                rules: [
                        {
                                test: /\.js$/,
                                exclude: /(node_modules|bower_components)/,
                                use: "babel-loader"
                        },
                ]
        },
        externals: [nodeExternals()],
}
