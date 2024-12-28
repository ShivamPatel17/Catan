import { catanCfg } from "config/catanConfig";

export function DrawBoard(scene) {
  scene.children.removeAll(scene);
  drawHexagons(scene);
  drawEdges(scene);
  drawVertices(scene);
}

function drawHexagons(scene) {
  // Ensure gameState and its tiles property exist before attempting to use them
  if (!scene.gameState || !scene.gameState.tiles) {
    console.error("Game state or tiles are undefined");
    return;
  }

  let hexagons = scene.gameState.tiles;
  for (let i = 0; i < hexagons.length; i++) {
    let x = hexagons[i].x;
    let y = hexagons[i].y;
    let resource = hexagons[i].resource;

    let sprite;
    switch (resource) {
      case "sheep":
        sprite = scene.add.sprite(x, y, "sheep_hex");
        break;
      case "wheat":
        sprite = scene.add.sprite(x, y, "wheat_hex");
        break;
      case "ore":
        sprite = scene.add.sprite(x, y, "ore_hex");
        break;
      case "wood":
        sprite = scene.add.sprite(x, y, "wood_hex");
        break;
      case "brick":
        sprite = scene.add.sprite(x, y, "brick_hex");
        break;
    }
    sprite.setDisplaySize(catanCfg.HexWidth, catanCfg.HexHeight);
    sprite.setInteractive();
  }
}

function drawVertices(scene) {
  // Ensure gameState and its vertices property exist before attempting to use them
  if (!scene.gameState || !scene.gameState.vertices) {
    console.error("Vertices data is undefined or not an array");
    return;
  }

  let vertices = scene.gameState.vertices;

  Object.entries(vertices).forEach(([_, vertice]) => {
    // Check if vertice contains the necessary properties (x, y, id)
    if (
      typeof vertice.x !== "number" ||
      typeof vertice.y !== "number" ||
      typeof vertice.id !== "string"
    ) {
      console.error(
        `Vertice at index ${i} is missing 'x', 'y', or 'id' properties`,
      );
      return;
    }

    let sprite = scene.add.sprite(vertice.x, vertice.y, "brick_hex");

    sprite.setDisplaySize(30, 30);
    sprite.setInteractive();
    sprite.setDepth(2);

    // Add click functionality that sends WebSocket message with vertice id
    sprite.on("pointerdown", () => {
      const message = {
        MessageType: "vertexClicked",
        Data: {
          Id: vertice.id,
        },
      };

      // Assuming you have a WebSocket connection stored in scene.socket
      if (scene.socket && scene.socket.readyState === WebSocket.OPEN) {
        scene.socket.send(JSON.stringify(message));
        console.log("Sent message:", message);
      } else {
        console.error("WebSocket connection is not open");
      }
    });
  });
}

function drawEdges(scene) {
  if (!scene.gameState || !scene.gameState.edges) {
    console.error("Edges data is undefined");
    return;
  }

  let edges = scene.gameState.edges;
  for (let i = 0; i < edges.length; i++) {
    let edge = edges[i];
    let sprite = scene.add.sprite(edge.x, edge.y, "sheep_hex");
    sprite.setDisplaySize(30, 30);
    sprite.setDepth(2);
  }
}

