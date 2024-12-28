// Import the playGame class from playGame.js
import { PlayGame } from "game/catanGame";
import { MenuScene } from "game/menu";
import { fetchData } from "utils/fetchData";
import Phaser from "phaser";

const baseUrl = "http://localhost:3000";

// call this before the Phaser Game starts
export async function loadBackendConfig() {
  try {
    // Fetch the game configuration
    const url = "http://localhost:3000/config";
    let backendConfig = await fetchData(url); // Update the exported variable
    catanCfg = { ...catanCfg, ...backendConfig };
    return true;
  } catch (error) {
    console.error("Error initializing the game: ", error);
  }
}

// Catan board game config
export let catanCfg = {
  diceWidth: 64,
  diceHeight: 64,
};

// Phaser Game Config
export let phaserGameCfg = {
  baseUrl: baseUrl,
  type: Phaser.AUTO,
  backgroundColor: 0x4488aa,
  scale: {
    mode: Phaser.Scale.FIT,
    autoCenter: Phaser.Scale.CENTER_BOTH,
    parent: "thegame",
    width: 1400,
    height: 900,
  },
  scene: [PlayGame, MenuScene], // Reference to the playGame class
};

