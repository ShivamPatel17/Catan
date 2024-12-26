# Add for MVP

- custom audio for my friends

# Nice to Haves

- see if you can get everything working offline
  - docker offline
  - phaser downloaded for offline

# Try out

- https://pixijs.com/8.x/playground

## Feature Ideas

- make "levels" where you can unlock skins/rewards for winning
- introduce new mechanics to the game such as deciding what resource a hex will provide
- add variations on all the maps
  - maybe some things cost different amounts?
- introduce some kind of battleing system like AOE? turn based strategy
- allow clients to write their algorithms to compete against one another..
  - i really love this Ideas. Let's just do one small step at a time
  - set up API to allow clients to set up
  - put the website up on the public

## nvim stuff

- get a javascript LSP
- get teh matcher to stop bugging on on asteriks

# moving to vite

- ok so ill use vite so i can get some better lsp support

## step i've taken

- adding a package.json with `npm init -y`
- adding a `.vite.config.js`
- rewriting my .js files to use the new phaser module

## to do later

- test that the new set up works offline
  - need to get vite working in docker
  - ok moving over to yarn instead of npm
- remove the vendored phaser.js
- figure out how to deploy this in prod with `npx vite build`

