import { phaserGameCfg, loadBackendConfig } from "config/catanConfig";
import Phaser from "phaser";

export let gameConfigGo = null; // Declare and export gameConfigGo

(async function () {
  try {
    // Fetch the game configuration
    var configIsloaded = await loadBackendConfig();

    if (configIsloaded) {
      console.log("starting game");
      let game = new Phaser.Game(phaserGameCfg);
      window.focus();
    } else {
      console.error("Failed to load game configuration.");
    }
  } catch (error) {
    console.error("Error initializing the game: ", error);
  }
})();

