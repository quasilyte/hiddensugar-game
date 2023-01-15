A simple game that was created during a live coding session from scratch.

It's written in Go and based on [ebitengine](github.com/hajimehoshi/ebiten/).

## Game Overview

In this game you help a tired gopher to find a sugar for their coffee.

Unfortunately, it's really dark. You can barely see anything.

Try to avoid the traps and reach your glorious goal to have a nice sweet morning coffee!

Every time a trap is activated, it deals some damage. Every next damage from
that trap kind will be doubled aftet that. Therefore, you should avoid
stepping in the same traps too many times.

Bear traps are special: they can't be triggered more than once. All other traps
will hurt you again and again if you enter the same tile.

## Game Controls

* **Move left**: left arrow (keyboard), a key (keyboard), gamepad d-pad left
* **Move right**: right arrow (keyboard), d key (keyboard), gamepad d-pad right
* **Move up**: up arrow (keyboard), w key (keyboard), gamepad d-pad up
* **Move down**: down arrow (keyboard), s key (keyboard), gamepad d-pad down

## Running the Game

```bash
$ git clone github.com/quasilyte/hiddensugar-game
$ cd hiddensugar-game
$ go run .
```
