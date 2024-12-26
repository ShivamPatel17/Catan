// Import the playGame class from playGame.js
import { PlayGame } from "./playGame.js";
import { MenuScene } from "./menuscene.js";
import { fetchData } from "./fetchData.js";

import Phaser from "phaser";
const baseUrl = "http://localhost:3000";

// Export game options
export let gameOptions = {
  diceWidth: 64,
  diceHeight: 64,
};

// Export game configuration
export let gameConfig = {
  baseUrl: baseUrl,
  type: Phaser.AUTO,
  backgroundColor: 0x4488aa,
  scale: {
    mode: Phaser.Scale.EXPAND,
    autoCenter: Phaser.Scale.CENTER_BOTH,
    parent: "thegame",
    width: 1200,
    height: 900,
  },
  scene: [PlayGame, MenuScene], // Reference to the playGame class
};

export let config = null; // Declare and export gameConfigGo

(async function () {
  try {
    // Fetch the game configuration
    const url = "http://localhost:3000/config";
    let backendConfig = await fetchData(url); // Update the exported variable
    console.log(backendConfig);
    config = { ...backendConfig };
  } catch (error) {
    console.error("Error initializing the game: ", error);
  }
})();

