import Image from "next/image";

export default function RscyBb() {
  async function weOnLaServer() {
    'use server'
    console.log('we strictly server here bro')
    return { data: 'mmm lil bits', }
  }
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <p>this is a rsc I guess?</p>
      {/* <button onClick={weOnLaServer}>here be RSCs</button> */}
    </main>
  );
}
