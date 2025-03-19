import { SpriteCfg } from "config/catanConfig";

/**
 * @param  {Phaser.Scene} scene
 */
export function loadAssets(scene) {
  loadDiceSpritesheet(scene);
  loadHexagonImages(scene);
  loadSettlementImages(scene);
  loadRoads(scene);
}

const BrickHex = "brick_hex";
const OreHex = "ore_hex";
const SheepHex = "sheep_hex";
const WoodHex = "wood_hex";
const WheatHex = "wheat_hex";
const OpenSettlement = "open_settlement";
const Settlement = "settlement";
const TopRoad = "top_road";
const RedDie = "red_die";

export const Assets = {
  BrickHex,
  OreHex,
  SheepHex,
  WoodHex,
  WheatHex,
  OpenSettlement,
  Settlement,
  TopRoad,
  RedDie,
};

/**
 * @param  {Phaser.Scene} scene
 */
function loadDiceSpritesheet(scene) {
  scene.load.spritesheet(RedDie, "boardgamePack_v2/Spritesheets/diceRed.png", {
    frameWidth: SpriteCfg.diceWidth,
    frameHeight: SpriteCfg.diceHeight,
  });
}

/**
 * @param  {Phaser.Scene} scene
 */
function loadRoads(scene) {
  scene.load.image(TopRoad, edges_path() + "top_road.png");
}

/**
 * @param  {Phaser.Scene} scene
 */
function loadHexagonImages(scene) {
  const hexagonImages = ["brick", "sheep", "wood", "ore", "wheat"];
  hexagonImages.forEach((hex) => {
    scene.load.image(`${hex}_hex`, hexagon_path() + `${hex}.png`);
  });
}

/**
 * @param  {Phaser.Scene} scene
 */
function loadSettlementImages(scene) {
  scene.load.image(OpenSettlement, vertices_path() + "open_settlement.png");
  scene.load.image(Settlement, vertices_path() + "settlement.png");
}

function hexagon_path() {
  return "assets/board/hexagon/";
}
function vertices_path() {
  return "assets/board/vertices/";
}
function edges_path() {
  return "assets/board/edges/";
}

