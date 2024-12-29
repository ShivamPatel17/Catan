import { Assets } from "assets/loadAssets";
import { CatanCfg } from "config/catanConfig";

/**
 * @param  {Phaser.Scene} scene
 */
export function DrawOpenSettlement(scene, x, y) {
  let sprite = scene.add.sprite(x, y, Assets.OpenSettlement);
  sprite.setDisplaySize(
    CatanCfg.openSettlements.width,
    CatanCfg.openSettlements.height,
  );
  sprite.setInteractive();
  sprite.setDepth(CatanCfg.openSettlements.depth);
  return sprite;
}

export function DrawSettlement(scene, x, y) {
  let sprite = scene.add.sprite(x, y, Assets.Settlement);
  sprite.setDisplaySize(
    CatanCfg.settlements.width,
    CatanCfg.settlements.height,
  );
  sprite.setInteractive();
  sprite.setDepth(CatanCfg.settlements.depth);
  return sprite;
}

