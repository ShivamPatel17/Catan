// Import the playGame class from playGame.js
import { PlayGame } from "./playGame.js";
import { fetchData } from "./fetchData.js";

// Define the base scale
const hexagonImageScale = 0.3;

// Base dimensions
const hexagonImageHeight = 508;
const hexagonImageWidth = 440;

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
    mode: Phaser.Scale.FIT,
    autoCenter: Phaser.Scale.CENTER_BOTH,
    parent: "thegame",
    width: 1200,
    height: 900,
  },
  scene: PlayGame, // Reference to the playGame class
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
