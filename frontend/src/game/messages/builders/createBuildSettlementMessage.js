export function CreateBuildSettlementMessage(vertexUuid) {
  return {
    MessageType: "buildSettlement",
    Data: {
      VertexUuid: vertexUuid,
    },
  };
}
