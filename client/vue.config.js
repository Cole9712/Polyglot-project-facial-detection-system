module.exports={
    devServer: {
        proxy: {
            '^/uploadMine': {
                target: 'http://127.0.0.1:8082/',
                secure: false,
                ws: true,
                changeOrigin:true,
                headers: {
                    Connection: 'keep-alive'
                }
                // pathRewrite: {'^/api' : ''}
            }
        }
    }
}