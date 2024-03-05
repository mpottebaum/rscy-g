import { createRoot } from 'react-dom/client'
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

function GatesOfRsc() {
  return <p>yo world</p>
}

const domNode = document.getElementById('gate');

if(domNode) {
  const root = createRoot(domNode);
  root.render(<GatesOfRsc />);
} else {
  console.error('yikes no node bro')
}
