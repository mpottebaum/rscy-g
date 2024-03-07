// import { createRoot } from 'react-dom/client'
// import {
// //@ts-expect-error we're hella experimental here, bro, dunno what else to say - you simply gotta learn to expect errors round here, ya dig? this is essential, my guy
//   experimental_taintUniqueValue,
// } from 'react';
// import { renderToPipeableStream } from 'react-dom/server';

// const THE_SENSITIVEST_OF_SECRETARY_SECRETS = 'leaked, son -> get wrecked'

// experimental_taintUniqueValue(
//   'Ya done fucked up brah we got all kinds of secrets in the client dog',
//   process,
//   THE_SENSITIVEST_OF_SECRETARY_SECRETS
// );
// NEED TO GET THIS STUFF WORKING IN A COMPONENT
// import path = require('path')
// import env from 'dotenv'

// env.config({
//   path: path.resolve('../.env')
// })

// const dbUrl = process.env.DB_URL
// const dbToken = process.env.DB_TOKEN

// console.log('db', {
//   dbUrl,
//   dbToken,
// })

export function GatesOfRsc() {
  console.log('where am I? am I on the client?')
  return <p>yo world</p>
}

// const domNode = document.getElementById('gate');

// if(domNode) {
//   const root = createRoot(domNode);
//   root.render(<GatesOfRsc />);
// } else {
//   console.error('yikes no node bro')
// }
