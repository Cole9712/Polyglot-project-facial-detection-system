module.exports={
    devServer: {
        proxy: {
            '/uploadMine': {
                target: 'http://127.0.0.1:8082'
                // pathRewrite: {'^/api' : ''}
            }
        }
    }
}