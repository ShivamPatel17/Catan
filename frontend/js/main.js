import { gameConfig } from "./gameConfig.js";
import { PlayGame } from "./playGame.js";
import { fetchData } from "./fetchData.js";

export let gameConfigGo = null; // Declare and export gameConfigGo

(async function () {
  try {
    // Fetch the game configuration
    const url = "http://localhost:3000/config";
    gameConfigGo = await fetchData(url); // Update the exported variable
    // console.log(gameConfigGo);

    if (gameConfigGo) {
      let game = new Phaser.Game(gameConfig);
      window.focus();
    } else {
      console.error("Failed to load game configuration.");
    }
  } catch (error) {
    console.error("Error initializing the game: ", error);
  }
})();
