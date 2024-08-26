import { fetchData } from "./fetchData.js";
import { gameConfig } from "./gameConfig.js";
import { config } from "./gameConfig.js";
import { loadAssets } from "./loadAssets.js";
export class PlayGame extends Phaser.Scene {
  constructor() {
    super({ key: "PlayGame" });
    this.gameState = null; // stores game state from the back end
  }

  preload() {
    loadAssets(this, gameConfig);
  }

  async create() {
    // Fetch the initial game state from the backend
    this.gameState = await this.fetchGameState();
    this.setupWebSocket();
    this.die = this.add.sprite(700, 550, "redDie").setInteractive();
    this.input.keyboard.on("keydown-SPACE", this.rollDie, this);
    this.loadhex();
    this.loadVertices();
    this.loadEdges();
  }

  setupWebSocket() {
    // Create a new WebSocket connection
    this.socket = new WebSocket("ws://localhost:3000/ws");

    // Handle the connection opening
    this.socket.onopen = () => {
      console.log("WebSocket connection established");

      // Optionally send an initial message to the server
      this.socket.send(JSON.stringify({ type: "join", data: {} }));
    };

    // Handle incoming messages from the server
    this.socket.onmessage = (event) => {
      const message = JSON.parse(event.data);
      this.handleServerMessage(message);
    };

    // Handle connection closure
    this.socket.onclose = () => {
      console.log("WebSocket connection closed");
    };

    // Handle connection errors
    this.socket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };
  }

  handleServerMessage(message) {
    switch (message.type) {
      case "gameState":
        this.updateGameState(message.data);
        break;
      case "actionResult":
        this.processActionResult(message.data);
        break;
      // Handle other message types...
      default:
        console.warn("Unknown message type:", message.type);
    }
  }

  async fetchGameState() {
    try {
      const response = await fetch("http://localhost:3000/board");
      return await response.json();
    } catch (error) {
      console.error("Failed to fetch game state:", error);
    }
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

  loadhex() {
    let hexagons = this.gameState.tiles;
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
  }

  loadVertices() {
    let vertices = this.gameState.vertices;
    for (let i = 0; i < vertices.length; i++) {
      let vertice = vertices[i];
      let sprite = this.add.sprite(vertice.x, vertice.y, "brick_hex");
      sprite.setDisplaySize(30, 30);
      sprite.setDepth(2);
    }
  }

  loadEdges() {
    let edges = this.gameState.edges;
    console.log(edges);
    for (let i = 0; i < edges.length; i++) {
      let edge = edges[i];
      let sprite = this.add.sprite(edge.x, edge.y, "sheep_hex");
      sprite.setDisplaySize(30, 30);
      sprite.setDepth(2);
    }
  }

  update() {
    // listen to backend
    // if new game state is the same as backend, only move the sprites that are necessary to the new position
    // if the game state is different, rerender everything
  }
}

function onHover(sprite) {
  sprite.setTint(0xff0000); // Change the color of the sprite on hover
}

function onHoverOut(sprite) {
  sprite.clearTint(); // Revert the color of the sprite
}
