import { fetchData } from "./fetchData.js";
import { gameConfig } from "./gameConfig.js";
import { config } from "./gameConfig.js";
import { loadAssets } from "./loadAssets.js";
import Phaser from "phaser";

export class PlayGame extends Phaser.Scene {
  constructor() {
    super({ key: "PlayGame" });
    this.gameState = null; // stores game state from the back end
  }

  preload() {
    loadAssets(this, gameConfig);
  }

  async create() {
    this.setupWebSocket();
    this.socket.send(JSON.stringify("joining"));
    this.drawBoard();
  }

  drawBoard() {
    this.children.removeAll(true);
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
      // Check if the message is a Blob
      if (event.data instanceof Blob) {
        event.data.text().then((text) => {
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
    switch (message.messageType) {
      case "gameState":
        this.updateGameState(message.data); // Use the new state from the message
        break;
      default:
        console.warn("Unknown message type:", message.messageType);
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
    }
  }

  loadVertices() {
    // Ensure gameState and its vertices property exist before attempting to use them
    if (!this.gameState || !this.gameState.vertices) {
      console.error("Vertices data is undefined or not an array");
      return;
    }

    let vertices = this.gameState.vertices;

    Object.entries(vertices).forEach(([_, vertice]) => {
      // Check if vertice contains the necessary properties (x, y, id)
      if (
        typeof vertice.x !== "number" ||
        typeof vertice.y !== "number" ||
        typeof vertice.id !== "string"
      ) {
        console.error(
          `Vertice at index ${i} is missing 'x', 'y', or 'id' properties`,
        );
        return;
      }

      let sprite = this.add.sprite(vertice.x, vertice.y, "brick_hex");

      sprite.setDisplaySize(30, 30);
      sprite.setInteractive();
      sprite.setDepth(2);

      // Add click functionality that sends WebSocket message with vertice id
      sprite.on("pointerdown", () => {
        const message = {
          MessageType: "vertexClicked",
          Data: {
            Id: vertice.id,
          },
        };

        // Assuming you have a WebSocket connection stored in this.socket
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
          this.socket.send(JSON.stringify(message));
          console.log("Sent message:", message);
        } else {
          console.error("WebSocket connection is not open");
        }
      });
    });
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
    console.log(newState);
    if (!newState || !newState.tiles || !newState.vertices || !newState.edges) {
      console.error("Received invalid game state from server");
      return;
    }

    this.gameState = newState;
    this.drawBoard();
  }
}

