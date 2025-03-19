import { CatanCfg } from "config/catanConfig";
import { loadAssets } from "assets/loadAssets"; import { DrawBoard } from "game/internal/createBoard";
import Phaser from "phaser";
import { SetupWebSocket } from "utils/setupWebSocket";
import { JoinGame } from "messages/joinGame";

export class CatanGame extends Phaser.Scene {
	constructor() {
		console.log("CatanGame constructor")
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
		// join as a new player
		JoinGame(this)
	}

	// set up from SetupWebSocket in init()
	handleServerMessage(message) {
		let g = {
			playerUuid: message.playerUuid,
			gameState: message.data,
		}
		console.log(g)
		switch (message.messageType) {
			case "gameState":
				this.updateGameState(g); // Use the new state from the message
				break;
			default:
				console.warn("Unknown message type:", message.messageType);
		}
	}

	updateGameState(newState) {
		console.log(newState)
		let gameState = newState.gameState
		if (!gameState || !gameState.tiles || !gameState.vertices || !gameState.edges) {
			console.error("Received invalid game state from server");
			return;
		}

		this.gameState = gameState;
		this.playerUuid = newState.playerUuid
		DrawBoard(this);
	}
}
