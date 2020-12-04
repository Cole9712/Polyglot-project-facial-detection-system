module.exports={
    devServer: {
        proxy: {
            '^/uploadMinePost': {
                target: 'http://127.0.0.1:8082/',
                secure: false,
                ws: true,
                changeOrigin:true,
                headers: {
                    Connection: 'keep-alive'
                }
                // pathRewrite: {'^/api' : ''}
            },
            '^/sockjs-node': {
                target: 'http://127.0.0.1:5555/',
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