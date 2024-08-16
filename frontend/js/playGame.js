import { fetchData } from "./fetchData.js";
import { gameOptions } from "./gameConfig.js";

export class playGame extends Phaser.Scene {
  constructor() {
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
    this.load.image("hexagon", "assets/images/hexagon.png");
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
        const sprite = this.add.sprite(x, y, "hexagon");
        sprite.setScale(2.0);
      }
    } catch (error) {
      console.log("Error loading the hexagon tiles:", error);
    }
  }
}
