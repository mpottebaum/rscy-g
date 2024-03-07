import express = require( 'express')
import path = require('path')
import cors from 'cors'
import { renderToPipeableStream } from 'react-dom/server'
import { GatesOfRsc } from './gates-of-rsc'
import { createElement } from 'react'

const port = 3000

const app = express()

app.use(cors())

let gatePath = '/js/gates-of-rsc.js'
if(process.env.NODE_ENV === 'dev') {
  gatePath = '../dist' + gatePath
}
app.use('/gates-of-rsc.js', express.static(path.join(__dirname, gatePath)))

app.use('/gate', (_, res) => {
  const { pipe } = renderToPipeableStream(
    createElement(GatesOfRsc),
    {
      bootstrapScripts: ['/files/main.js'],
      onShellReady() {
        res.setHeader('content-type', 'text/html');
        pipe(res);
      }
    }
  )
})

app.listen(port, () => {
  console.log(`rscy world has risen on port ${port}`)
})
