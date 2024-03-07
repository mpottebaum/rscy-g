# RSC-Y G

trailblazing the Go/RSC stack: Risky a Go Go

Now Playing:
- RSC server -> aka `rscy-biz`

`rscy-biz` game plan:
  - serve react code from `rscy-biz` and mount it on `the-dopness` via `<script>`
  - `the-dopness` will then need to open some sort of portal into the `rscy-biz` netherworld
    - real-world architectural diagram&#8482;:
    ![Sisyphus or some shit I don't know](/the-dopness/static/pics/Journeys-to-the-Underworld.jpg)

Todoozers:
- get rscy <-- doin
  - yooo ok I think we're cookin here we're gonna stream some react dude
    - [check it](https://react.dev/reference/react-dom/server/renderToPipeableStream) dude we're grabbing the future by the hornz, g
  - NEXT:
    - need a `main.js` served up on `the-dopness`
      - thinking I'll keep all react in `rscy-biz` and build what I need to `the-dopness`
      - hopefully I can just bundle the `react` dependency with the lil bit of js I need in that file
    - lil bit of js I need in `main.js`:
    ```
    // source: "check it" link above
    import { hydrateRoot } from 'react-dom/client';
    import { GatesOfRsc } from './GatesOfRsc.js';

    hydrateRoot(document, <GatesOfRsc />);
    ```
      - now realizing we'll need to include our component as well
- (I think this is OUTDATED now|see above) clean up rsc script tag
  - pull in `rscy-biz` base URL via env

Up next:
- fix the `/rscy` ui
  - mobile responsiveness (flip flex boxes, reduce grid columns, etc.)
  - abbreviate `Busyness` text + some kinda expando
    - for expando: thinking modal on hover kinda like a netflix tile
  - cooler `Dopness` representation
- e2e tests
  - probs [playwright](https://playwright.dev/) but who knows

![L.A. Lakers Legends Kareem Abdul-Jabbar, Shaquille O'Neal, George Mikan](/the-dopness/static/pics/lakeys.jpg)

## RSC IT ALL
gotta have it all:
- [go](https://go.dev/doc/install) - to run go
- [air](https://github.com/cosmtrek/air#installation) - to watch go files
- [node](https://nodejs.org/en/download) - to run tailwind/watch views files
- [turso](https://docs.turso.tech/cli/installation) - db, bb

_you can have it all_
![Nine Inch Nailer Trent "Rezzy" Reznor - Tiger Beat](/the-dopness/static/pics/9-incher.jpg)


## THA DB, bb
**update db schema**
from the root directory:
```
# open up turso db shell
turso db shell <db name>

# in shell, read .sql file
.read ./db/rscy-gs.sql
```

that `rscy-gs.sql` file's gonna drop the current db so ya best tread carefully, ya dig?

## THE DOPNESS, bb
**set up that dev env, yo**
get in `/the-dopness` dir

add deez variables to a `.env` file in that dir, yo:
```
DB_TOKEN=
DB_URL=
```

still inside `/the-dopness`:
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

## RSCY BIZ
**set up the devvy envi**
lil busted at the moment, my dude. first you gotta build the gates of rsc from inside the `/rscy-biz`:
```
npm run build:gates
```

then you can dev that shit, bruv:
```
npm run dev
```

not gonna do a `nodemon` at this time because the `rscy-world` server shouldn't change much if any

**prod stewart**
if you build it
```
npm run build:gates
npm run build:world
```
they will come
```
npm start
```

## DADA (data)

*RSCY G*
- name - rsc'r
- email - limit number of rscy's by email
- busyness - buncha text
- dopness - integer betweem 0-100 representing **D**ennis h**OP**per**NESS**

![The Dop Man himself, domrade Dennis Hopper](/the-dopness/static/pics/dopper.jpg)

## OTHER STUFF

format `/views`:
```
npm run format
```

lint `/views`:
```
npm run lint
```
