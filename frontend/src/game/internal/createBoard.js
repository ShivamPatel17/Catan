import { CreateVertex } from "game/internal/createVertex";
import { CreateHexagon } from "game/internal/createHexagon";
import { CreateEdge } from "game/internal/createEdge";
import { CreateDice } from "game/internal/createDice";

/**
 * @param  {Phaser.Scene} scene
 */
export function DrawBoard(scene) {
  if (!scene.gameState) {
    console.error("Edges data is undefined");
    return;
  }
  scene.children.removeAll(scene);
  createHexagons(scene);
  createEdges(scene);
  createVertices(scene);
  createDie(scene);
}

/**
 * @param  {Phaser.Scene} scene
 */
function createHexagons(scene) {
  let hexagons = scene.gameState.tiles;
  console.log(scene.gameState);
  console.log(hexagons.length);
  Object.entries(hexagons).forEach(([_, hexagon]) => {
    CreateHexagon(scene, hexagon);
  });
}

/**
 * @param  {Phaser.Scene} scene
 */
function createVertices(scene) {
  let vertices = scene.gameState.vertices;

  Object.entries(vertices).forEach(([_, vertice]) => {
    CreateVertex(scene, vertice);
  });
}

/**
 * @param  {Phaser.Scene} scene
 */
function createEdges(scene) {
  let edges = scene.gameState.edges;

  return;
  Object.entries(edges).forEach(([_, edge]) => {
    CreateEdge(scene, edge);
  });
}

/**
 * @param  {Phaser.Scene} scene
 */
function createDie(scene) {
  CreateDice(scene);
}
