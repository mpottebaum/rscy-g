import express = require( 'express')
import path = require('path')

const port = 3000

const app = express()
app.use('/gates-of-rsc.js', express.static(path.join(__dirname, '/js/gates-of-rsc.js')))

app.listen(port, () => {
  console.log(`rscy world has risen on port ${port}`)
})
