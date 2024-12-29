import {
  DrawOpenSettlement,
  DrawSettlement,
} from "assets/draw/drawSettlements";

/**
 * @param  {Phaser.Scene} scene
 */
export function DrawVertex(scene, vertex) {
  switch (vertex.building) {
    // building == 0 , open settlement
    case 0:
      let openSettleSprite = DrawOpenSettlement(scene, vertex.x, vertex.y);

      // Add click functionality that sends WebSocket message with vertex id
      openSettleSprite.on("pointerdown", () => {
        buildSettlementOnVertex(scene, vertex);
      });
      break;
    // building == 1, built settlement
    case 1:
      DrawSettlement(scene, vertex.x, vertex.y);
      break;
  }
}

function buildSettlementOnVertex(scene, vertex) {
  const message = {
    MessageType: "buildSettlement",
    Data: {
      VertexUuid: vertex.uuid,
    },
  };

  // Assuming you have a WebSocket connection stored in scene.socket
  if (scene.socket && scene.socket.readyState === WebSocket.OPEN) {
    scene.socket.send(JSON.stringify(message));
  } else {
    console.error("WebSocket connection is not open");
  }
}

