import { fetchData } from "utils/fetchData";
import { CatanCfg } from "config/catanConfig";
import { loadAssets } from "assets/loadAssets";
import { DrawBoard } from "game/internal/createBoard";
import Phaser from "phaser";
import { SendWSMessage } from "utils/sendWSMessage";
import { SetupWebSocket } from "utils/setupWebSocket";

export class CatanGame extends Phaser.Scene {
	constructor() {
		super({ key: "CatanGame" });
		this.gameState = null; // stores game state from the back end
	}

	init() {
		console.log("CatanGame.init()")
		SetupWebSocket(this);
	}

	preload() {
		console.log("CatanGame.preload()")
		loadAssets(this, CatanCfg);
	}

	async create() {
		let randomId = Math.round(Math.random() * 100) + 10_000_000
		// join as a new player
		let message = {
			MessageType: "playerConnecting",
			PlayerUuid: "00000000-0000-4000-0000-0000" + randomId.toString()
		};
		SendWSMessage(this.socket, message);
	}

	handleServerMessage(message) {
		switch (message.messageType) {
			case "gameState":
				this.updateGameState(message.data); // Use the new state from the message
				break;
			default:
				console.warn("Unknown message type:", message.messageType);
		}
	}

	updateGameState(newState) {
		if (!newState || !newState.tiles || !newState.vertices || !newState.edges) {
			console.error("Received invalid game state from server");
			return;
		}

		this.gameState = newState;
		DrawBoard(this);
	}
}
