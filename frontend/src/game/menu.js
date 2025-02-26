import { loadAssets } from "assets/loadAssets";
import { CatanCfg } from "config/catanConfig";

export class MenuScene extends Phaser.Scene {
	constructor() {
		super({ key: "MenuScene" });
	}

	preload() {
		loadAssets(this, CatanCfg);
	}

	create() {
		const logo = this.add.sprite(600, 350, "red_die").setInteractive();
		logo.setInteractive();
		logo.on("pointerdown", () => this.scene.start("CatanGame")); // Start game on click
	}
}

