import {
  DrawOpenSettlement,
  DrawSettlement,
} from "assets/draw/drawSettlements";
import { CreateBuildSettlementMessage } from "builders/createBuildSettlementMessage";
import { SendWSMessage } from "utils/sendWSMessage";
/**
 * @param  {Phaser.Scene} scene
 */
export function CreateVertex(scene, vertex) {
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
  const message = CreateBuildSettlementMessage(vertex.uuid);

  SendWSMessage(scene.socket, message);
}
