# RSC-Y G

trailblazing the Go/RSC stack: Risky a Go Go

Now Playing:
- get databased

Todoozers:

Up next:
- fix the `/rscy` ui
  - mobile responsiveness (flip flex boxes, reduce grid columns, etc.)
  - abbreviate `Busyness` text + some kinda expando
    - for expando: thinking modal on hover kinda like a netflix tile
  - cooler `Dopness` representation
- RSC server

![L.A. Lakers Legends Kareem Abdul-Jabbar, Shaquille O'Neal, George Mikan](/static/pics/lakeys.jpg)

## RSC IT ALL
gotta have it all:
- [go](https://go.dev/doc/install) - to run go
- [air](https://github.com/cosmtrek/air#installation) - to watch go files
- [node](https://nodejs.org/en/download) - to run tailwind/watch views files
- [turso](https://docs.turso.tech/cli/installation) - db, bb

_you can have it all_
![Nine Inch Nailer Trent "Rezzy" Reznor - Tiger Beat](/static/pics/9-incher.jpg)


## if ya gotta set up tha db, bb
in the directory:
```
# open up turso db shell
turso db shell <db name>

# in shell, read .sql file
.read ./db/rscy-gs.sql
```

run it again and you got yourself a fresh new db, ya dig?

## use turso
add deez variables to a `.env` file in the root dir, yo:
```
DB_TOKEN=
DB_URL=
```

## set up the dev env, bb

in the directory:
```
# install tailwind + build css
npm install
npm run build

# tailwind watch /views dir
npm run watch
```

meanwhile in another shell...
start er up:
```
# air watch go code
air
```

we're watching files and rebuilding and shit but we're not reloading it in the browser yet ok sorry you're gonna have to do that yourself for now my bad

## DADA (data)

*RSCY G*
- name - rsc'r
- email - limit number of rscy's by email
- busyness - buncha text
- dopness - integer betweem 0-100 representing **D**ennis h**OP**per**NESS**

![The Dop Man himself, domrade Dennis Hopper](/static/pics/dopper.jpg)

## OTHER STUFF

format `/views`:
```
npm run format
```

lint `/views`:
```
npm run lint
```
