// NEED TO GET THIS STUFF WORKING IN A COMPONENT
// import path = require('path')
// import env from 'dotenv'

import React, { useState } from 'react'

// env.config({
//   path: path.resolve('../.env')
// })

// const dbUrl = process.env.DB_URL
// const dbToken = process.env.DB_TOKEN

// console.log('db', {
//   dbUrl,
//   dbToken,
// })

async function dbCall() {
  console.log('look at me callin a db. where could I be?')
  return [
    {
      name: 'yo',
    }
  ]
}

export function RscySqirl() {
  const [ rscyGs, setRscyGs ] = useState<{ name: string;}[]>([])
  console.log('where am I? am I on the client?')
  
  async function sup() {
    const rgs = await dbCall()
    setRscyGs(rgs)
  }
  return (
    <section>
      <p>yo world</p>
      <button
        onClick={sup}
      >rsc it all</button>
      {rscyGs.length > 0 && (
        <ul>
          {rscyGs.map(({name})=>(
            <li key={name}>
              {name}
            </li>
          ))}
        </ul>
      )}
    </section>
  )
}
