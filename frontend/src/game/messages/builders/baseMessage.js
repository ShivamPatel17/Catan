export function CreateBaseMessage(type, playerUuid) {
  return {
    MessageType: type,
    PlayerUuid: playerUuid,
  };
}
