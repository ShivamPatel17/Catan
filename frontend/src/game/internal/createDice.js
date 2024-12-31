import { DrawDie } from "assets/draw/drawDie";
import { fetchData } from "utils/fetchData";
import { CatanCfg } from "config/catanConfig";

/**
 * @param  {Phaser.Scene} scene
 */
export function CreateDice(scene) {
  console.log(CatanCfg.dice.red.x);
  let die = DrawDie(scene, CatanCfg.dice.red.x, CatanCfg.dice.red.y);
  scene.input.keyboard.on("keydown-SPACE", () => rollDie(), scene);
  die.on("pointerdown", () => scene.scene.start("MenuScene")); // Start game on click
}

async function rollDie() {
  try {
    const number = await fetchData("http://localhost:3000/roll");
    console.log("Random number from backend:", number);
    const dieNumberToFrame = [1, 2, 5, 6, 4, 0];
    ej;
    die.setFrame(dieNumberToFrame[number - 1]);
  } catch (error) {
    console.error("Error rolling die:", error);
  }
}

