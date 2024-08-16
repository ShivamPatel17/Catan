import { gameConfig } from "./gameConfig.js";
import { playGame } from "./playGame.js";

window.onload = function () {
  gameConfig.scene = playGame; // Ensure the scene is set
  let game = new Phaser.Game(gameConfig);
  window.focus();
};
