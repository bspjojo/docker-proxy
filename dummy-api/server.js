const express = require('express')
const app = express()
const port = 3000

app.use((req, res, next) => {
    console.log(req.path, req.query)
    res.send(req.path)
})

app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`)
})