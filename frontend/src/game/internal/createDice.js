import { DrawDie } from "assets/draw/drawDie";
import { fetchData } from "utils/fetchData";
import { CatanCfg } from "config/catanConfig";

/**
 * @param  {Phaser.Scene} scene
 */
export function CreateDice(scene) {
	let die = DrawDie(scene, CatanCfg.dice.red.x, CatanCfg.dice.red.y);
	scene.input.keyboard.on("keydown-SPACE", () => rollDie(), scene);
	die.on("pointerdown", () => scene.scene.start("MenuScene")); // Start game on click
}

async function rollDie() {
	try {
		const number = await fetchData("http://localhost:3000/roll");
		const dieNumberToFrame = [1, 2, 5, 6, 4, 0];
		die.setFrame(dieNumberToFrame[number - 1]);
	} catch (error) {
		console.error("Error rolling die:", error);
	}
}

