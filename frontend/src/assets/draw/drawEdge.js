import { CatanCfg } from "config/catanConfig";
import { Assets } from "assets/loadAssets";

/**
 * @param  {Phaser.Scene} scene
 */
export function DrawTopRoad(scene, x, y) {
  let sprite = scene.add.sprite(x, y, Assets.TopRoad);
  sprite.setDisplaySize(CatanCfg.roads.width, CatanCfg.roads.height);
  sprite.setDepth(CatanCfg.roads.depth);
}

