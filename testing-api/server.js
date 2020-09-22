const express = require('express')
const app = express()
const port = 3000

app.use((req, res, next) => {
    let { path, query, headers, body, url, originalUrl, baseUrl, method, hostname } = req;
    let t = { path, query, headers, body, url, originalUrl, baseUrl, method, hostname };
    console.log(t);
    res.send(t)
})

app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`)
})