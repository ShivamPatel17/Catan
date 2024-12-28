import { catanCfg } from "config/catanConfig";

export function loadAssets(scene) {
  loadDiceSpritesheet(scene);
  loadHexagonImages(scene);
}

function loadDiceSpritesheet(scene) {
  scene.load.spritesheet(
    "redDie",
    "boardgamePack_v2/Spritesheets/diceRed.png",
    {
      frameWidth: catanCfg.diceWidth,
      frameHeight: catanCfg.diceHeight,
    },
  );
}

function loadHexagonImages(scene) {
  const hexagonImages = ["brick", "sheep", "wood", "ore", "wheat"];
  hexagonImages.forEach((hex) => {
    scene.load.image(`${hex}_hex`, `assets/board/hexagon/${hex}.png`);
  });
}

