//index.js
const httpProxy = require('express-http-proxy');
const express = require('express');
 
const app = express();
 
function selectProxyHost(req) {
    if (req.path.startsWith('/admin'))
        return 'http://localhost:8001/';
    else if (req.path.startsWith('/todos'))
        return 'http://localhost:8002/';
}
 
app.use((req, res, next) => {
    httpProxy(selectProxyHost(req))(req, res, next);
});
 
app.listen(8000, () => {
    console.log('API Gateway running!');
});