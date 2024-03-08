import express from 'express'
import path from 'path'
import cors from 'cors'
import { renderToPipeableStream } from 'react-dom/server'
import { createElement } from 'react'
import { RscySqirl } from './rscy-sqirl'

const port = 3000

const app = express()

app.use(cors())

let gatePath = '/gates-of-rsc.js'
if(process.env.NODE_ENV === 'dev') {
  gatePath = '../dist' + gatePath
}
app.use('/gates-of-rsc.js', express.static(path.join(__dirname, gatePath)))

app.use('/gate', (_, res) => {
  const { pipe } = renderToPipeableStream(
    createElement(RscySqirl),
    {
      // keeping full URL here in prep for usage on the-dopness
      bootstrapScripts: ['http://localhost:3000/gates-of-rsc.js'],
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
