import { DrawVertex } from "game/internal/createVertex";
import { DrawHexagon } from "game/internal/createHexagon";
import { DrawEdge } from "game/internal/createEdge";

/**
 * @param  {Phaser.Scene} scene
 */
export function DrawBoard(scene) {
  if (!scene.gameState) {
    console.error("Edges data is undefined");
    return;
  }
  scene.children.removeAll(scene);
  drawHexagons(scene);
  drawEdges(scene);
  drawVertices(scene);
  drawDie(scene);
}

/**
 * @param  {Phaser.Scene} scene
 */
function drawHexagons(scene) {
  let hexagons = scene.gameState.tiles;
  console.log(scene.gameState);
  console.log(hexagons.length);
  Object.entries(hexagons).forEach(([_, hexagon]) => {
    DrawHexagon(scene, hexagon);
  });
}

/**
 * @param  {Phaser.Scene} scene
 */
function drawVertices(scene) {
  let vertices = scene.gameState.vertices;

  Object.entries(vertices).forEach(([_, vertice]) => {
    // Check if vertice contains the necessary properties (x, y, id)
    if (
      typeof vertice.x !== "number" ||
      typeof vertice.y !== "number" ||
      typeof vertice.uuid !== "string"
    ) {
      console.error(
        `Vertice at index ${i} is missing 'x', 'y', or 'id' properties`,
      );
      return;
    }

    DrawVertex(scene, vertice);
  });
}

/**
 * @param  {Phaser.Scene} scene
 */
function drawEdges(scene) {
  let edges = scene.gameState.edges;
  for (let i = 0; i < edges.length; i++) {
    DrawEdge(scene, edges[i]);
  }
}

/**
 * @param  {Phaser.Scene} scene
 */
function drawDie(scene) {
  scene.die = scene.add.sprite(1000, 800, "redDie").setInteractive();
}

