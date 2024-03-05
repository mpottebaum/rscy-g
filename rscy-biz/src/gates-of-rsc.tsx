import { createRoot } from 'react-dom/client'

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
