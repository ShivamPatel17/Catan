import { fetchData } from "./fetchData.js";
import { gameOptions } from "./gameConfig.js";
import { config } from "./gameConfig.js";
export class playGame extends Phaser.Scene {
  constructor(config) {
    super("PlayGame");
    this.redDieNum = 1;
  }

  preload() {
    this.load.spritesheet(
      "redDie",
      "boardgamePack_v2/Spritesheets/diceRed.png",
      {
        frameWidth: gameOptions.diceWidth,
        frameHeight: gameOptions.diceHeight,
      }
    );
    this.load.image("brick_hex", "assets/board/hexagon/brick.png");
    this.load.image("sheep_hex", "assets/board/hexagon/sheep.png");
    this.load.image("wood_hex", "assets/board/hexagon/wood.png");
    this.load.image("ore_hex", "assets/board/hexagon/ore.png");
    this.load.image("wheat_hex", "assets/board/hexagon/wheat.png");
  }

  create() {
    this.die = this.add.sprite(700, 550, "redDie").setInteractive();
    this.input.keyboard.on("keydown-SPACE", this.rollDie, this);
    this.loadhex();
  }

  async rollDie() {
    try {
      const number = await fetchData("http://localhost:3000/roll");
      console.log("Random number from backend:", number);
      const dieNumberToFrame = [1, 2, 5, 6, 4, 0];
      this.die.setFrame(dieNumberToFrame[number - 1]);
    } catch (error) {
      console.error("Error rolling die:", error);
    }
  }

  async loadhex() {
    try {
      const hexagons = await fetchData("http://localhost:3000/hexagon");
      console.log(hexagons);
      for (let i = 0; i < hexagons.length; i++) {
        let x = hexagons[i].X;
        let y = hexagons[i].Y;
        let resource = hexagons[i].Resource;

        let sprite;
        switch (resource) {
          case "sheep":
            console.log("adding sheep");
            sprite = this.add.sprite(x, y, "sheep_hex");
            break;
          case "wheat":
            sprite = this.add.sprite(x, y, "wheat_hex");
            break;
          case "ore":
            sprite = this.add.sprite(x, y, "ore_hex");
            break;
          case "wood":
            sprite = this.add.sprite(x, y, "wood_hex");
            break;
          case "brick":
            sprite = this.add.sprite(x, y, "brick_hex");
            break;
        }
        sprite.setDisplaySize(config.HexWidth, config.HexHeight);
        sprite.setInteractive();
        sprite.on("pointerover", function () {
          // https://newdocs.phaser.io/docs/3.80.0/Phaser.GameObjects.GameObject#On
          onHover(sprite);
        });

        // Add the 'pointerout' event listener (optional)
        sprite.on("pointerout", function () {
          onHoverOut(sprite);
        });
      }
    } catch (error) {
      console.log("Error loading the hexagon tiles:", error);
    }
  }
}

// Function to call when hovering over the sprite
function onHover(sprite) {
  console.log("Hovering over sprite");
  // You can also add more actions here, like changing the sprite's appearance
  sprite.setTint(0xff0000); // Change the color of the sprite on hover
}

// Function to call when the pointer leaves the sprite (optional)
function onHoverOut(sprite) {
  console.log("No longer hovering over sprite");
  // Revert any changes made during hover
  sprite.clearTint(); // Revert the color of the sprite
}
