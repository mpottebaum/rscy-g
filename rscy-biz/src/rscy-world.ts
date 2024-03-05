import express = require( 'express')
import path = require('path')

const port = 3000

const app = express()

let gatePath = '/js/gates-of-rsc.js'
if(process.env.NODE_ENV === 'development') {
  gatePath = '../dist' + gatePath
}
app.use('/gates-of-rsc.js', express.static(path.join(__dirname, gatePath)))

app.listen(port, () => {
  console.log(`rscy world has risen on port ${port}`)
})
