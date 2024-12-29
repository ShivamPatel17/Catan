import { Assets } from "assets/loadAssets";
import { CatanCfg } from "config/catanConfig";

export function DrawBrickHex(scene, x, y) {
  let sprite = scene.add.sprite(x, y, Assets.BrickHex);
  sprite.setDisplaySize(CatanCfg.hexagons.width, CatanCfg.hexagons.height);
  sprite.setInteractive();
  sprite.setDepth(CatanCfg.hexagons.depth);
  return sprite;
}

export function DrawOreHex(scene, x, y) {
  let sprite = scene.add.sprite(x, y, Assets.OreHex);
  sprite.setDisplaySize(CatanCfg.hexagons.width, CatanCfg.hexagons.height);
  sprite.setInteractive();
  sprite.setDepth(CatanCfg.hexagons.depth);
  return sprite;
}
export function DrawSheepHex(scene, x, y) {
  let sprite = scene.add.sprite(x, y, Assets.SheepHex);
  sprite.setDisplaySize(CatanCfg.hexagons.width, CatanCfg.hexagons.height);
  sprite.setInteractive();
  sprite.setDepth(CatanCfg.hexagons.depth);
  return sprite;
}
export function DrawWoodHex(scene, x, y) {
  let sprite = scene.add.sprite(x, y, Assets.WoodHex);
  sprite.setDisplaySize(CatanCfg.hexagons.width, CatanCfg.hexagons.height);
  sprite.setInteractive();
  sprite.setDepth(CatanCfg.hexagons.depth);
  return sprite;
}
export function DrawWheatHex(scene, x, y) {
  let sprite = scene.add.sprite(x, y, Assets.WheatHex);
  sprite.setDisplaySize(CatanCfg.hexagons.width, CatanCfg.hexagons.height);
  sprite.setInteractive();
  sprite.setDepth(CatanCfg.hexagons.depth);
  return sprite;
}

