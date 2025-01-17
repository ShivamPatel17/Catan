import { fetchData } from "utils/fetchData";
import { CatanCfg } from "config/catanConfig";
import { loadAssets } from "assets/loadAssets";
import { DrawBoard } from "game/internal/createBoard";
import Phaser from "phaser";

export class CatanGame extends Phaser.Scene {
  constructor() {
    super({ key: "CatanGame" });
    this.gameState = null; // stores game state from the back end
  }

  preload() {
    loadAssets(this, CatanCfg);
  }

  async create() {
    this.setupWebSocket();
    this.socket.send(JSON.stringify("joining"));
    DrawBoard(this);
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

  updateGameState(newState) {
    if (!newState || !newState.tiles || !newState.vertices || !newState.edges) {
      console.error("Received invalid game state from server");
      return;
    }

    this.gameState = newState;
    DrawBoard(this);
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
}

