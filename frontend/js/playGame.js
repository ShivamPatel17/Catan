import { fetchData } from "./fetchData.js";
import { gameConfig, gameOptions } from "./gameConfig.js";
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
    this.loadVertices();
    this.loadEdges();
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
      const board = await fetchData("http://localhost:3000/board");
      console.log(board);
      let hexagons = board.tiles;
      for (let i = 0; i < hexagons.length; i++) {
        let x = hexagons[i].x;
        let y = hexagons[i].y;
        let resource = hexagons[i].resource;

        let sprite;
        switch (resource) {
          case "sheep":
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
        sprite.setDepth(1);
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
      sprite.setInteractive();
      sprite.on("pointerover", function () {
        // https://newdocs.phaser.io/docs/3.80.0/Phaser.GameObjects.GameObject#On
        onHover(sprite);
      });
    }
  }

  async loadVertices() {
    try {
      const board = await fetchData(gameConfig.baseUrl + "/board");
      let vertices = board.vertices;
      for (let i = 0; i < vertices.length; i++) {
        let vertice = vertices[i];
        let sprite = this.add.sprite(vertice.x, vertice.y, "brick_hex");
        sprite.setDisplaySize(30, 30);
        sprite.setDepth(2);
      }
    } catch (error) {
      console.log("Error loading vertices");
    }
  }

  async loadEdges() {
    try {
      const board = await fetchData(gameConfig.baseUrl + "/board");
      let edges = board.edges;
      console.log(edges);
      for (let i = 0; i < edges.length; i++) {
        let edge = edges[i];
        let sprite = this.add.sprite(edge.x, edge.y, "sheep_hex");
        sprite.setDisplaySize(30, 30);
        sprite.setDepth(2);
      }
    } catch (error) {
      console.log("Error loading edges");
    }
  }
}

function onHover(sprite) {
  sprite.setTint(0xff0000); // Change the color of the sprite on hover
}

function onHoverOut(sprite) {
  sprite.clearTint(); // Revert the color of the sprite
}
