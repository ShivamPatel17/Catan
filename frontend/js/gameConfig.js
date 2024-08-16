// Import the playGame class from playGame.js
import { playGame } from "./playGame.js";

// Export game options
export let gameOptions = {
  diceWidth: 64,
  diceHeight: 64,
  cardScale: 0.8,
};

// Export game configuration
export let gameConfig = {
  type: Phaser.AUTO,
  backgroundColor: 0x4488aa,
  scale: {
    mode: Phaser.Scale.FIT,
    autoCenter: Phaser.Scale.CENTER_BOTH,
    parent: "thegame",
    width: 800,
    height: 600,
  },
  scene: playGame, // Reference to the playGame class
};
