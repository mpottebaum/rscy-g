import express = require( 'express')
import path = require('path')
import cors from 'cors'

const port = 3000

const app = express()

app.use(cors())

let gatePath = '/js/gates-of-rsc.js'
if(process.env.NODE_ENV === 'dev') {
  gatePath = '../dist' + gatePath
}
app.use('/gates-of-rsc.js', express.static(path.join(__dirname, gatePath)))

app.listen(port, () => {
  console.log(`rscy world has risen on port ${port}`)
})
