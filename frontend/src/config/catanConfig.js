import { CatanGame } from "game/catanGame";
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
    CatanCfg = { ...CatanCfg, ...backendConfig };
    return true;
  } catch (error) {
    console.error("Error initializing the game: ", error);
  }
}

// sprite config
export let SpriteCfg = {
  diceWidth: 64,
  diceHeight: 64,
};

// Frontend Catan board game config
export let CatanCfg = {
  dice: {
    red: {
      x: 1200,
      y: 800,
    },
    depth: 2,
    width: 80,
    height: 80,
  },
  hexagons: {
    width: 440 * 0.4, // hex sprite is 440px wide
    height: 508 * 0.4, // hex sprite is 508px tall
    depth: 1,
  },
  settlements: {
    depth: 2,
    width: 60,
    height: 60,
  },
  openSettlements: {
    depth: 2,
    width: 20,
    height: 20,
  },
  roads: {
    depth: 2,
    width: 30,
    height: 20,
  },
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
  scene: [CatanGame, MenuScene], // Reference to the playGame class
};

