// Import the playGame class from playGame.js
import { playGame } from "./playGame.js";

// Define the base scale
const hexagonImageScale = 0.3;

// Base dimensions
const hexagonImageHeight = 508;
const hexagonImageWidth = 440;

// Export game options
export let gameOptions = {
  diceWidth: 64,
  diceHeight: 64,

  // Hexagon tile configs
  hexagonImageScale: hexagonImageScale,
  hexagonImageHeight: hexagonImageHeight * hexagonImageScale,
  hexagonImageWidth: hexagonImageWidth * hexagonImageScale,
};

// Export game configuration
export let gameConfig = {
  type: Phaser.AUTO,
  backgroundColor: 0x4488aa,
  scale: {
    // mode: Phaser.Scale.FIT,
    autoCenter: Phaser.Scale.CENTER_BOTH,
    parent: "thegame",
    width: 1200,
    height: 900,
  },
  scene: playGame, // Reference to the playGame class
};
