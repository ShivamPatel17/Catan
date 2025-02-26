import { CatanCfg } from "config/catanConfig";
import { Assets } from "assets/loadAssets";

export function DrawDie(scene, x, y) {
	let die = scene.add.sprite(x, y, Assets.RedDie);

	die.setDisplaySize(CatanCfg.dice.width, CatanCfg.dice.height);
	die.setInteractive();
	die.setDepth(CatanCfg.dice.depth);

	return die;
}

