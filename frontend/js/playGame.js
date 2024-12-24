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
    console.log(this.gameState);
    this.setupWebSocket();
    this.die = this.add.sprite(400, 350, "redDie").setInteractive();
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
    };

    // Handle incoming messages from the server
    this.socket.onmessage = (event) => {
      console.log("MESSAGE RECEIVED!");
      console.log(event); // Log the entire event

      // Check if the message is a Blob
      if (event.data instanceof Blob) {
        event.data.text().then((text) => {
          console.log("Raw message data:", text); // Log the raw message data as a string
          try {
            const message = JSON.parse(text); // Parse the string as JSON
            console.log("Parsed JSON message:", message); // Log the parsed JSON
            this.handleServerMessage(message); // Call your handler with the parsed message
          } catch (error) {
            console.error("Error parsing JSON:", error);
          }
        });
      } else {
        // If the message is not a Blob, handle it as a string (if needed)
        try {
          const message = JSON.parse(event.data); // Parse it directly as JSON
          console.log("Parsed JSON message:", message);
          this.handleServerMessage(message);
        } catch (error) {
          console.error("Error parsing JSON:", error);
        }
      }
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
    console.log(message);
    console.log("handling server message?!?");
    switch (message.messageType) {
      case "gameState":
        this.updateGameState(message.data); // Use the new state from the message
        break;
      case "actionResult":
        this.processActionResult(message.data);
        break;
      // Handle other message types...
      default:
        console.warn("Unknown message type:", message.messageType);
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
    console.log("fetching load hex");
    // Ensure gameState and its tiles property exist before attempting to use them
    if (!this.gameState || !this.gameState.tiles) {
      console.error("Game state or tiles are undefined");
      return;
    }

    let hexagons = this.gameState.tiles;
    console.log(hexagons);
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
      setOnHover(sprite);
    }
  }

  loadVertices() {
    // Ensure gameState and its vertices property exist before attempting to use them
    if (!this.gameState || !this.gameState.vertices) {
      console.error("Vertices data is undefined or not an array");
      return;
    }

    let vertices = this.gameState.vertices;

    for (let i = 0; i < vertices.length; i++) {
      let vertice = vertices[i];

      // Check if vertice contains the necessary properties (x, y, id)
      if (
        typeof vertice.x !== "number" ||
        typeof vertice.y !== "number" ||
        typeof vertice.id !== "string"
      ) {
        console.error(
          `Vertice at index ${i} is missing 'x', 'y', or 'id' properties`,
        );
        continue;
      }

      let sprite = this.add.sprite(vertice.x, vertice.y, "brick_hex");

      sprite.setDisplaySize(30, 30);
      sprite.setInteractive();
      sprite.setDepth(2);

      // Set hover functionality
      setOnHover(sprite);

      // Add click functionality that sends WebSocket message with vertice id
      sprite.on("pointerdown", () => {
        const message = {
          MessageType: "vertexClicked",
          Data: vertice.id,
        };

        // Assuming you have a WebSocket connection stored in this.socket
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
          this.socket.send(JSON.stringify(message));
          console.log("Sent message:", message);
        } else {
          console.error("WebSocket connection is not open");
        }
      });
    }
  }

  loadEdges() {
    if (!this.gameState || !this.gameState.edges) {
      console.error("Edges data is undefined");
      return;
    }

    let edges = this.gameState.edges;
    for (let i = 0; i < edges.length; i++) {
      let edge = edges[i];
      let sprite = this.add.sprite(edge.x, edge.y, "sheep_hex");
      sprite.setDisplaySize(30, 30);
      sprite.setDepth(2);
    }
  }

  updateGameState(newState) {
    // Ensure newState exists and has the expected properties before updating game state
    console.log(newState);
    if (!newState || !newState.tiles || !newState.vertices || !newState.edges) {
      console.error("Received invalid game state from server");
      return;
    }

    console.log(newState);
    // Update the local game state with the new state
    this.gameState = newState;
    console.log("HERE!!!");
    console.log(this.gameState);

    // commenting because I think this does nothing
    //// Clear the current game objects (e.g., tiles, vertices, edges)
    //this.clearGameObjects();

    // Re-render the game objects with the new state
    this.loadhex();
    this.loadVertices();
    this.loadEdges();
  }

  clearGameObjects() {
    // Remove all existing sprites (e.g., hexes, vertices, edges)
    this.children.removeAll(); // This removes all game objects from the scene
  }
}

function setOnHover(sprite) {
  sprite.on("pointerover", function () {
    sprite.setTint(0xff0000); // Change the color of the sprite on hover
  });
  sprite.on("pointerout", function () {
    sprite.clearTint(); // Remove the tint on pointer out
  });
}

