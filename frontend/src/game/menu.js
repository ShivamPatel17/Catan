import { loadAssets } from "assets/loadAssets";
import { catanCfg } from "config/catanConfig";

export class MenuScene extends Phaser.Scene {
  constructor() {
    super({ key: "MenuScene" });
  }

  preload() {
    // defintely don't need to preload all this. Just adding to get quick access to the sprites
    loadAssets(this, catanCfg);
  }

  create() {
    const logo = this.add.sprite(600, 350, "redDie").setInteractive();
    logo.setInteractive();
    logo.on("pointerdown", () => this.scene.start("PlayGame")); // Start game on click
  }
}
