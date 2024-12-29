import {
  DrawBrickHex,
  DrawOreHex,
  DrawSheepHex,
  DrawWoodHex,
  DrawWheatHex,
} from "assets/draw/drawHexagon";

const Brick = "brick";
const Ore = "ore";
const Sheep = "sheep";
const Wood = "wood";
const Wheat = "wheat";

export function DrawHexagon(scene, hexagon) {
  let x = hexagon.x;
  let y = hexagon.y;
  let resource = hexagon.resource;

  switch (resource) {
    case Brick:
      DrawBrickHex(scene, x, y);
      break;
    case Ore:
      DrawOreHex(scene, x, y);
      break;
    case Sheep:
      DrawSheepHex(scene, x, y);
      break;
    case Wood:
      DrawWoodHex(scene, x, y);
      break;
    case Wheat:
      DrawWheatHex(scene, x, y);
      break;
  }
}

