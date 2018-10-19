const path = require('path');
const webpack = require('webpack');

module.exports = {
    mode: 'development',
    entry: path.resolve(__dirname, './index.js'), //指定入口文件，程序从这里开始编译,__dirname当前所在目录, ../表示上一级目录, ./同级目录
    output: {
        path: path.resolve(__dirname, '../static/js'), // 输出的路径
        filename: 'a.js'  // 打包后文件
    },
    module: {
        rules: [
            {
                test: /\.(js|jsx)$/,
                loader: 'babel-loader',
                exclude: /node_modules/
            },
            {
                test: /\.svg/,
                exclude: /node_modules/,
                loader: "file-loader"
            },
            {
                test: /\.jpg/,
                exclude: /node_modules/,
                loader: "file-loader"
            },
            {
                test: /\.scss$/,
                use: [
                    { loader: "file-loader", options: { name: "[name].css" } },
                    {
                        loader: "sass-loader",
                        options: {
                            outputStyle: "compressed",
                            includePaths: ["./node_modules"]
                        }
                    }
                ]
            }
        ]
    }
}