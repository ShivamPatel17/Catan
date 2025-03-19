import { SendWSMessage } from "utils/sendWSMessage";

export function JoinGame(scene) {
	let randomId = Math.round(Math.random() * 100) + 10_000_000
	// join as a new player
	let message = {
		MessageType: "playerConnecting",
		PlayerUuid: "00000000-0000-4000-0000-0000" + randomId.toString()
	};
	SendWSMessage(scene.socket, message);
}
